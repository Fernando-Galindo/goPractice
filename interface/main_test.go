package main

import (
	"testing"
)

func TestSquareArea(t *testing.T) {
	s := SQUARE{
		length: 8,
		width:  3,
	}
	a := s.Area()
	if a != 24 {
		t.Errorf("Area returned as %v when expected 24", a)
	}
}
