package main

import "fmt"

//判断s是否是有效的括号的函数
func isValid(s string) bool {
	s_len := len(s)
	//创建用于暂时存储字符的栈
	stack := []byte{}
	for i := 0; i < s_len; i++ { //循环遍历字符串
		switch s[i] {
		//当为左侧括号时直接压栈
		case '[':
			stack = append(stack, s[i])
		case '{':
			stack = append(stack, s[i])
		case '(':
			stack = append(stack, s[i])
		case ']':
			if len(stack) > 0 { //当前栈不空才能取出内容
				top := stack[len(stack)-1] //取出栈顶元素
				if top == '[' {            //如果栈顶元素与之匹配
					stack = stack[:len(stack)-1] //就出栈
				} else { //否则说明括号不匹配返回错误信息
					return false
				}
			} else { //栈是空的但是有右括号说明不符合条件
				return false //返回错误信息
			}

		case '}': //同理
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if top == '{' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		case ')': //同理
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if top == '(' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 { //最后如果栈空了说明都匹配上了
		return true //直接返回正确信息
	} else { //栈如果不空说明有没匹配上的
		return false //返回错误信息
	}
}
func main() {
	s := "]"
	isV := isValid(s)
	fmt.Println(isV)
}
