package types

import validation "github.com/go-ozzo/ozzo-validation"

// response struct | marshalled into json fromat from struct
type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookName"`
	Author      string `json:"author"`
	Publication string `json:"publication,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.Author,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(1, 50)))
}

// package types

// import validation "github.com/go-ozzo/ozzo-validation"

// response struct | marshalled into JSON format from struct
type AuthorRequest struct {
	ID            uint   `json:"authorID"`
	AuthorName    string `json:"authorName"`
	AuthorAddress string `json:"authorAddress"`
	Phone         string `json:"phone,omitempty"`
}

func (author AuthorRequest) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.AuthorName,
			validation.Required.Error("First name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&author.AuthorAddress,
			validation.Required.Error("Last name cannot be empty"),
			validation.Length(1, 50)))
}

type UserRequest struct {
	ID			 string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
}

func (user UserRequest) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username,
			validation.Required.Error("username cannot be empty"),
			validation.Length(3, 50)),
		validation.Field(&user.Password,
			validation.Required.Error("password cannot be empty"),
			validation.Length(1, 50)))
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Validate validates the request body for the LoginRequest.
func (request LoginRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.UserName,
			validation.Required.Error("Username cannot be empty")),
		validation.Field(&request.Password,
			validation.Required.Error("Password cannot be empty")))
}