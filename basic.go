package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa=3
var bb="kkk"
//bb:=true 函数外不支持这样写
//不是全局变量，作用域是包内变量
var(
	cc=4
	dd=true
)
//集中定义变量
func variableZeorValue()  {
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}
func variableInitialValue()  {
	var a,b int =3,4
	var s string="abc"
	fmt.Println(a,b,s)
}
func variableTypeDeduction()  {
	var a,b,c,s=3,4,true,"def"
	fmt.Println(a,b,c,s)
}
func variableShorter()  {
	a,b,c,s:=3,4,true,"def"
	b=5
	fmt.Println(a,b,c,s)
}
func euler()  {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}
func triangle()  {
	var a,b int =3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	println(c)
}

func consts()  {
	const(
		filename="abc.txt"
		a,b=3,4
	)
	var c int
	c=int(math.Sqrt(a*a+b*b))
	println(c)
}
func enums(){
	const(
		cpp=iota
		java
		_
		python
		golang
	)
	const(
		b=1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang)
	println(b,kb,mb,gb,tb,pb)
}
func main()  {
	fmt.Println("HelloWolrd")
	variableZeorValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,cc,dd)
	euler()
	triangle()
	consts()
	enums()
}
