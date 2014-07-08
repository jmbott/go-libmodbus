package modbus

import (
        "fmt"
)

/*
The modbus_read_registers() function uses the Modbus function code 0x03 (read holding registers).
This code reads all the available meter addresses.
*/

/*

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>

#include <modbus.h>

int modbus_tcp (void) {
    modbus_t *ctx;
    uint16_t tab_reg[64];
    int rc;
    int i;

    ctx = modbus_new_tcp("192.168.1.35", 502);
    if (modbus_connect(ctx) == -1) {
        fprintf(stderr, "Connection failed: %s\n", modbus_strerror(errno));
        modbus_free(ctx);
        return -1;
    }

    modbus_set_debug(ctx, 0x01);
    rc = modbus_set_slave(ctx, 1);
    if (rc == -1) {
        fprintf(stderr, "%s\n", modbus_strerror(errno));
        return -1;
    }

    // Read 1 register starting at address 25 
    rc = modbus_read_registers(ctx, 24, 1, tab_reg);
    if (rc == -1) {
        fprintf(stderr, "%s\n", modbus_strerror(errno));
        return -1;
    }
    
    for (i=0; i < rc; i++) {
        printf("reg[%d]=%d (0x%X)\n", i, tab_reg[i], tab_reg[i]);
    }
    
    modbus_close(ctx);
    modbus_free(ctx);
    
    return 0;
}

*/
import "C"

func Read_holding_registers_tcp() {
        fmt.Println("Modbus Wrapper")
        C.modbus_tcp()
}

