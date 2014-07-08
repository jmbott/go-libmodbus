package modbus

import (
        "fmt"
)

/*
#include <stdio.h>
#include <stdlib.h>

void test() {
        printf("Hello World in C\n");
}

*/
import "C"

func Read_holding_registers_tcp() {
        fmt.Println("X I am here in go")
        C.test()
}

