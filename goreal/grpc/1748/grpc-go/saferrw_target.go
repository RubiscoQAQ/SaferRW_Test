package grpc

import (
	"golang.org/x/net/http2"
	"io"
	"net"
	"time"
)

func main() {
	mctBkp := minConnectTimeout
	// Call this only after transportMonitor goroutine has ended.
	defer func() {
		minConnectTimeout = mctBkp
	}()
	//defer leakcheck.Check(t)
	minConnectTimeout = time.Millisecond * 500
	server, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		//t.Fatalf("Error while listening. Err: %v", err)
	}
	defer server.Close()
	done := make(chan struct{})
	clientDone := make(chan struct{})
	go func() { // Launch the server.
		defer func() {
			if done != nil {
				close(done)
			}
		}()
		conn1, err := server.Accept()
		if err != nil {
			//t.Errorf("Error while accepting. Err: %v", err)
			return
		}
		defer conn1.Close()
		// Don't send server settings and make sure the connection is closed.
		time.Sleep(time.Millisecond * 1500) // Since the first backoff is for a second.
		conn1.SetDeadline(time.Now().Add(time.Second))
		b := make([]byte, 24)
		for {
			// Make sure the connection was closed by client.
			_, err = conn1.Read(b)
			if err == nil {
				continue
			}
			if err != io.EOF {
				//t.Errorf(" conn1.Read(_) = _, %v, want _, io.EOF", err)
				return
			}
			break
		}

		conn2, err := server.Accept() // Accept a reconnection request from client.
		if err != nil {
			//t.Errorf("Error while accepting. Err: %v", err)
			return
		}
		defer conn2.Close()
		framer := http2.NewFramer(conn2, conn2)
		if err := framer.WriteSettings(http2.Setting{}); err != nil {
			//t.Errorf("Error while writing settings. Err: %v", err)
			return
		}
		time.Sleep(time.Millisecond * 1500) // Since the first backoff is for a second.
		conn2.SetDeadline(time.Now().Add(time.Millisecond * 500))
		for {
			// Make sure the connection stays open and is closed
			// only by connection timeout.
			_, err = conn2.Read(b)
			if err == nil {
				continue
			}
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				return
			}
			//t.Errorf("Unexpected error while reading. Err: %v, want timeout error", err)
			break
		}
		close(done)
		done = nil
		<-clientDone

	}()
	client, err := Dial(server.Addr().String(), WithInsecure())
	if err != nil {
		//t.Fatalf("Error while dialing. Err: %v", err)
	}
	<-done
	client.Close()
	close(clientDone)
}
