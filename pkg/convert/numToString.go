package convert

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func NumToStr[T Number](num T) string {
	return fmt.Sprint(num)
}
