package ch03

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	// Create a listener on a random port.
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Listener created, listening on:", listener.Addr().String())

	done := make(chan struct{}) //initiate a channel for signaling purpose

	// Open a go routine to handle accepting connections.
	go func() {
		// Defer the done signal to indicate when the goroutine is done.
		defer func() {
			t.Log("Listener goroutine exiting.")
			done <- struct{}{}
		}()

		// Keep listening for incoming client connections.
		t.Log("Listener started accepting connections.")
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Log("Error accepting connection:", err)
				return
			}
			t.Log("Accepted a new connection.")

			// Open a connection in a new goroutine to handle this specific connection.
			go func(c net.Conn) {
				defer func() {
					t.Log("Closing connection.")
					c.Close()
					done <- struct{}{}
				}()

				// Buffer to intercept handshake and read data
				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error("Error reading from connection:", err)
						}
						return
					}
					t.Logf("Received data: %q", buf[:n])
				}
			}(conn)
		}
	}()

	// Dial to the listener before accepting any connections.
	t.Log("Dialing the server...")
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Dial successful, connected to server.")

	// Closing the connection from the client side.
	t.Log("Closing client connection.")
	conn.Close()

	// Wait for the listener goroutine to finish.
	<-done
	t.Log("Listener closed.")

	// Close the listener after the client connection is closed.
	listener.Close()

	// Wait for the listener goroutine to completely finish.
	<-done
	t.Log("Test completed.")
}
