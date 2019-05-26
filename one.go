package main //包名，main表示当前是一个可执行程序，而不是一个库

//import "fmt"

import (
	"fmt"
	"reflect"
	"strings"
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

var string1 string = "string1" // var string1 = 'string1' //会报错，不能使用单引号 invalid character literal (more than one character)

// sting2 := "sting2"  //该写法只能在函数体内使用non-declaration statement outside function body

//执行的优先级比main函数高
func init() {

}

//程序主函数
func main() {
	//局部变量定义了必须使用
	inta = 1

	var int0 int32
	fmt.Println(fmt.Sprintf("%T", int0))
	fmt.Println(getType(int0))

	//复数
	var complex_a complex64
	complex_a = 3.2 + 12i
	complex_b := complex(3.2, 12)
	//获取复数的实部和虚部
	vr := real(complex_a)
	vi := imag(complex_b)
	fmt.Println(getType(vr))
	fmt.Println(getType(vi))


	funcString()
}


func funcString() string {
	//字符串是一个定长的字节数组，字符串是不可改变的UTF-8字符序列
	//字符串的拼接，用 + 拼接
	string0 := "Hello,世界!"
	string0 += "!!!"
	lens := len(string0) //返回字符串的字节数，一个中文有3个字节(包括中文标点符号)
	lens2 := len([]rune(string0)) //这样可以返回字符串的字符数
	fmt.Println(lens)
	fmt.Println(lens2)

	//反单引号和双引号
	//\n   \r   \t

	//字符串类似数组操作截取字符串
	fmt.Println(string0[0:1])  //截取的字符串不包括后面索引的值
	fmt.Println(string0[:2])
	fmt.Println(string0[0:])

	//没有意义
	for i := 0; i < lens; i++ {
		// fmt.Println(string0[i])
	}

	for _, v := range string0 {
		fmt.Printf("%c\n",v)
	}
	fmt.Printf("%d", 120)

	str := "a,b,c,d,e,f,g"
	stringSlice := strings.Fields(str)	//只能用于区分空格 
	fmt.Printf("%q", stringSlice)

	//修改字符串，只能先复制到另外一个可写变量中，一般用[]byte或[]rune类型
	//如果要修改字节，则使用[]byte；如果要修改字符，则使用[]rune



	//1.包含判断
	strings.HasPrefix(str, "find")
	strings.HasSuffix(str, "find")
	strings.Contains(str, "find")
	strings.ContainsAny(str, "find")
	strings.Index(str, "find")
	strings.LastIndex(str, "find")
	strings.Replace(str, "old", "new", -1)
	strings.Count(str, "find")
	strings.ToUpper(str)
	strings.ToLower(str)
	strings.Trim(str,"trimstr")
	strings.TrimLeft(str,"trimstr")
	strings.TrimSpace(str,"trimstr")
	strings.Split(str,",")
	strings.Join(slice,",")



	return	"1"
}

/*
获取变量的类型
*/
func getType(v interface{}) string {
	return fmt.Sprintf("%T", v)
	return reflect.TypeOf(v).String()
}

//变量定义注意
// 只能在函数体内使用
// inta := 1

//基础数据类型
/*
1.整型int/uint
默认int类型
有符号整型：int8,int16,int32(rune),int64
无符号整型：uint8(byte),uint16,uint32,uint64
uintptr


2.浮点型
float32,float64(默认)  分别为4字节和8字节
复数
complex64,complex128(默认)

3.字符串类型
string


4.布尔型
false(默认)、true
*/
