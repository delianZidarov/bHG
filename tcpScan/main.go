package main

import (
	"fmt"
	"net"
)

func main() {
 tooFast()
}

func iterScanner(){
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

func tooFast(){
  for i := 1; i < 1025; i++ {
    go func(j int){
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
