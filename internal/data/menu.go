package data

import (
	"ludum-dare-54/pkg/typeface"
)

type MenuItem struct {
	Text *typeface.Text
	Func func()
}

var Starting bool
