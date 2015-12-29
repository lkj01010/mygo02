package main
import (


//	"golang.org/x/net/websocket"
	"time"
	"math/rand"
	"net/http"
//	"strconv"
	"fmt"
	"strconv"
	"golang.org/x/net/websocket"
	"mygo02/agent"
	"mygo02/chat"
)

var accessCnt = 0
func sendRecvServer(ws *websocket.Conn){
	accessCnt ++

	fmt.Printf("sendRecvServer %#v\n", ws)
	for {
		var buf string
		// Receive receives a binary message from client, since buf is []byte.
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("recv:%#v\n", buf)
		// Send sends a binary message to client, since buf is []byte.
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", buf)
	}
	fmt.Printf("sendRecvBinaryServer finished, access cnt=%s", strconv.Itoa(accessCnt))
}

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

	r := chat.Room{map[string]agent.User{}}

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

