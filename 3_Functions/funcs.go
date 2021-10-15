package main // 声明 main 包
import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
)

//函数的定义
func sum(num1, num2 int) int {
	return num1 + num2
}

//多个返回值的函数
func sumAndComp(num1, num2 int) (int, bool) {
	var nSumRet int = num1 + num2
	var bCompRet bool = false
	if num1 > num2 {
		bCompRet = true
	}
	return nSumRet, bCompRet
}

//变长参数的函数
func sumAll(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	var retSum int = 0
	for _, tmp := range nums { //此处 _符号用于表示忽略索引，因为我们这里不会用到索引值，否则可以使用一个变量代替_ ,并打印出来
		retSum += tmp
	}
	return retSum
}
func main() { // 声明 main 主函数

	//函数的使用
	var result int = sum(2, 3)
	fmt.Printf("两数和为%d\r\n", result) //格式化输出

	retSum, retComp := sumAndComp(2, 3)
	fmt.Printf("两数和为%d\r\n", retSum) //格式化输出
	if retComp {
		fmt.Println("更大")
	} else {
		fmt.Println("更小")
	}
	fmt.Printf("\r\n多个数的和为:%d", sumAll(1, 2, 3, 4, 5, 6))
}
