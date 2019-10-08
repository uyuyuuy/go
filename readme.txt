
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

字符串的截取
string[0:1]
string[2:]


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


2.切片（长度和容量，不允许创建容量比长度小的切片（会报错）） https://www.cnblogs.com/f-ck-need-u/p/9854932.html
虽然slice实际上包含了3个属性，它的数据结构类似于[3/5]0xc42003df10，但仍可以将slice看作一种指针。这个特性直接体现在函数参数传值上。

my_slice := make([]int,3,5)
println(my_slice)      // [3/5]0xc42003df10


slice := make([]string,5)
slice := make([]int,3,5) //长度为3，容量为5

slice := []int{1,2,3,4}
空切片 var slice []int
切片的访问和修改值跟数组一样，但是增加元素必须用append，不能slice[] = 1或者 slice[10] = 10，否则报错 panic: runtime error: index out of range
slice := make([]int, 10)    //将切片变成数组，定义了切片的数量

由于多个slice共享同一个底层数组，所以当修改了某个slice中的元素时，其它包含该元素的slice也会随之改变，因为slice只是一个指向底层数组的指针(只不过这个指针不纯粹，多了两个额外的属性length和capacity)，
实际上修改的是底层数组的值，而底层数组是被共享的。
slice := []int{1,2,3,4,5,6,7,8,9}
slice1 := slice[0:1:3]

slice1 := slice[A:B:C]
其中A表示从 slice 的第几个元素开始切，B控制切片的长度(B-A)，C控制切片的容量(C-A)，如果没有给定C，则表示切到底层数组的最尾部。截取时"左闭右开",跟php数组的截取类似


// 创建长度和容量为100的slice，并为第100个元素赋值为3
slice := []int{99:3}

切片扩容
slice = append(slice, 10)

由于string的本质是[]byte，所以string可以append到byte slice中：
s1 := []byte("Hello")
s2 := append(s1, "World"...)
fmt.Println(string(s2))   // 输出：HelloWorld

合并切片
slice1 := []int{1,2,3,4}
slice2 := []int{5,6,7}
slice3 := append(slice1, slice2...)     //主要slice2后面的...

对切片进行迭代，slice是一个集合，所以可以进行迭代。
s1 := []int{11,22,33,44}
for index,value := range s1 {
    println("index:",index," , ","value",value)
}


切片作为参数传递给函数
当切片作为参数传给函数时，其实是传递的副本还是指向该切片的底层数组，函数内部对改切片进行修改会直接影响到底层数组，除非对改切片扩容（扩容会新创建一个底层数组）

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)
	modify(slice)
	fmt.Println(slice)
}

func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}

0xc420082060
0xc420082080
[1 10 3 4 5]
仔细看，这两个切片的地址不一样，所以可以确认切片在函数间传递是复制的。而我们修改一个索引的值后，发现原切片的值也被修改了，说明它们共用一个底层数组。


切片与内存浪费
由于slice的底层是数组，很可能数组很大，但slice所取的元素数量却很小，这就导致数组占用的绝大多数空间是被浪费的。
特别地，垃圾回收器(GC)不会回收正在被引用的对象，当一个函数直接返回指向底层数组的slice时，这个底层数组将不会随函数退出而被回收，而是因为slice的引用而永远保留，除非返回的slice也消失。
因此，当函数的返回值是一个指向底层数组的数据结构时(如slice)，应当在函数内部将slice拷贝一份保存到一个使用自己底层数组的新slice中，并返回这个新的slice。这样函数一退出，
原来那个体积较大的底层数组就会被回收，保留在内存中的是小的slice。

多维切片
slice := [][]int{{10},{20,30}}

内置函数：
len() 返回切片的长度
cap() 返回切片的容量

高级部分
当同一个底层数组有很多slice的时候，一切将变得混乱不堪，因为我们不可能记住谁在共享它，通过修改某个slice的元素时，将也会影响那些可能我们不想影响的slice。
所以，需要一种特性，保证各个slice的底层数组互不影响，相关内容见下面的"扩容"。
当slice的length已经等于capacity的时候，再使用append()给slice追加元素，会自动扩展底层数组的长度。
底层数组扩展时，会生成一个新的底层数组。所以旧底层数组仍然会被旧slice引用，新slice和旧slice不再共享同一个底层数组。



3.字典（映射）
用于存储一系列无序的键值对，映射基于键来存储值
dict := make(map[string]int){"red":1,"blue":2}
dict := map[string]int{}
dict := map[string][]int{}
dict2 := map[string]interface{}{}

判断字典的某个键是否存在
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

判断结构体为空
if objectA== (structname{}){ // your code }

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


make、new、cap
简单的说，new只分配内存并返回类型的指针，make用于slice，map，和chan的初始化，返回的是传入的类型
make只能构建slice、map和channel这3种结构的数据对象，因为它们都指向底层数据结构，都需要先为底层数据结构分配好内存并初始化。
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

匿名函数
闭包内取外部函数的参数的时候是取的地址,而不是调用闭包时刻的参数值
所以我们在使用go func的时候最好把可能改变的值通过值传递的方式传入到闭包之中,避免在协程运行的时候参数值改变导致结果不可预期
func main()  {
    i := 1
    go func(i int) {
        time.Sleep(100*time.Millisecond)
        fmt.Println("i =", i)
    } (i)

    i++
    time.Sleep(1000*time.Millisecond)
}





方法
func (m mm) add(a int, b string)(c int, err error){

}

*/

==========================================================
包
包名都采用小写，和目录名一致，多个单词用 - 拼接
package main 包中的main函数，是项目入口文件，相当于index.php文件（main函数只能出现在package main 中）
go run main.go

引入包，该包根目录所有go文件中的init函数会依次触发，按照文件名排序依次执行(如果根目录下引入了其他包，则会去到引入的包根目录下执行所有go文件中的init函数，一直这样递归执行)



