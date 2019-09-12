package main

import (
	"fmt"
)

func main() {
	var First float32
	var FuHao string
	var Two float32
	var exit string
	
	for {
	fmt.Print("输入1为计算器，输入2为退出:")
	fmt.Scanln(&exit)
	fmt.Print("\n")	
	if exit == "1"{
		fmt.Print("请输入第一个数字:")
		fmt.Scanln(&First)
		

		fmt.Print("请输入运算符号:")
		fmt.Scanln(&FuHao)

		fmt.Print("请输入第二个数字:")
		fmt.Scanln(&Two)
		if FuHao == "+" {
			var JieGuo float32 = First + Two
			fmt.Print("计算结果为：",JieGuo)
			fmt.Print("\n")
			
		}
		if FuHao == "-" {
			var JieGuo float32 = First - Two
			fmt.Print("计算结果为：",JieGuo)
			fmt.Print("\n")
		}
		if FuHao == "*" {
			var JieGuo float32 = First * Two
			fmt.Print("计算结果为：",JieGuo)
			fmt.Print("\n")
		}
		if FuHao == "/" {
			var JieGuo float32 = First / Two
			fmt.Print("计算结果为：",JieGuo)
			fmt.Print("\n")

			
			}
		
		}
	if exit == "2"{
		fmt.Print("已为您退出")
		break

		}	
	}

		
	
}
