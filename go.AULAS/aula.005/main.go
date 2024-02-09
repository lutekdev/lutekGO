package main

import (
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("lutekdev@gmail.com")
	fmt.Println(erro)
}
