 package chat

import (
	"fmt"
	"mygo02/agent2"
)

type Chater struct {
	RW agent.ReadWriter
	Id string
}

type Room struct {
	agents map[string] *chatAgent
	Broadcast chan string
	ChaterAdd chan *chatAgent
	ChaterRM chan string
}

func NewRoom() (*Room){
	return &Room{
		map[string]*chatAgent{},
		make(chan string),
		make(chan *chatAgent),
		make(chan string),
	}
}

func(r *Room)Serve(){
	fmt.Println("room:Serve:0")
	for {
		select {
		case c := <- r.ChaterAdd:
			fmt.Println("room:Serve:1")
			r.AddUser(c)
		//		go ra.Serve()
		case id := <- r.ChaterRM:
			fmt.Println("room:Serve:2")
			r.RemoveUser(id)
		case msg := <- r.Broadcast:
			fmt.Println("room:Serve:3")
			r.Receive(msg)
		}
	}
}

func(r *Room)AddUser(u *chatAgent){
	r.agents[u.Id()] = u
}

func(r *Room)RemoveUser(id string){
	delete(r.agents, id)
}

func(r *Room)Receive(msg string){
	for _, v := range r.agents{
		if err := v.Write(msg); err != nil{
			fmt.Println(err)
//			r.RemoveUser(v.Id())
			r.ChaterRM <- v.Id()
		}
	}
}



