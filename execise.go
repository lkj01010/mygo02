package main
import (
	"time"
	"math/rand"
	"net/http"
	"fmt"
	"strconv"
	"golang.org/x/net/websocket"
//	"mygo02/agent2"
	"mygo02/chat2"
"mygo02/agent2"
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
var s_id = 0
//func sendRecvServer(ws *websocket.Conn){
//	accessCnt ++
//	s_id ++
//	fmt.Printf("agent come, access cnt=%s\n", strconv.Itoa(accessCnt))
//
//	id := `user` + strconv.Itoa(s_id)
//
//	rw := &wsReadWriter{ws}
//	a := chat.NewRoomAgent(agent.NewAgent(rw, id), chatroom)
//
//	a.Serve()
//
//	accessCnt --
//	fmt.Printf("agent leave, access cnt=%s\n", strconv.Itoa(accessCnt))
//	chatroom.RemoveUser(a)
//}

func sendRecvServer2(ws *websocket.Conn){
	accessCnt ++
	s_id ++
	fmt.Printf("agent come, access cnt=%s\n", strconv.Itoa(accessCnt))

	id := `user` + strconv.Itoa(s_id)

	rw := &wsReadWriter{ws}
	a := chat.NewRoomAgent(agent.NewAgent(rw, id), chatroom)

	fmt.Printf("sendRecvServer2::2, a=%#v", a)
//	chatroom.ChaterAdd <- a
	fmt.Println("sendRecvServer2::3")
	a.Serve()


	accessCnt --
	fmt.Printf("agent leave, access cnt=%s\n", strconv.Itoa(accessCnt))
//	chatroom.RemoveUser(a)
//	chatroom.ChaterRM <- id
}

var chatroom = chat.NewRoom()

func main() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	a := [10]int{1, 1, 3, 4, 5, 6, 7, 8, 9, 0}
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

	go chatroom.Serve()

	http.Handle("/", websocket.Handler(sendRecvServer2))
	http.ListenAndServe(":8000", nil)
}

