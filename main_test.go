package main

import "testing"

type S struct {
	Name string
}

func TestRun(t *testing.T) {
	var list []*S
	list = append(list, &S{"summer"})
	modify(list)
	t.Log(*list[0])
}

func modify(list []*S) {
	list[0].Name = "sunny"
}
