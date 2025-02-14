# tcp_server
How to implement a tcp server in go, let's play with tcp.
this repoistory is just a warming up to unserstand the mechanisme of go by using channels and go routines and that by implementing tcp server under multiple test files.

#TestDial
net package from go has a rich tools for networking developemnt, under this function we have used net.Dial() and net.Listen().

1. We have to create a listener on a port; in our case we prefer to make random port, the Listen(<connection_type>, <address_ip:port>) return two interfaces, listener and err.

```
// Create a listener on a random port.
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Listener created, listening on:", listener.Addr().String())
```
