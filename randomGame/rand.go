package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func input(s string)int  {
	userGuess := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez une valeur : ")
	userGuess.Scan()
	userInput := userGuess.Text()
	UserTrueInput, err := strconv.Atoi(userInput)
	for err != nil {
		fmt.Print("entrer une bonne valeur : ")
				userGuess.Scan()
				userInput := userGuess.Text() 
				NewUserTrueInput, err := strconv.Atoi(userInput)
				UserTrueInput = NewUserTrueInput
				if err == nil {
					break
				}
	}
	return UserTrueInput
}

func main() {
	random := rand.Intn(100)
	fmt.Printf("random integer: %v\n", random)
	/*
	userGuess := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez quelque chose : ")
	userGuess.Scan()
	userInput := userGuess.Text()
	UserTrueInput, err := strconv.Atoi(userInput)
	for err != nil {
		fmt.Print("entrer une bonne valeur : ")
				userGuess.Scan()
				userInput := userGuess.Text() 
				NewUserTrueInput, err := strconv.Atoi(userInput)
				UserTrueInput = NewUserTrueInput
				if err == nil {
					break
				}
	}
	*/
	var UserInputint string
	UserTrueInput := input(UserInputint)
	tentative := 3
	for i := 1; i <= tentative; i++ {
		//		fmt.Printf("random integer: %v\n", random)
		if UserTrueInput == random {
			fmt.Println("good guess you nik the game")
			break
		}
		
		if UserTrueInput < random {
			fmt.Println(tentative)
			fmt.Println(i)
			fmt.Println("trop petit")
			fmt.Print("essayer encore : ")
			/*
			userGuess.Scan()
			userInput := userGuess.Text()
			NewUserTrueInput, err := strconv.Atoi(userInput)
			var UserInputint2 string
			NewUserTrueInput := input(UserInputint2)
			*/
			UserTrueInput = NewUserTrueInput
			for err != nil {
				fmt.Print("entrer une bonne valeur : ")
				userGuess.Scan()
				userInput := userGuess.Text() 
				NewUserTrueInput, err := strconv.Atoi(userInput)
				UserTrueInput = NewUserTrueInput
				if err == nil {
					break
				}
			}
			if i >= tentative{
				fmt.Println("Game start again")
				userGuess.Scan()
				userInput := userGuess.Text() 
				NewUserTrueInput,_ := strconv.Atoi(userInput)
				UserTrueInput = NewUserTrueInput
				i = 0
			}
			} else if UserTrueInput > random {
				fmt.Println(tentative)
				fmt.Println(i)
				fmt.Println("trop grand")
				fmt.Print("réessayer a nouveau : ")
				userGuess.Scan()
				userInput := userGuess.Text() 
				NewUserTrueInput, err := strconv.Atoi(userInput)
				UserTrueInput = NewUserTrueInput
				for err != nil {
					fmt.Print("entrer une bonne valeur : ")
					userGuess.Scan()
					userInput := userGuess.Text() 
					NewUserTrueInput, err := strconv.Atoi(userInput)
					UserTrueInput = NewUserTrueInput
					if err == nil {
						break
					}
				}
				if i >= tentative{
					fmt.Println("Game start again")
					i = 0
				}
			}
		}
	}