package field

import (
	"errors"
	"math/big"
)

type NumberType struct {
	Minimum *big.Rat
	Maximum *big.Rat
	Step    big.Rat
}

func (n NumberType) NewValue() Value {
	return &Number{}
}

func (n NumberType) Validate() error {
	if n.Minimum != nil && n.Maximum != nil && n.Maximum.Cmp(n.Minimum) <= 0 {
		errors.New("Maximum cannot be less than or equal to minimum")
	}

	return nil
}

type Number struct {
	value big.Rat
}

func (n *Number) Set(x, y int64) error {
	if y == 0 {
		return errors.New("Denominator cannot be zero")
	}

	n.value.Set(big.NewRat(x, y))
	return nil
}
