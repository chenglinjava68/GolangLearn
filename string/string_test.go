package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值“”

	s = "hello"
	t.Log(len(s))

	//s[1] = '3' //string是不可变的byte slice，不可通过此方法修改

	//s = "\xE4\xB8\xA5"  //可以存储任何十六进制数据
	s = "\xE4\xBA\xBB\xFF" //乱码亦可存储
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s)) // 3，是 byte 数
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestUnicode(t *testing.T) {
	s := "中"
	t.Log(len(s)) // 3，是 byte 数

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

//字符串的分割和拼接, 调用 strings 包中的函数
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

//字符串与 int 的转换
func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	//Atoi 函数有两个返回值
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
