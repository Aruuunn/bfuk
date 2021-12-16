package bfuk

import (
	"errors"
	"io"
)

type CommandTape struct {
	commandPointer int
	parsedCommands []rune
	commandReader  io.RuneReader
}

func NewCommandTape(commandReader io.RuneReader) *CommandTape {
	return &CommandTape{
		commandPointer: -1,
		parsedCommands: make([]rune, 0),
		commandReader:  commandReader,
	}
}

func (ct *CommandTape) GetCurrentCommand() rune {
	return ct.parsedCommands[ct.commandPointer]
}

func (ct *CommandTape) MoveRight() error {
	if ct.commandPointer == len(ct.parsedCommands)-1 {
		r, _, err := ct.commandReader.ReadRune()

		if err != nil {
			return err
		}

		ct.parsedCommands = append(ct.parsedCommands, r)
	}

	ct.commandPointer++

	return nil
}

func (ct *CommandTape) MoveLeft() error {
	if ct.commandPointer <= 0 {
		return errors.New("reached start of tape")
	}

	ct.commandPointer--

	return nil
}
