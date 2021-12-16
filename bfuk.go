package bfuk

import (
	"io"
)

type Bfuk struct {
	tape        *Tape
	commandTape *CommandTape
}

func (bf *Bfuk) handleCommand(inputReader io.ByteReader, outputWriter io.ByteWriter) error {
	tape := bf.tape

	cmd := bf.commandTape.GetCurrentCommand()

	switch cmd {
	case '>':
		tape.MoveRight()
	case '<':
		tape.MoveLeft()
	case ',':
		b, err := inputReader.ReadByte()
		if err != nil {
			return err
		}

		tape.SetCurrentCell(uint8(b))
	case '.':
		err := outputWriter.WriteByte(byte(tape.GetCurrentCell()))

		if err != nil {
			return err
		}
	case '+':
		tape.IncrementCurrentCell()
	case '-':
		tape.DecrementCurrentCell()
	case '[':
		if tape.GetCurrentCell() == 0 {
			for bf.commandTape.GetCurrentCommand() != ']' {
				err := bf.commandTape.MoveRight()

				if err != nil {
					return err
				}
			}
		}
	case ']':
		if tape.GetCurrentCell() == 0 {
			for bf.commandTape.GetCurrentCommand() != '[' {
				err := bf.commandTape.MoveLeft()

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func NewBfuk(commandReader io.RuneReader) *Bfuk {
	return &Bfuk{
		commandTape: NewCommandTape(commandReader),
		tape:        NewTape(),
	}
}

func (bf *Bfuk) Run(inputReader io.ByteReader, outputWriter io.ByteWriter) error {
	for {
		err := bf.commandTape.MoveRight()

		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		bf.handleCommand(inputReader, outputWriter)
	}
}
