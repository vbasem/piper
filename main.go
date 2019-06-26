package main



import "github.com/fatih/color"


func main() {
	println("hello")
	GetConsole("blue")("Prints text in blue.")

}



func GetConsole(c string) func(string, ...interface {}) {

	if c == "blue" {
		return color.Blue
	} else {
		return color.Yellow
	}

}