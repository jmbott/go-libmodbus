package main

import "fmt"

/*
#include <stdio.h>
#include <stdlib.h>

*/
import "C"


func one() {
	fmt.println("Hello World in Go")
}

func two() {
	C.printf("Hello, World in C\n");
}

