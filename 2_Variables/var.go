package main // 声明 main 包
import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
)

func main() { // 声明 main 主函数
	var a = "Hello World"
	fmt.Println(a) // 打印 Hello World!
	// var case = "Hello World 2"  <--编译无法通过，因为case是保留关键字
	// var 1str = "Hello World 2"  <--编译无法通过，因为变量名不能以数字开头

}
