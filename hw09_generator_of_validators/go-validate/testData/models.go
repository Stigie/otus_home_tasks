package testData

type UserRole string

type User struct {
	ID     string `json:"id" validate:"len:36"`
	Name   string
	Age    int      `validate:"min:18|max:50"`
	Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
	Role   UserRole `validate:"in:admin,stuff"`
	Phones []string `validate:"len:11"`
}

type App struct {
	Version string `validate:"len:5"`
}

type App1 struct {
	Version struct {
		Build int `validate:"min:18|max:50"`
	}
	Vqweqe struct {
		Build []int `validate:"min:18|max:50"`
	}
	Vqweqe1 App
}
