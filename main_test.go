package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	first      int
	last       int
	pair       []int
	desiredCar int
	desiredCdr int
}

var td []Test

func TestTheQuestion(t *testing.T) {
	td = append(td, Test{
		3,
		4,
		[]int{3, 4},
		3,
		4,
	})

	for _, v := range td {
		b := theQuestion{v.first, v.last}

		assert.Equal(
			t,
			v.pair,
			theQuestion.Cons(b),
			"Expected %v, got %v\n", v.pair, theQuestion.Cons(b),
		)

		if theQuestion.Car(b) != v.desiredCar {
			t.Logf("Expected %v, got %v\n", v.desiredCar, theQuestion.Car(b))
			t.Fail()
		}

		if theQuestion.Cdr(b) != v.desiredCdr {
			t.Logf("Expected %v, got %v\n", v.desiredCdr, theQuestion.Cdr(b))
			t.Fail()
		}

		assert.Equal(
			t,
			v.pair,
			theQuestion.Cons(theQuestion{v.first, v.last}),
			"Expected %v, got %v\n", v.pair, theQuestion.Cons(theQuestion{v.first, v.last}),
		)

	}
}
