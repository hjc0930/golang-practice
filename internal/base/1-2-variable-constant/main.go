package main

import "fmt"

/*
*

  - 变量

  - 常量

  - 基本数据类型 (int, bool, string, float etc...)

  - bool 1byte true/false

  - int8 1byte -128 ~ 127

  - int16 2byte -32768 ~ 32767

  - int32 4byte -2^31 ~ 2^31-1

  - int64 8byte -2^63 ~ 2^63-1

  - int 根据操作系统决定 32位系统为4byte 64位系统为8byte

  - uint8 1byte 0 ~ 255

  - uint16 2byte 0 ~ 65535

  - uint32 4byte 0 ~ 2^32 - 1

  - uint64 8byte 0 ~ 2^64 - 1

  - uint 根据操作系统决定 32位系统为4byte 64位系统为8byte

  - uintptr 根据操作系统决定 32位系统为4byte 64位系统为8byte 用于存放指针

  - byte 1byte 0 ~ 255 等同于uint8

  - rune 4byte -2^31 ~ 2^31-1 等同于int32 表示unicode码点

  - float32 4byte IEEE-754 32位浮点数

  - float64 8byte IEEE-754 64位浮点数

  - complex64 8byte 实部和虚部都是float32类型的复数

  - complex128 16byte 实部和虚部都是float64类型的复数

  - string 16byte 字符串

  - 复合数据类型 (数组，切片，map，结构体)
*/
func main() {
	var str string = "hello" // 标准定义
	var str2 = "hello"       // 自动推导类型
	str3 := "world"          // 短变量定义，只能用在函数中

	const pi = 3.14 // 常量，在编译时就确定了
	fmt.Println(str, str2, str3, pi)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

const (
	MonDay = iota + 1
	TueDay
	WedDay
	ThuDay
	FriDay
	SatDay
	SunDay
)

/**
 * 计算平均分
 */
func average(scoreMap map[string]float64) float64 {
	var total float64

	for _, v := range scoreMap {
		total += v
	}

	avg := total / float64(len(scoreMap))

	return avg
}
