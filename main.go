package main

type User struct {
	Username string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"gte=18,lte=60"` // Age must be between 18-60
}

func main() {
	test1()

}
