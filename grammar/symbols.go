package grammar

import (
	"strings"

	set "github.com/sef-computin/go-set"
)

type production string

func NewProd(args ...string) production {
	prod := production(strings.Join(args, " "))

	return prod
}

func (this production) ToSlice() []string {
	return strings.Split(string(this), " ")
}

func (this production) ToString() string {
	return string(this)
}

type productions set.Set

func NewProds() productions {
	return set.NewSet()
}
