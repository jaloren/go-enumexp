package enumcar

import (
	"fmt"
	"golang.org/x/exp/maps"
)

const (
	Honda = iota
	Toyota
	GM
	unknown = "UNKNOWN"
)

var values = map[Brand]car{
	Honda:  {brand: Honda},
	Toyota: {brand: Toyota},
	GM:     {brand: GM},
}

type Car interface {
	Brand() string
	Equal(enum Car) bool
	String() string
	internal()
}

type Brand int

func (c Brand) String() string {
	switch c {
	case Honda:
		return "Honda"
	case Toyota:
		return "Toyota"
	case GM:
		return "GM"
	default:
		return unknown
	}
}

func (c Brand) Equal(input string) bool {
	return c.String() == input
}

func New(brand Brand) Car {
	val, ok := values[brand]
	if !ok {
		panic(fmt.Sprintf("enum value %d is invalid", brand))
	}
	return val
}

func Values() []car {
	return maps.Values(values)
}

type car struct {
	brand Brand
}

func (c car) Brand() string {
	return c.brand.String()
}

func (c car) String() string {
	return c.brand.String()
}

func (c car) Equal(enum Car) bool {
	return enum.Brand() == c.Brand()
}

func (c car) internal() {
	// no-op
}
