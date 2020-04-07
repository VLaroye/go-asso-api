package app

type Member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Members []Member
