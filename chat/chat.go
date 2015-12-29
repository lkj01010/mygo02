package chat

import (
	"mygo02/agent"
)

type Room struct {
	users map[string]agent.User
}

func(r *Room)AddUser(u *agent.User){
	r.users[u.Name] = u
}

func(r *Room)RemoveUser(u *agent.User){
	delete(r.users, u.Name)
}

func(r *Room)OnReceive(msg string){
	for _, v := range r.users{
		v.Write([]byte(msg))
	}
}



