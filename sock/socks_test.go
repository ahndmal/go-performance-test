package sock

import (
	"log"
	"net"
	"testing"
)

func TestSocketSimple(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:25")
	if err != nil {
		t.Fatal(err)
	}

	req := "GET / HTTP" +
		"Sec-WebSocket-Protocol: tcp" +
		"Upgrade: websocket" +
		"Connection: Upgrade"

	_, err2 := conn.Write([]byte(req))
	if err2 != nil {
		log.Println(err2)
	}

	conn.Read([]byte(""))
}
