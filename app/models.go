package app

type Member struct {
	ID    int    `json:"id" gorm:"primary_key""`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type Members []Member
