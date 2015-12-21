package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
)

func httpGetTest(i int){
	resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		// handle error
	}

	defer func(){
		resp.Body.Close()
		if(i==99){
			flag<-0
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
var flag = make(chan interface{})

func main(){
	t1 :=time.Now()
	for i := 0; i < 100; i++{
		go httpGetTest(i)
	}
	<- flag

	t2 := time.Now()
	d := t2.Sub(t1)
	fmt.Println(`time consume: `, d)
}
