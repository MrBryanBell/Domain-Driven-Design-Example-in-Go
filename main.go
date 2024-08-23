package main

import (
	"fmt"
	UserRegister "mymodule/User/Application/Register"
	UserUpdater "mymodule/User/Application/Updater"
	User "mymodule/User/Domain"
	InMemoryUserRepository "mymodule/User/Infrastructure/InMemory"
	"strings"
)

func main() {

	{
		fmt.Println("---- TESTING AGAINST AN EMPTY NAME ----")
		/* It Should fail if name is empty  */
		var _, err = User.Create("", "", "woods@my.com")

		AssertNot(err, nil)
	}

	{
		fmt.Println("---- TESTING AGAINST AN INVALID SHORT NAME ----")
		/* It Should fail if name is lower than 3 characters  */
		var _, err = User.Create("", "Li", "woods@my.com")

		AssertNot(err, nil)
	}

	{
		fmt.Println("---- TESTING AGAINST AN INVALID LONG NAME ----")
		var INVALID_NAME = strings.Repeat("a", 51)
		/* It Should fail if name is longer than 50 characters  */
		var _, err = User.Create("", INVALID_NAME, "woods@my.com")

		AssertNot(err, nil)
	}

	{
		fmt.Println("---- TESTING AGAINST AN INVALID EMAIL EXTENSION ----")
		var INVALID_EMAIL = "pupusa@hotmail.com"
		/* It Should fail if email domain is not gmail  */
		var _, err = User.Create("", "Bryan", INVALID_EMAIL)

		AssertNot(err, nil)
	}

	{
		fmt.Println("---- REGISTERING A USER WITH AN INVALID EMAIL ----")
		var repository = InMemoryUserRepository.New()
		var register = UserRegister.Setup(repository)
		var error = register.Run("", "Bryan", "woods@my.com")

		AssertNot(error, nil)
	}

	{
		fmt.Println("---- FINDING A USER WITH UNEXISTENT ID ----")
		/* It should fail if ID doesn't exist */
		const UNEXISTENT_ID = "10"
		var repository = InMemoryUserRepository.New()
		var _, err = repository.FindByID(UNEXISTENT_ID)

		AssertNot(err, nil)
		

		fmt.Println("---- REGISTERING A USER WITH VALID DATA ----")
		/* It should create a valid user */
		repository = InMemoryUserRepository.New()
		var register = UserRegister.Setup(repository)
		err = register.Run("1", "Bryan", "bebell@gmail.com")
		Assert(err, nil)


		fmt.Println("---- FINDING A USER BY A VALID ID ----")
		/* It should find a user with a valid ID */
		const VALID_ID = "1"
		var user, _ = repository.FindByID(VALID_ID)
		
		Assert(user.ID, VALID_ID)


		fmt.Println("---- UPDATING A USER EMAIL ----")
		/* It should update a user email */
		var updater = UserUpdater.Setup(repository)
		newUser, _ := updater.Run("1", "Bryan", "newemail@gmail.com")

		AssertNot(user, newUser)


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


/* LESSON: Al utilizar definición de Alias, 
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

func mainn() {
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
