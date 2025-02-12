package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Test4 struct {
	St1 string `validate:"required"`
	St2 string
	St3 string `validate:"min=2,max=4"`
	St4 string `validate:"min=2,max=4"`
	St5 string `validate:"min=2,max=4"`
}

// customerize message
func test4() {
	validate := validator.New()

	test := Test4{
		St3: "1",
		St4: "12345",
		St5: "1234",
	}

	errs := validate.Struct(test)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			print("Field: ", err.Field(), ", Error: ", err.Tag(),
				", Threshold: ", err.Param(),
				", Current Val: ", fmt.Sprintf("%v", err.Value()))
			println("")
		}
	} else {
		print("Success")
	}
}
