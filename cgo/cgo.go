package main

import (
        "fmt"
        "go-libmodbus/hey_go"
)

/*
#include <stdio.h>
#include <stdlib.h>

*/
import "C"

func main() {
        hey_go.Hey()
        fmt.Println("I am here in go")
	//C.printf("Hello, World in C\n");
}

