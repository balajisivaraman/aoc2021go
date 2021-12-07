package solutions

import (
	"strings"
	"github.com/balajisivaraman/aoc2021go/utils"
)

type Command struct {
	Movement string
	Amount int
}

func (command *Command) isMoveForward() bool {
	return command.Movement == "forward"
}

func (command *Command) isMoveDown() bool {
	return command.Movement == "down"
}

func (command *Command) isMoveUp() bool {
	return command.Movement == "up"
}

type Coordinate struct {
	HorizontalPosition int
	Depth int
}

func (coordinate *Coordinate) moveForward(by int) {
	coordinate.HorizontalPosition += by
}

func (coordinate *Coordinate) moveDown(by int) {
	coordinate.Depth += by
}

func (coordinate *Coordinate) moveUp(by int) {
	coordinate.Depth -= by
}

func Day02Part1(input []string) int {
	coordinate := Coordinate {
		HorizontalPosition: 0,
		Depth: 0,
	}
	commands := []Command{}
	for i := 0; i < len(input); i++ {
		split := strings.Split(input[i], " ")
		command := Command {
			Movement: split[0],
			Amount: utils.ParseStringToIntIgnoringError(split[1]),
		}
		commands = append(commands, command)
	}
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		if command.isMoveForward() {
			coordinate.moveForward(command.Amount)
		}
		if command.isMoveDown() {
			coordinate.moveDown(command.Amount)
		}
		if command.isMoveUp() {
			coordinate.moveUp(command.Amount)
		}
	}
	return coordinate.HorizontalPosition * coordinate.Depth
}
