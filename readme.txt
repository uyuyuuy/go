
关键字：25个
包导入与申明：import、package
程序实体申明与定义：var、type、func、interface、map、struct、chan、const
流程控制：if、else、continue、for、return、go、case、goto、switch、select、break、default、defer、fallthrough、range

/*
基础数据类型
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

复合数据类型
1.数组
var array [5]int
array := [5]int{1,2,3,4,5}
array := [...]int{1,2,3,4}
访问和修改元素
a0 := array[0]
array[0] = 0

多维数组：
array := [5][2]int{{1,2},{3,4},{5,6},{7,8},{9,10}}


2.切片（长度和容量，不允许创建容量比长度小的切片（会报错））
slice := make([]string,5)
slice := make([]int,3,5) //长度为3，容量为5

slice := []int{1,2,3,4}
空切片 var slice []int
切片的访问和赋值跟数组一样

创建一个新切片
slice := []int{1,2,3,4,5,6,7,8,9}
slice1 := slice[0:1:3]

切片扩容
slice = append(slice, 10)

多维切片
slice := [][]int{{10},{20,30}}

内置函数：
len() 返回切片的长度
cap() 返回切片的容量

3.映射
用于存储一系列无序的键值对，映射基于键来存储值
dict := make(map[string]int){"red":1,"blue":2}
dict := map[string]int{}
dict := map[string][]int{}
value, exists := dict['red']
if exists {
	fmt.Println(value)
}

删除元素
delete(dict, 'red')

4.结构体 struct{}
自定义类型，由一系列属性组成
	type xiangdong struct {
		name string
		age  int
	}

	people1 := new(xiangdong)
	people1.name = "xd"
	people1.age = 29


5.接口 interface{}
定义了一组方法（方法集），但这些方法不包含实现

流程控制
if a := 1; a = 1 {

} else if 1 > 2 {

} else {

}

switch a {
case 1:
	fmt.Println(a)
case 2:
	fmt.Println(a)
default:
	fmt.Println(a)

}

switch a := 1; a {
case a > 1:
	fmt.Println(a)
case a >2:
	fmt.Println(a)
default:
	fmt.Println(a)

}

select 配合通道的读写操作，用于多个channel的并发读写操作

for a := 0; a < 5; a++ {
	fmt.Println(a)
	if a == 1 {
		break
	}
	if a == 2 {
		continue
	}
}

a := [5]int{1,2,3,4,5};
for  index, value := range a{
	fmt.Println(value)
}

函数
func name(int, int, string) (int, int) {
	...
}


方法
func (m mm) add(a int, b string)(c int, err error){

}

*/