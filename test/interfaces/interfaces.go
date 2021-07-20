package interfaces

import "fmt"

type People struct {
	Name string
	Age int
}

type Animate struct {

}

type Eat interface {
	Eat()
}


type Live interface {
	Run(num int)
	Eat
}


func(p *People) Run(num int){
	fmt.Println(p.Name,"正在跑步，跑了",num,"米")
}

func (p *People) Eat() {
	fmt.Println(p.Name,"正在吃钣")
}

func (p *Animate) Eat() {
	fmt.Println("动物正在吃东西")
}
