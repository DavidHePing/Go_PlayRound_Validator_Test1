package main

import "github.com/go-playground/validator/v10"

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
