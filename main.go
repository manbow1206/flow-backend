package main

import (
	f "fmt" // 別名
)

// var message string = "変数を使用"

func main() {

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
	f.Println(Str)

    // if
    a, b := 10 ,100
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

}
