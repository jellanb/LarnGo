package main

//bcryp is a package to protect information from users
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `clave123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	claveLogin := `clave1234`
	err = CompareHashAndPassword(bs, []byte(claveLogin))
	if err != nil {
		fmt.Println("password is incorrect")
		return
	}
	fmt.Println("you are login ok!")

}
