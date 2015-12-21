package main
import (


//	"golang.org/x/net/websocket"
	"time"
	"math/rand"
	"net/http"
//	"strconv"
	"fmt"
	"strconv"
)

var accessCnt = 0
func rootHandler(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello"))
	accessCnt ++
	fmt.Println(req.Body, strconv.Itoa(accessCnt))
}

func main() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)


	fmt.Print("123")
	fmt.Print(`123`)


//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//		s := http.Server{Handler: websocket.Handler(wsHandler)}
//		s.ServeHTTP(w, req)
//	})
//
//	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
//	if err != nil {
//		panic("Error: " + err.Error())
//	}

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8000", nil)
}

