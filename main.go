package main

import (
	f "fmt" // 別名
)

// var message string = "変数を使用"

func main() {

	// variable
	f.Println("---variable syntax---")
	// var message string = "変数を使用"

	// 複数変数宣言
	// var foo, bar, buz = "foo", "bar", "buz"
	var (
		foo = "foo"
		bar = "bar"
		buz = "buz"
	)
	f.Println(foo)
	f.Println(bar)
	f.Println(buz)

	// 省略変数
	message := "略変数を使用"
	f.Println(message)

	// 定数
	const Constant string = "定数を使用"
	f.Println(Constant)

	// # ゼロ値
	// ## 整数型
	var Zero int
	f.Println(Zero)

	// ## 浮動小数型
	var Float float64
	f.Println(Float)

	// ## 真偽型
	var Bool bool
	f.Println(Bool)

	// ## 文字列
	var Str string
	f.Println("※↓stringゼロ値")
	f.Println(Str)

	// if
	f.Println("---if syntax---")
	a, b := 10, 100
	if a > b {
		f.Println("a is larger than b")
	} else if a < b {
		f.Println("a is smaller than b")
	} else {
		f.Println("a equals b")
	}

	// NG if 省略
	// if n == 10
	//     f.Println(n)

	//  NG 三項演算子
	// n == 10 ? yes : no

	// for
	f.Println("---for syntax---")
	for i := 0; i < 10; i++ {
		f.Println(i)
	}

	// for for while
	n := 0
	for n < 10 {
		f.Println(n)
		n++
	}

	// infinite loop
	// for {
	//     doSmething()
	// }

	// break, continue
	m := 0
	for {
		m++
		if m > 10 {
			break
		}
		if m%2 == 0 {
			continue
		}
		f.Println(m)
	}

	// switch
	f.Println("---switch syntax---")
	SwitchValiable := 0

	switch SwitchValiable {
	case 15:
		f.Println("FizzBuzz")
	case 5, 10:
		f.Println("Fizz")
	case 3, 6, 9:
		f.Println("Buzz")
	default:
		f.Println(n)
	}

	// switch fallthrough
	SwitchFallthrough := 3

	switch SwitchFallthrough {
	case 3:
		SwitchFallthrough = SwitchFallthrough - 1
		fallthrough
	case 2:
		SwitchFallthrough = SwitchFallthrough - 1
		fallthrough
	case 1:
		SwitchFallthrough = SwitchFallthrough - 1
		f.Println(SwitchFallthrough)
	default:
		f.Println(SwitchFallthrough)
	}

	// switch if Ver
	SwitchIf := 10

	switch {
	case SwitchIf%15 == 0:
		f.Println("FizzBuzz")
	case SwitchIf%5 == 0:
		f.Println("Fizz")
	case SwitchIf%3 == 0:
		f.Println("Buzz")
	default:
		f.Println(SwitchIf)
	}

	// // function
	f.Println("---function syntax---")

	// // basic
	// func hello() {
	//     f.Println("Hello")
	// }

	// arg Ver
	// func Arg(i, j int) {
	//     sum i + j
	//     f.Println(sum)
	// }

	// return Ver
	// func RetunrFunc(i, j int) int {
	//     sum i + j
	//     return sum
	// }

	// returns Ver
    // 戻り値の変数を個数分用意
	x, y := 3, 4
	x, y = RetunrsFunc(x, y)
    f.Println(x, y)

    // NG 戻り値の変数を個数分用意していない
    // cannot assign 2 values to 1 variables
	// x = RetunrsFunc(x, y)


    // 必要な戻り値のみ用意
    x, _ = RetunrsFunc(x, y)
    f.Println(x)

}

// returns Ver
func RetunrsFunc(i, j int) (int, int) {
	i = i + 1
    j = j + 1
    return i, j
}
