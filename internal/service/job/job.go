package job

import (
	"context"
	"fmt"
	"math"
)

type repositoryCredit interface {
	AddCredit(ctx context.Context, username string, amount int) error
	SubtractCredit(ctx context.Context, username string, amount int) error
}

type repositoryEmail interface {
	GetEmail(ctx context.Context, username string) (string, error)
}

type EmailService interface {
	SendHTML(message, subject, receiver string) error
}

type Service struct {
	repository      repositoryCredit
	repositoryEmail repositoryEmail
	email           EmailService
}

func NewJobService(repository repositoryCredit, repositoryEmail repositoryEmail, email EmailService) Service {
	return Service{repository: repository, repositoryEmail: repositoryEmail, email: email}
}

func (s Service) AddCredit(ctx context.Context, username string, amount int) error {
	return s.repository.AddCredit(ctx, username, amount)
}

func (s Service) SubtractCredit(ctx context.Context, username string, amount int) error {
	if err := s.repository.SubtractCredit(ctx, username, amount); err != nil {
		email, err := s.GetEmail(ctx, username)
		if err != nil {
			return err
		}
		if err := s.SendHTML("Лимит вашего кредита закончился!<br> Обратитесь к администратору или пополните счет", "", email); err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetEmail(ctx context.Context, username string) (string, error) {
	return s.repositoryEmail.GetEmail(ctx, username)
}

func (s Service) SendHTML(message, subject, receiver string) error {
	return s.email.SendHTML(message, subject, receiver)
}

func (s Service) SlowTaskMissingNumbers(ctx context.Context, nums []int, username string) error {

	message := SlowTaskMissingNumbers + fmt.Sprintf("\n\n %v", findMissingNumbers(nums))

	email, err := s.GetEmail(ctx, username)
	if err != nil {
		return err
	}

	return s.SendHTML(message, "", email)
}

func findMissingNumbers(nums []int) []int {
	missingNumbers := []int{}

	// Пройдемся по всем элементам массива
	for i := 0; i < len(nums); i++ {
		// Берем абсолютное значение текущего элемента
		// и используем его как индекс
		index := int(math.Abs(float64(nums[i]))) - 1

		// Если элемент по индексу больше нуля, делаем его отрицательным
		// это обозначает, что число присутствует в массиве
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	// Пройдемся по массиву еще раз
	for i, num := range nums {
		// Если число положительное, это означает, что индекс + 1
		// не присутствует в массиве, и это одно из пропущенных чисел
		if num > 0 {
			missingNumbers = append(missingNumbers, i+1)
		}
	}

	return missingNumbers
}
