package models

type UserRole string

// NOTE: Several struct specs in one type declaration are allowed.

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
	}

	App struct {
		Version string `validate:"len:5"`
	}
	App1 struct {
		Version struct {
			Build int `validate:"min:18|max:50"`
		}
		Vqweqe struct {
			Build []int `validate:"min:18|max:50"`
		}
		Vqweqe1 App
	}
	App2 struct {
		App1
		Build int `validate:"min:18|max:50"`
	}
)

type Token struct {
	Header    []byte
	Payload   []byte
	Signature []byte
}

type Response struct {
	Code int    `validate:"in:200,404,500"`
	Body string `json:"omitempty"`
}

//go:generate go-validate ./models.go
