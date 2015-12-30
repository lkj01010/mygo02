package main
import (
	"time"
	"math/rand"
	"net/http"
	"fmt"
	"strconv"
	"golang.org/x/net/websocket"
	"mygo02/agent"
	"mygo02/chat"
)


type wsReadWriter struct {
//	agent.ReadWriter
	ws *websocket.Conn
}

func (w *wsReadWriter)Read(msg *string) (err error){
	err = websocket.Message.Receive(w.ws, msg)
	fmt.Printf("recv:%#v\n", *msg)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (w *wsReadWriter)Write(msg string)(err error){
	err = websocket.Message.Send(w.ws, msg)
	fmt.Printf("send:%#v\n", msg)
	if err != nil{
		fmt.Println(err)
	}
	return
}

var accessCnt = 0
func sendRecvServer(ws *websocket.Conn){
	accessCnt ++
	fmt.Printf("sendRecvBinaryServer finished, access cnt=%s\n", strconv.Itoa(accessCnt))

	id := `user` + strconv.Itoa(accessCnt)

	rw := &wsReadWriter{ws}
	a := chat.NewRoomAgent(agent.NewAgent(rw, id), chatroom)

	for {
		var buf string
		err := rw.Read(&buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		err = a.Handle(buf)
		if err != nil{
			fmt.Println(err)
			break
		}
	}
}

var chatroom = chat.NewRoom()

func main() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sa := a[2:7]
	fmt.Println(sa)
	sa = append(sa, 100)
	sb := sa[3:8]
	sb[0] = 99
	fmt.Println(a)  //输出：[1 2 3 4 5 99 7 100 9 0]
	fmt.Println(sa) //输出：[3 4 5 99 7 100]
	fmt.Println(sb) //输出：[99 7 100 9 0]


//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//		s := http.Server{Handler: websocket.Handler(wsHandler)}
//		s.ServeHTTP(w, req)
//	})
//
//	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
//	if err != nil {
//		panic("Error: " + err.Error())
//	}

	http.Handle("/", websocket.Handler(sendRecvServer))
	http.ListenAndServe(":8000", nil)
}

