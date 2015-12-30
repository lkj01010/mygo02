package chat
import (
)

type Agent interface {
	Id() (string)
	Write(msg string)(error)
}

type chatAgent struct{
	a Agent
	r *Room
}

func NewRoomAgent(a Agent, r *Room)(ra *chatAgent){
	ra =  &chatAgent{a, r}
	r.AddUser(ra)
	return
}

func(ra *chatAgent)Id()(string){
	return ra.a.Id()
}

func(ra *chatAgent)Write(msg string)(error){
	return ra.a.Write(msg)
}

func (ra *chatAgent)Handle(msg string)(err error){
	reply := ra.Id() + `:` + msg

//	err = ra.Write(reply)
	ra.r.Receive(reply)
	return
}
