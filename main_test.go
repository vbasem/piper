package main



import (
	"github.com/fatih/color"
	"testing"
	"reflect"
)


func TestGetConsole_blue(t *testing.T) {

	sf1 := reflect.ValueOf(color.Blue)
	sf2 := reflect.ValueOf(GetConsole("blue"))
	if sf1 != sf2 {
		t.Error("dang it bad")
	}

}


func TestGetConsole_others(t *testing.T) {

	sf1 := reflect.ValueOf(color.Red)
	sf2 := reflect.ValueOf(GetConsole(""))
	if sf1 != sf2 {
		t.Error("dang it bad")
	}

}
