package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	TestContains()
	TestCount()
	TestIndex()
	TestIndexFunc()
	TestLastIndex()
	TestLastIndexFunc()
	res := GetFileSuffix("abc.xyz.lmn.jpg")
	fmt.Println(res)
}

//判断是否包含子串
func TestContains() {
	fmt.Println(strings.Contains("dsadsfda", "sad"))
	fmt.Println(strings.Contains("dadaf", "web"))
	fmt.Println(strings.Contains("dgsaudai", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("stevnen大家哦哦", "大"))
}

//判断字符串是否包含另一字符串中的任一字符
func TestContainsAny() {
	fmt.Println(strings.ContainsAny("test", "i"))
	fmt.Println(strings.ContainsAny("failure", "u&i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

//判断字符串是否包含unicode码值
func TestContainsRune() {
	fmt.Println(strings.ContainsRune("一丁", '丁'))
	fmt.Println(strings.ContainsRune("一丁", 19969))
}

//返回字符串中包含另一字符串的个数
func TestCount() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("one", ""))
}

//判断字符串s是否有前缀字符串
func TestHasPrefix() {
	fmt.Println(strings.HasPrefix("1000phone newa", "1000"))
	fmt.Println(strings.HasPrefix("1000phone news", "1000a"))
}

//判断字符串是否有后缀字符串
func TestHasSuffix() {
	fmt.Println(strings.HasSuffix("1000phone news", "news"))
	fmt.Println(strings.HasSuffix("1000phone news", "new"))
}

//返回字符串中另一字符串首次出现的位置
func TestIndex() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
}

//返回字符串中的任一unicode码值首次出现的位置
func TestIndexAny() {
	fmt.Println(strings.IndexAny("andaABV123", "教育基地A"))
}

//返回字符串中字符首次出现位置
func TestIndexByte() {
	fmt.Println(strings.IndexByte("123acz", 'a'))
}

//判断字符串中是否包含unicode码值
func TestIndexRune() {
	fmt.Println(strings.IndexRune("abcABD120", 'A'))
	fmt.Println(strings.IndexRune("It教育培训", '教'))
}

//返回字符串中满足函数f(r)==true字符首次出现的位置
func TestIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello123,中国", f))
}

//返回字符串中子串最后一次出现的位置
func TestLastIndex() {
	fmt.Println(strings.LastIndex("Steven learn english", "e"))
	fmt.Println(strings.Index("go goer", "go"))
	fmt.Println(strings.LastIndex("go goer", "go"))
	fmt.Println(strings.LastIndex("go gier", "ad"))
}

//返回字符串中任意一个unicode码值最后一次出现的位置
func TestLastIndexAny() {
	fmt.Println(strings.LastIndexAny("chicken", "aeipuo"))
	fmt.Println(strings.LastIndexAny("crwth", "aeiouy"))
}

//返回字符串中字符最后一次出现的位置
func TestLastIndexByte() {
	fmt.Println(strings.LastIndexByte("abcAdfAdasvA", 'A'))
}

//返回字符串中满足函数f(r)==true字符最后一次出现的位置
func TestLastIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.LastIndexFunc("Hello 世界", f))
	fmt.Println(strings.LastIndexFunc("Hello, world中国人", f))
}

//获取文件后缀
func GetFileSuffix(str string) string {
	arr := strings.Split(str, ".")
	fmt.Printf("the arr value is", arr)
	return arr[len(arr)-1]
}
