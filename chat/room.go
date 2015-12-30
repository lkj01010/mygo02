package chat

import (
	"fmt"
	"sync"
)

type Room struct {
	agents map[string] *chatAgent
	mu sync.Mutex
}

func NewRoom() (*Room){
	return &Room{agents: map[string]*chatAgent{}}
}

func(r *Room)AddUser(u *chatAgent){
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()
	r.agents[u.Id()] = u
}

func(r *Room)RemoveUser(u *chatAgent){
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.agents, u.Id())
}

func(r *Room)Receive(msg string){
	for _, v := range r.agents{
		if err := v.Write(msg); err != nil{
			fmt.Println(err)
			r.RemoveUser(v)
		}
	}
}



