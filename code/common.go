package code

import "errors"

const DefaultProgressBarLength = 40.0

var InvalidSyntax = errors.New("invalid file structure")

const (
	DoneSymbol     = "X"
	UndoneSymbol   = " "
	ProgressSymbol = "#"
)
