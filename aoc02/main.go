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
		case Rock:
			return 1
		case Paper:
			return 2
		case Scissors:
			return 3
		}

		panic("unreachable!")
	}

	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	outcomeScore := func(opponent, me RPS) int {
		switch opponent {
		case Rock:
			{
				switch me {
				case Rock:
					return 3
				case Paper:
					return 6
				case Scissors:
					return 0
				}
			}

		case Paper:
			{
				switch me {
				case Rock:
					return 0
				case Paper:
					return 3
				case Scissors:
					return 6
				}
			}

		case Scissors:
			{
				switch me {
				case Rock:
					return 6
				case Paper:
					return 0
				case Scissors:
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

	getRPS := func(component RPS, result Result) RPS {
		switch component {
		case Rock:
			{
				switch result {
				case Lose:
					return Scissors
				case Draw:
					return Rock
				case Win:
					return Paper
				}
			}
		case Paper:
			{
				switch result {
				case Lose:
					return Rock
				case Draw:
					return Paper
				case Win:
					return Scissors
				}
			}
		case Scissors:
			{
				switch result {
				case Lose:
					return Paper
				case Draw:
					return Scissors
				case Win:
					return Rock
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
		case Rock:
			{
				newRound.me = getRPS(round.opponent, Lose)
			}

		// need to end in a draw
		case Paper:
			{
				newRound.me = getRPS(round.opponent, Draw)

			}

		// need to win
		case Scissors:
			{
				newRound.me = getRPS(round.opponent, Win)
			}
		}
	}

	part1(newRounds)
}

type RPS int

const (
	Rock     RPS = iota
	Paper        = iota
	Scissors     = iota
)

type Result int

const (
	Win  Result = iota
	Draw        = iota
	Lose        = iota
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
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	default:
		panic(fmt.Errorf("invalid RPS string %s", str))
	}

}
