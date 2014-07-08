package main

import (
	"fmt"
	"go-libmodbus/hey_go"
)

func main() {
	hey_go.Hey()
	fmt.Println("I am here")
}

/*
https://www.youtube.com/watch?feature=player_embedded&v=XCsL89YtqCs&noredirect=1
http://blog.golang.org/c-go-cgo
http://golang.org/cmd/cgo/
http://www.goinggo.net/2013/08/using-c-dynamic-libraries-in-go-programs.html
http://stackoverflow.com/questions/1760468/interface-go-with-c-libraries
http://tour.golang.org/#1
http://www.golang-book.com/
http://dennisforbes.ca/index.php/2013/07/31/demonstrating-gos-easy-c-interop/
*/
