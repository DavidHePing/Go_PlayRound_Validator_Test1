package main

import "github.com/go-playground/validator/v10"

type Test3 struct {
	Num1 int `validate:"min=18,max=60"`
	Num2 int `validate:"min=18,max=60"`
}

//customerize message
func test3() {
	validate := validator.New()

	test := Test3{
		Num1: 17,
		Num2: 61,
	}

	errs := validate.Struct(test)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			print("Field: ", err.Field(), ", Error: ", err.Tag(),
				", Threshold: ", err.Param(), ", Current Val: ", err.Value())
			println("")
		}
	} else {
		print("Success")
	}
}
