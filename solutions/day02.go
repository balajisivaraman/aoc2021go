package solutions

import (
	"strings"
	"github.com/balajisivaraman/aoc2021go/utils"
)

type Command struct {
	Movement string
	Amount int
}

func (command *Command) isForwardCommand() bool {
	return command.Movement == "forward"
}

func (command *Command) isDownCommand() bool {
	return command.Movement == "down"
}

func (command *Command) isUpCommand() bool {
	return command.Movement == "up"
}

type Coordinate struct {
	HorizontalPosition int
	Depth int
	Aim int
}

func (coordinate *Coordinate) moveForwardWrong(by int) {
	coordinate.HorizontalPosition += by
}

func (coordinate *Coordinate) moveDownWrong(by int) {
	coordinate.Depth += by
}

func (coordinate *Coordinate) moveUpWrong(by int) {
	coordinate.Depth -= by
}

func (coordinate *Coordinate) increaseAim(by int) {
	coordinate.Aim += by
}

func (coordinate *Coordinate) decreaseAim(by int) {
	coordinate.Aim -= by
}

func (coordinate *Coordinate) forward(by int) {
	coordinate.HorizontalPosition += by
	coordinate.Depth += (coordinate.Aim * by)
}

func Day02Part1(input []string) int {
	coordinate := Coordinate {
		HorizontalPosition: 0,
		Depth: 0,
	}
	commands := parseInputToCommand(input)
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		if command.isForwardCommand() {
			coordinate.moveForwardWrong(command.Amount)
		}
		if command.isDownCommand() {
			coordinate.moveDownWrong(command.Amount)
		}
		if command.isUpCommand() {
			coordinate.moveUpWrong(command.Amount)
		}
	}
	return coordinate.HorizontalPosition * coordinate.Depth
}

func Day02Part2(input []string) int {
	coordinate := Coordinate {
		HorizontalPosition: 0,
		Depth: 0,
		Aim: 0,
	}
	commands := parseInputToCommand(input)
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		if command.isForwardCommand() {
			coordinate.forward(command.Amount)
		}
		if command.isDownCommand() {
			coordinate.increaseAim(command.Amount)
		}
		if command.isUpCommand() {
			coordinate.decreaseAim(command.Amount)
		}
	}
	return coordinate.HorizontalPosition * coordinate.Depth
}

func parseInputToCommand(input []string) []Command {
	commands := []Command{}
	for i := 0; i < len(input); i++ {
		split := strings.Split(input[i], " ")
		command := Command {
			Movement: split[0],
			Amount: utils.ParseStringToIntIgnoringError(split[1]),
		}
		commands = append(commands, command)
	}
	return commands
}
