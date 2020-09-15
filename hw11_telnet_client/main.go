package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout=time")
}

func main() {
	flag.Parse()

	tail := flag.Args()
	if len(tail) != 2 {
		log.Fatal("Please, provide host and port")
	}
	host, port := tail[0], tail[1]

	flag.Parse()

	tc := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)

	err := tc.Start()
	if err != nil {
		log.Println(err)
	}

	// Place your code here
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
}
