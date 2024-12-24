package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func rpc(rounds int) (int, int) {
	pcnt := 0
	ccnt := 0
	fmt.Println("Let's Play Rock,Paper,Scissors!You Have", rounds, "Rounds.")
	for i := 0; i < rounds; i++ {
		choicenum := rand.Intn(rounds)
		var comchoice string
		switch choicenum {
		case 0:
			comchoice = "Rock"
		case 1:
			comchoice = "Paper"
		case 2:
			comchoice = "Scissors"
		}
		var playerchoice string
		fmt.Println("Enter your choice:")
		fmt.Scan(&playerchoice)
		if !slices.Contains([]string{"Rock", "Paper", "Scissors"}, playerchoice) {
			panic("Invalid choice:")
		}
		fmt.Println("Computer chose:", comchoice)
		if playerchoice == comchoice {
			fmt.Println("It's a tie!")
		} else if (playerchoice == "Rock" && comchoice == "Scissors") || (playerchoice == "Paper" && comchoice == "Rock") || (playerchoice == "Scissors" && comchoice == "Paper") {
			fmt.Println("Player Win This Round!")
			pcnt++
		} else {
			fmt.Println("Computer Win This Round!")
			ccnt++
		}
	}
	return pcnt, ccnt
}
func main() {
	const rounds = 3
	pcnt, ccnt := rpc(rounds)
	if pcnt == ccnt {
		fmt.Println("After ", rounds, " Rounds", " Match Tie!")
	} else if pcnt > ccnt {
		fmt.Println("After ", rounds, " Rounds", " Player Win The Match")
	} else {
		fmt.Println("After ", rounds, " Rounds", " Computer Win The Match")
	}
}
