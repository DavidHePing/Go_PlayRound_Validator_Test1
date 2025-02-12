package main

import "github.com/go-playground/validator/v10"

type Test2 struct {
	Num1 int `required,validate:"min=18,max=60"` // required is useless
	Num2 int `validate:"min=18,max=60"`
	Num3 int `validate:"min=18,max=60"`
	Num4 int `validate:"min=18,max=60"`
}

func test2() {
	validate := validator.New()

	test := Test2{
		Num2: 61,
		Num3: 17,
		Num4: -4,
	}

	errs := validate.Struct(test)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			print("Field: ", err.Field(), ", Error: ", err.Tag())
			println("")
		}
	} else {
		print("Success")
	}
}
