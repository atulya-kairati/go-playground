package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    addr, err := net.ResolveUDPAddr("udp", ":50500")
    if err != nil {
        fmt.Println("Error resolving UDP address:", err)
        return
    }
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Error listening on UDP:", err)
        return
    }

	fmt.Println("Listening...")
    var count int
    // var avg int
    buffer := make([]byte, 1024*64)
    for {
        size, _, err := conn.ReadFromUDP(buffer) // blocking function
        if err != nil {
            fmt.Println("Error reading UDP packet:", err)
            continue
        }
        count++
        // avg = (size + avg) / count
        // if count >= 5000{
        // 	timestamp := time.Now().Unix()
        // 	fmt.Printf("%d packets recieved with avg size %d bytes at %d sec\n", count, size, timestamp)
        // 	count = 0
        // 	avg = 0
        // }

        timestamp := time.Now().Unix()
        fmt.Printf("Received packet %d with %d bytes at %d ms\n", count, size, timestamp)
    }
    
    defer conn.Close()
}
