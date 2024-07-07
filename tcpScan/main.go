package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	waitGroup()
}

func iterScanner() {
	for i := 1; i < 1025; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port closed or filtered
			continue
		}
		conn.Close()
		fmt.Println(i, "open")
	}
}

func tooFast() {
	for i := 1; i < 1025; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println(j, "open")
		}(i)
	}
}

func waitGroup() {
	var wg sync.WaitGroup
	for i := 1; i < 1025; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Print(".")
				return
			}
			conn.Close()
			fmt.Println("\n", j, "open")
		}(i)
	}
	wg.Wait()
}
