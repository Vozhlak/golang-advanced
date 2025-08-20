package verify

type RequestVerify struct {
	Email string `json:"email" validate:"required,email"`
}

type ResponseVerify struct {
	Status string
}
