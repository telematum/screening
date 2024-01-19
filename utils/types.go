package utils

type CreateUserType struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserType struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
