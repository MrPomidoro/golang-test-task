package email

import "github.com/go-gomail/gomail"

type emailService struct {
	dialer   *gomail.Dialer
	sender   string
	host     string
	port     int
	username string
	password string
}

func NewEmailService(sender, host string, port int, username, password string) *emailService {
	return &emailService{
		sender:   sender,
		host:     host,
		port:     port,
		username: username,
		password: password,
		dialer:   gomail.NewDialer(host, port, username, password),
	}
}

func (s *emailService) SendHTML(message, subject, receiver string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.sender)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	err := s.dialer.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
