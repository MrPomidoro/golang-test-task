package dto

type CreateTaskDTO struct {
	Description string `json:"description"`
	Cost        int    `json:"cost"`
}

func NewCreateDTO(description string, cost int) CreateTaskDTO {
	return CreateTaskDTO{
		Description: description,
		Cost:        cost,
	}
}
