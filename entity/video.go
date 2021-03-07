package entity

type Person struct {
	FirsName string `json: "firstname" binding: "requiered"`
	LastName string `json: "lastname" binding: "requiered"`
	Age      int8   `json: "age" binding: "gte=1, lte=130"`
	Email    string `json: "email" validate: "requiered, email"`
}

type Video struct {
	Title       string `json: "title" binding: "min=2, max=10" validate:"is-cool"`
	Description string `json: "description" binding: "max=20"`
	URL         string `json: "url" binding: "required,url"`
	Auth        string `json: "author" binding: "required"`
}
