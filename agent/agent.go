package agent
import (
	"fmt"
)

type Writer interface {
//	Read(msg *string) (err error)
	Write(msg string)(err error)
}

type Agent struct {
	r Writer
	id string
}

func NewAgent(r Writer, id string) (a *Agent){
	fmt.Println("NewAgent, id=" + id)
	return &Agent{r, id}
}

func (a *Agent)Id()(string){
	return a.id
}

//func (a *Agent)Serve(){
//	for {
//		var buf string
//		err := a.rw.Read(&buf)
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//		err = a.handle(buf)
//		if err != nil{
//			fmt.Println(err)
//			break
//		}
//	}
//}

func (a *Agent)Handle(msg string)(err error){
	reply := a.id + `:` + msg

	err = a.Write(reply)
	return
}

func (a *Agent)Write(msg string)(error){
	return a.r.Write(msg)
}