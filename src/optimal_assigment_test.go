package src

import (
	"testing"

	"github.com/matryer/is"
)

func TestOptimalAssigment(t *testing.T) {

	is := is.New(t)

	is.Equal(OptimalAssignments([]string{"(5,4,2)", "(12,4,3)", "(3,4,13)"}),
		"(1-3)(2-2)(3-1)")

	is.Equal(OptimalAssignments([]string{"(1,1,4)", "(5,2,1)", "(1,5,2)"}),
		"(1-2)(2-3)(3-1)")

	is.Equal(OptimalAssignments([]string{"(1,4)", "(5,18)"}),
		"(1-2)(2-1)")

	is.Equal(OptimalAssignments([]string{"(13,2,7,1)", "(1,2,3,4)", "(6,7,2,3)", "(67,4,8,18)"}), "(1-4)(2-1)(3-3)(4-2)")

	is.Equal(OptimalAssignments([]string{"(90,75,75,80,82)", "(35,85,20,50,41)", "(40,2,24,1,67)", "(4,70,2,11,23)", "(23,25,56,44,21)"}),
		"(1-2)(2-3)(3-4)(4-1)(5-5)")

}
