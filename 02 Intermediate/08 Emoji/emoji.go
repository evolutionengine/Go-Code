// Prints emoji on the terminal
package main

import (
	"fmt"

	"gopkg.in/kyokomi/emoji.v1"
)

func main() {
	fmt.Println("Welcome to the world of emoji's !")
	emoji.Println(":beer: Beer!!!")
	fmt.Println("")
	emoji.Println("Behold the fury of :dragon: Dragon Slayer ... :skull: :cyclone: !!!")
}
