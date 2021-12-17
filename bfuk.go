package bfuk

import (
	"bufio"
	"io"
)

type Bfuk struct {
	tape                *Tape
	commandTape         *CommandTape
	loopStartIndexStack stack
}

func (bf *Bfuk) pushCurrentPointerToStack() {
	bf.loopStartIndexStack = bf.loopStartIndexStack.Push(bf.commandTape.commandPointer)
}

func (bf *Bfuk) handleCommand(inputReader *bufio.Reader, outputWriter *bufio.Writer) error {
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
		outputWriter.Flush()
		if err != nil {
			return err
		}
	case '+':
		tape.IncrementCurrentCell()
	case '-':
		tape.DecrementCurrentCell()
	case '[':
		bf.pushCurrentPointerToStack()

		if tape.GetCurrentCell() == 0 {
			ptr := bf.commandTape.commandPointer

			for {
				err := bf.commandTape.MoveRight()
				if err != nil {
					return err
				}

				if bf.commandTape.GetCurrentCommand() == '[' {
					bf.pushCurrentPointerToStack()
				}

				if bf.commandTape.GetCurrentCommand() == ']' {
					top := bf.loopStartIndexStack.Top()

					bf.loopStartIndexStack, _ = bf.loopStartIndexStack.Pop()

					if top == ptr {
						break
					}
				}
			}
		}

	case ']':
		if tape.GetCurrentCell() != 0 {
			bf.commandTape.commandPointer = bf.loopStartIndexStack[len(bf.loopStartIndexStack)-1]
		} else {
			bf.loopStartIndexStack, _ = bf.loopStartIndexStack.Pop()
		}
	}

	return nil
}

func NewBfuk(commandReader io.RuneReader) *Bfuk {
	return &Bfuk{
		commandTape:         NewCommandTape(commandReader),
		tape:                NewTape(),
		loopStartIndexStack: make(stack, 0),
	}
}

func (bf *Bfuk) Run(inputReader *bufio.Reader, outputWriter *bufio.Writer) error {
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
