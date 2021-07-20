package encapsulation

type People struct {
	name string
	age  int
}

//设置get set方法

func (p *People) SetName(name string) {
	p.name = name
}

func (p *People) GetName() string {
	return p.name
}

func (p *People) SetAge(age int) {
	if age < 1 || age > 200 {
		age = 0
	} else {
		p.age = age
	}
}

func (p *People) GetAge() int {
	return p.age
}
