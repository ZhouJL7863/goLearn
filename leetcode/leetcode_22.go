package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res

}
func dfs(left, right, idx int, bytes []byte, res *[]string) {
	//所有符号都添加完毕
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}
	//"("不用担心匹配问题，
	//只要left > 0 就可以直接添加
	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}
	//想要添加")"时
	//需要left < right
	//即在bytes[:idx]至少有一个"("可以有与之匹配
	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}
func main() {
	var n int = 3
	res := generateParenthesis(n)
	fmt.Println(res)
}

//回溯法三部曲：
//(1):选择，对于每个特定的解，肯定是由一步步构建而来的，而每一步怎么构建，肯定都是有限个选择，要怎么选择，这个要知道
//；同时，在编程时候要定下，优先或合法的每一步选择的顺序，一般都是通过多个if或者for循环来排列
//(2):条件。对于每个特定解的某一步，他必然要符合某个解要求符合的条件，如果不符合，就要回溯，起始回溯也就是递归调用的返回
//(3):结束：当达到一个特定结束条件时候，就认为这个一步步构建的解是符合要求的解了，把解存下来或者打印下来。对于这一步来说，
//有时候也可以另外写一个issolution函数来判断，注意，当达到第三步时，有时候需要构建一个数据结构，把符合要求的解存起来，便于
//当得到所有解时，把解空间输出来，这个数据结构必须是全局的，作为参数之一传递给递归函数。
