package chat
import (
	"fmt"
)

type Agent interface {
	Id() (string)
	Read(msg *string)(error)
	Write(msg string)(error)
}

type chatAgent struct{
	a Agent
	r *Room
}

func NewRoomAgent(a Agent, r *Room)(ra *chatAgent){
	ra =  &chatAgent{a, r}
//	r.AddUser(ra)
	return
}

func(ra *chatAgent)Id()(string){
	return ra.a.Id()
}

func(ra *chatAgent)Write(msg string)(error){
	return ra.a.Write(msg)
}

func (ra *chatAgent)handle(msg string)(err error){
//	reply := ra.Id() + ":<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;" + msg
	reply := ra.Id() + ":" + msg

//	ra.r.Receive(reply)
	ra.r.Broadcast <- reply
	return
}

func (ra *chatAgent)Serve(){
	ra.r.ChaterAdd <- ra
	for {
		var buf string
		fmt.Println("Serve Read")
		err := ra.a.Read(&buf)
		if err != nil {
			fmt.Println("chatagent:Serve:readerror:")
			fmt.Println(err)
			ra.r.ChaterRM <- ra.Id()
			break
		}
		fmt.Println("id=" + ra.Id() + ", read=" + buf)
		err = ra.handle(buf)
		if err != nil{
			fmt.Println("chatagent:Serve:handleerror:")
			fmt.Println(err)
			break
		}
	}
}
