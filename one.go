package main //包名，main表示当前是一个可执行程序，而不是一个库

//import "fmt"

import (
	"fmt"
)

//常量 建议全部大写
const A int = 1
const B = 1
const X, Y, Z int = 1, 2, 3 //并行赋值

const (
	U, V, W = 1, 2, 3
)

var inta, intb int 
//inta = 1  // non-declaration statement outside function body

var string1 string = "string1"	// var string1 = 'string1' //会报错，不能使用单引号 invalid character literal (more than one character)

// sting2 := "sting2"  //该写法只能在函数体内使用non-declaration statement outside function body

//执行的优先级比main函数高
func init() {


}

//程序主函数
func main() {
	//局部变量定义了必须使用
	inta = 1
	sting2 := "sting2"	
	fmt.Println("Hello World")
	fmt.Println(sting2)
}





//变量定义注意
// 只能在函数体内使用
// inta := 1 

//基础数据类型
/*
1.整型
有符号整型：int8,int16,int32(rune),int64
无符号整型：uint8(byte),uint16,uint32,uint64
uintptr


2.浮点型
float32,float64(默认)  分别为4字节和8字节
复数
complex64,complex128

3.字符串类型
string


4.布尔型
false、true

