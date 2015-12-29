package agent
import (
	"golang.org/x/net/websocket"
	"io"
)

type User struct {
	Name string
	io.ReadWriter
	ws *websocket.Conn
}

func (u *User) Write(msg []byte) (n int, err error) {
	return 0, nil
}