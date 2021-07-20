package test


func Demo(){

	//fmt.Println("---------------------------------封装---------------------------------------")
	////封装
	//p := new(encapsulation.People)
	//p.SetName("longzhoufeng")
	//fmt.Println(p.GetName())
	//
	//p.SetAge(210)
	//
	//if p.GetAge() > 200 {
	//	fmt.Println(p.GetAge())
	//} else {
	//	fmt.Println("年龄不符合规范")
	//}
	//fmt.Println("---------------------------------继承---------------------------------------")
	////继承
	//teacher := inherit.Teacher{"302班级", inherit.People{Name: "longzhoufeng", Age: 30}}
	//fmt.Println(teacher.Classroom, teacher.Name, teacher.Age)
	//
	//fmt.Println("---------------------------------接口---------------------------------------")
	////接口
	//peo := interfaces.People{Name: "张三", Age: 30}
	//peo.Run(30)
	//peo.Eat()
	////2、继承
	//animate := interfaces.Animate{}
	//animate.Eat()
	//
	//fmt.Println("---------------------------------多态---------------------------------------")
	////people := polymorphic.People{"李四"}
	////var live polymorphic.Live = &polymorphic.People{"李四"}
	////live.Run()
	////live.Eat()
	//
	////2、
	//p2 := &polymorphic.People{}
	//polymorphic.AllRun(p2)
	//
	//p3 := &polymorphic.Animate{}
	//polymorphic.AllRun(p3)
	//
	//
	//fmt.Println("---------------------------------断言---------------------------------------")
	//assertion.Assertion(1234)
}
