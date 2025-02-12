package main

import "github.com/go-playground/validator/v10"

type User struct {
	Username string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"gte=18,lte=60"` // Age must be between 18-60
}

//simple test
func test1() {
	validate := validator.New()

	user := User{
		Username: "GoDev",
		Email:    "invalid-email",
		Age:      17,
	}

	errs := validate.Struct(user)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			print("Field: ", err.Field(), ", Error: ", err.Tag())
			println("")
		}
	} else {
		print("Success")
	}
}
