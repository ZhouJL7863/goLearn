package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	//string转换为[]byte,容易取得相应位上的具体值
	bsi := []byte(num1)
	bsj := []byte(num2)

	//temp的长度只能为len(num1)+len(num2) 或len(num1)+len(num2)-1
	temp := make([]int, len(num1)+len(num2))
	for i := 0; i < len(bsi); i++ {
		for j := 0; j < len(bsj); j++ {
			//每个位上的乘积，可以直接存入temp中相应的位置
			temp[i+j+1] += int(bsi[i]-'0') * int(bsj[j]-'0')
		}
	}
	//统一处理进位
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10
		temp[i] = temp[i] % 10
	}
	if temp[0] == 0 {
		temp = temp[1:]
	}
	//转换结果
	//temp 选用为[]int,而不是[]byte,是因为
	//go中，byte的基础结构是Unit8，最大值为255
	//不考虑进位的话，temp会溢出
	res := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		res[i] = '0' + byte(temp[i])
	}

}

//只能说两个字，优秀!!!
