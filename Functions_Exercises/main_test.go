package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T){
	total:=Add(5,5)
	if total!=10 {
		t.Errorf("Sum was incorrect, got: %d,want: %d.",total,10)
	}
}

func TestAddition(t *testing.T){
	got:=Addition(5,3)
	want:=8
	if got!=want{
		log.Fatalf("Sum was incorrect, got: %v,want: %v.",got,want)
	}
}


func TestParadise(t *testing.T){
	got:=Paradise("Singapore")
	want:="My idea of paradise is Singapore"
	if got!=want{
		log.Fatalf("Sum was incorrect, got: %v,want: %v.",got,want)
	}
}