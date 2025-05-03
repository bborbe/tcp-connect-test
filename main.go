package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	libservice "github.com/bborbe/service"
)

func main() {
	app := &application{}
	os.Exit(libservice.MainCmd(context.Background(), app))
}

type application struct {
	Target string `required:"true" arg:"target" env:"TARGET" usage:"target to connect to"`
	Amount int    `required:"true" arg:"amount" env:"AMOUNT" usage:"number of tcp connections" default:"1"`
}

func (a *application) Run(ctx context.Context) error {

	conns := make([]net.Conn, 0, a.Amount)
	for i := 0; i < a.Amount; i++ {
		conn, err := net.DialTimeout("tcp", a.Target, 5*time.Second)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Connection %d failed: %v\n", i+1, err)
			continue
		}
		fmt.Printf("Connection %d established\n", i+1)
		conns = append(conns, conn)
	}

	fmt.Printf("Successfully opened %d connections. Press Enter to close them.\n", len(conns))
	fmt.Scanln()

	for i, conn := range conns {
		conn.Close()
		fmt.Printf("Connection %d closed\n", i+1)
	}

	return nil
}
