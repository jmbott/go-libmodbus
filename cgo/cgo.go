package main

import (
        "fmt"
        "go-libmodbus/hey_go"
        "go-libmodbus/modbus"
)

func main() {
        hey_go.Hey()
        fmt.Println("I am here in go")
	modbus.Read_holding_registers_tcp()
}
