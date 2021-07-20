package polymorphic

import "fmt"

//1、
//type Live interface {
//	Run()
//	Eat()
//}
//
//type People struct {
//	Name string
//}
//
//func (p *People) Run() {
//	fmt.Println(p.Name,"在路步")
//}
//
//func (p *People) Eat() {
//	fmt.Println(p.Name,"在吃饭")
//}

//2、
type Live interface {
	Run()
}

type People struct {
}

type Animate struct {
}

func (p *People) Run() {
	fmt.Println("有人在路步")
}

func (p *Animate) Run() {
	fmt.Println("动物在奔跑")
}


func AllRun(live Live){
	live.Run()
}
