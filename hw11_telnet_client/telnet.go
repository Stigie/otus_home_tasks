package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
	Start() error
}

type Client struct {
	address string
	conn    net.Conn
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

var connectionRefuseErr = errors.New("...Connection refused")
var connectionCloseByPeerErr = errors.New("...Connection closed by peer")

func (tc *Client) Connect() error {
	conn, err := net.DialTimeout("tcp", tc.address, tc.timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot connect: %v", err)

		return err
	}

	fmt.Fprintf(os.Stderr, "...Connected to %s\n", tc.address)

	tc.conn = conn

	return nil
}

func (tc *Client) Close() error {
	if tc.conn == nil {
		return nil
	}
	err := tc.conn.Close()
	if err != nil {
		return connectionRefuseErr
	}

	return nil
}

func (tc *Client) Send() error {
	_, err := bufio.NewWriter(tc.conn).ReadFrom(tc.in)
	if err == io.EOF {
		return connectionCloseByPeerErr
	}

	return err
}

func (tc *Client) Receive() error {
	_, err := io.Copy(tc.out, tc.conn)

	return err
}

func (tc *Client) Start() error {
	err := tc.Connect()
	if err != nil {
		return err
	}
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)

	errChan := make(chan error, 1)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := tc.Receive()
				if err != nil {
					cancel()
					errChan <- err

					return
				}
			}
		}
	}(ctxWithCancel)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := tc.Send()
				if err != nil {
					cancel()
					errChan <- err
					return
				}
			}
		}
	}(ctxWithCancel)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ctxWithCancel.Done():
			err := <-errChan
			fmt.Fprintf(os.Stderr, "...%s\n", err)
			tc.Close()
			os.Exit(1)
		case <-termChan:
			fmt.Fprintln(os.Stderr, "...Closing the connection")
			tc.Close()
			os.Exit(1)
		}
	}
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{address: address, timeout: timeout, in: in, out: out}
}

// Place your code here
// P.S. Author's solution takes no more than 50 lines
