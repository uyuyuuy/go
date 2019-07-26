
关键字：25个
包导入与申明：import、package
程序实体申明与定义：var、type、func、interface、map、struct、chan、const
流程控制：if、else、continue、for、return、go、case、goto、switch、select、break、default、defer、fallthrough、range

变量命名规则：驼峰命名，避免使用下划线

/*
基础数据类型
1.整型int/uint
默认int类型（根据操作系统确定，32位操作系统就是int32,64位就是int64）
有符号整型：int8,int16,int32(rune),int64
无符号整型：uint8(byte),uint16,uint32,uint64
uintptr 一个大到足以存储指针值的未解释位的无符号整数

byte        alias for uint8 （别名）
rune        alias for int32 （别名）

uint8所有无符号8位整数的集合（0到255）
uint16所有无符号16位整数的集合（0到65535）
uint32所有无符号32位整数的集合（0到4294967295）
uint64所有无符号64位整数的集合（0到18446744073709551615）

int8所有带符号的8位整数集（-128到127）
int16所有带符号的16位整数集（-32768到32767）
int32所有带符号的32位整数的集合（-2147483648到2147483647）
int64所有带符号的64位整数的集合（-9223372036854775808到9223372036854775807）


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
切片的访问和修改值跟数组一样，但是增加元素必须用append，不能slice[] = 1或者 slice[10] = 10，否则报错 panic: runtime error: index out of range

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


[]byte

=================================
iota   从0开始递增
const (
	LevelDefault int = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
	LevelSerializable
	LevelLinearizable
)


简单的说，new只分配内存，make用于slice，map，和channel的初始化
数据类型（结构体）不能用range

==========================================
流程控制
GO没有while do相关函数
=========================================
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


============================================================================
函数
func name(int, int, string) (int, int) {
	...
}

func name(int0 int,strings ...string) (int, int) {
    //strings 为切片
	for _,name := range strings {
	    fmt.Println(name)
	}
}

name(1, "a", "b", "c")
slice := []string{"a", "b", "c"}
name(1, slice...)

slice2 := []string{"d", "e", "f"}
slice = append(slice, slice2...)

func append(slice []Type, elems ...Type) []Type

如果一个函数的某个参数是可变参数，可以传入切片类型数据，传入格式为slice...，后面需要加...



方法
func (m mm) add(a int, b string)(c int, err error){

}

*/