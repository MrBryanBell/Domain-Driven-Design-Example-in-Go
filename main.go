package main

import (
	"fmt"
	UserRegister "mymodule/User/Application/Register"
	User "mymodule/User/Domain"
	InMemoryUserRepository "mymodule/User/Infrastructure/InMemory"
	"strings"
)

func mainn() {

	{
		fmt.Println("---- TESTING AGAINST AN EMPTY NAME ----")
		/* It Should fail if name is empty  */
		var _, err = User.Create("", "woods@my.com")

		AssertNot(err, nil)
	}

	{
		fmt.Println("---- TESTING AGAINST AN INVALID SHORT NAME ----")
		/* It Should fail if name is lower than 3 characters  */
		var _, err = User.Create("Li", "woods@my.com")

		AssertNot(err, nil)
	}

	{
		fmt.Println("---- TESTING AGAINST AN INVALID LONG NAME ----")
		var INVALID_NAME = strings.Repeat("a", 51)
		/* It Should fail if name is lower than 3 characters  */
		var _, err = User.Create(INVALID_NAME, "woods@my.com")

		AssertNot(err, nil)
	}

	{
		fmt.Println("---- TESTING AGAINST AN INVALID EMAIL EXTENSION ----")
		var INVALID_EMAIL = "pupusa@hotmail.com"
		/* It Should fail if name is lower than 3 characters  */
		var _, err = User.Create("Bryan", INVALID_EMAIL)

		AssertNot(err, nil)
	}

	{
		fmt.Println("---- REGISTERING A USER WITH AN INVALID EMAIL ----")
		var repository = InMemoryUserRepository.New()
		var register = UserRegister.Setup(repository)
		var error = register.Run("Bryan", "woods@my.com")

		AssertNot(error, nil)
	}

}

const MESSAGE_PASSED_TEST = "✅ PASSED TEST"
const MESSAGE_FAILED_TEST = "❌ ERROR"

func Assert(actual any, expected any) {
	if expected == actual {
		fmt.Println(MESSAGE_PASSED_TEST)
	} else {
		fmt.Println(MESSAGE_FAILED_TEST)
	}
}

func AssertNot(actual any, expected any) {
	if expected != actual {
		fmt.Println(MESSAGE_PASSED_TEST)
	} else {
		fmt.Println(MESSAGE_FAILED_TEST)
	}
}

/* LESSON: Al utilziar definición de Alias,
estamos hablando siempre de la misma estructura, solo que ahora
podemos llamarle también con otro nombre (Alias). En el código de abajo,
incluso añadí el método Price() únicamente al Alias de "Coursito", y para
mi sopresa, también puedo fue añadido en el type original "Course".
*/

type Course struct {
	name  string
	price float64
}

func (self Course) Name() string {
	return self.name
}

type Coursito = Course

func (self Coursito) Price() float64 {
	return self.price
}

func main() {
	{
		/* .Price() method works on "Course" */
		var course = Course{name: "Golang", price: 100.0}
		fmt.Println(course.Name())
		fmt.Println(course.Price())
	}

	{
		/* .Price() method works on "Coursito" */
		var course = Coursito{name: "Golang", price: 100.0}
		fmt.Println(course.Name())
		fmt.Println(course.Price())
	}
}
