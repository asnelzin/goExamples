package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	ip := os.Args[1]

	const (
		PORTS_NUM = 65535 + 1
		BULKSIZE  = 20000
	)

	var wg sync.WaitGroup
	var start, end int
	for bulk := 1; end < PORTS_NUM; bulk++ {
		result := make(chan string)
		start = end
		end = Min(bulk*BULKSIZE, PORTS_NUM)
		for port := start; port < end; port++ {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 50*time.Millisecond)
				if err == nil {
					result <- fmt.Sprintf("%d port is open", port)
					conn.Close()
				}
			}(port)
		}

		go func() {
			for item := range result {
				fmt.Println(item)
			}
		}()

		wg.Wait()
	}
}
