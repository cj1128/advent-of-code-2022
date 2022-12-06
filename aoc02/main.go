package main

import (
	"fmt"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	parsed := parse(utils.ReadStdin())

	// 13484
	// part1(parsed)

	// 13433
	part2(parsed)
}

func part1(rounds []Round) {
	result := 0

	// 1 for Rock, 2 for Paper, and 3 for Scissors
	shapeSelectedScore := func(rps RPS) int {
		switch rps {
		case "rock":
			return 1
		case "paper":
			return 2
		case "scissors":
			return 3
		}

		panic("unreachable!")
	}

	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	outcomeScore := func(opponent, me RPS) int {
		switch opponent {
		case "rock":
			{
				switch me {
				case "rock":
					return 3
				case "paper":
					return 6
				case "scissors":
					return 0
				}
			}

		case "paper":
			{
				switch me {
				case "rock":
					return 0
				case "paper":
					return 3
				case "scissors":
					return 6
				}
			}

		case "scissors":
			{
				switch me {
				case "rock":
					return 6
				case "paper":
					return 0
				case "scissors":
					return 3
				}
			}
		}

		panic("unreachable!")
	}

	for _, round := range rounds {
		result += shapeSelectedScore(round.me) + outcomeScore(round.opponent, round.me)
	}

	fmt.Println(result)
}

// X(rock) means you need to lose, Y(paper) means you need to end the round in a draw, and Z(scissors) means you need to win
func part2(rounds []Round) {
	newRounds := make([]Round, len(rounds))

	getRPS := func(component RPS, result string) RPS {
		switch component {
		case "rock":
			{
				switch result {
				case "lose":
					return "scissors"
				case "draw":
					return "rock"
				case "win":
					return "paper"
				}
			}
		case "paper":
			{
				switch result {
				case "lose":
					return "rock"
				case "draw":
					return "paper"
				case "win":
					return "scissors"
				}
			}
		case "scissors":
			{
				switch result {
				case "lose":
					return "paper"
				case "draw":
					return "scissors"
				case "win":
					return "rock"
				}
			}
		}

		panic("unreachable!")
	}

	for idx, round := range rounds {
		newRound := &newRounds[idx]

		newRound.opponent = round.opponent

		switch round.me {
		// need to lose
		case "rock":
			{
				newRound.me = getRPS(round.opponent, "lose")
			}

		// need to end in a draw
		case "paper":
			{
				newRound.me = getRPS(round.opponent, "draw")

			}

		// need to win
		case "scissors":
			{
				newRound.me = getRPS(round.opponent, "win")
			}
		}
	}

	part1(newRounds)
}

type RPS string

const (
	Rock     RPS = "rock"
	Paper    RPS = "paper"
	Scissors RPS = "scissors"
)

type Round struct {
	opponent RPS
	me       RPS
}

func parse(str string) []Round {
	var result []Round

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		parts := strings.Split(line, " ")

		result = append(result, Round{
			opponent: toRPS(parts[0]),
			me:       toRPS(parts[1]),
		})
	}

	return result
}

func toRPS(str string) RPS {
	switch str {
	case "A":
		return "rock"
	case "B":
		return "paper"
	case "C":
		return "scissors"
	case "X":
		return "rock"
	case "Y":
		return "paper"
	case "Z":
		return "scissors"
	default:
		panic(fmt.Errorf("invalid RPS string %s", str))
	}

}
