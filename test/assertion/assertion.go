package assertion

import "fmt"

func Assertion(i interface{})  {
	result ,ok:= i.(int)
	if ok{
		fmt.Println("你输入的",result,"是正确的数字")
	}else{
		fmt.Println("请输入正确的数字")
	}
	fmt.Println("aaaaaaaaaaaaa")

}
