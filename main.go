package main

import (
	f "fmt" // 別名
	"reflect"
	// "golang.org/x/text/unicode/norm"
)

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

	// エラーを返り値に設定
	// エラー値は最後にするのが慣習、自作エラーパッケージを作成し、それを使用する事も可能
	// file, err := os.Open("hello world")
	// if err != nil {
	//     // エラー処理
	// }
	// fileを用いた処理

	// 名前付き戻り値
	// メリット : 関数の宣言時から戻り値の意味が把握できる。return時の書き間違いなど防止
	// func NamaReturn(i,j int) (result int, err error) {
	//     if j == 0 {
	//         err = errors.Ner("divied by zero")
	//         return //自動的に名前付き返り値が返される。この場合はreturn 0, err
	//     }
	//     result = i / j
	//     return //return result, nil と同じ
	// }

	// 関数リテラル
	// その場で使用する無名関数
	func(i, j int) {
		f.Println(i + j)
	}(2, 4)

	// 変数に関数リテラルを入れる場合
	var sum func(i, j int) = func(i, j int) {
		f.Println(i + j)
	}

	sum(2, 4)

	// 配列 固定長
	f.Println("---Array syntax---")
	var arr [4]string
	arr[0] = "マンボウ"
	arr[1] = "ジンベイザメ"
	arr[2] = "イルカ"

	f.Println(arr)
	f.Println(arr[1])
	f.Println(arr[2])

	// 宣言と同時に配列長を暗黙的に指定
	arr1 := [4]string{"マンボウ", "ジンベイザメ", "いるか", "まぐろ"}
	arr2 := [...]string{"マンボウ", "ジンベイザメ", "いるか", "まぐろ"}

	f.Println("配列長、直接指定")
	f.Println(arr1)
	f.Println("配列長、暗黙指定")
	f.Println(arr2)

	// 配列長も情報として含む
	// func CheckArrayLength(arr [4]string) {
	//     f.Printlna(arr)
	// }
	// var arr3 [4]string
	// var arr4 [4]string
	// CheckArrayLength(arr3) //ok
	// CheckArrayLength(arr4) //NG 長さが違う為

	// 配列が関数に渡された場合、配列はコピーされて、関数内で使用
	// func fn(arr [4]string) {
	//     arr[0] = "x"
	//     fmt.Println(arr) // [x b c d]
	// }

	// func main() {
	//     arr := [4]string{"a", "b", "c", "d"}
	//     fn(arr)
	//     fmt.Println(arr) // [a b c d]
	// }


    // Slice
	f.Println("------------- Slice syntax -------------")
    // 配列の使用は、シビアなメモリ管理が必要なプログラムなので、それ以外の場面では基本的にはスライスを使用

    // var slice []string

    slice1 := []string{"a", "b", "c"}
    f.Println(slice1[0])

    // スライス末尾に追加 append()
    f.Println("---- Slice append() ----")
    var slice2 []string
    f.Println("---- Slice append()前 ----")
    f.Println(slice2)
    CheckType(slice2)

    slice2 = append(slice2, "マンボウ")
    slice2 = append(slice2, "ジンベイザメ")
    slice2 = append(slice2, "まぐろ")
    slice2 = append(slice2, "いるか")

    f.Println("---- Slice append()後 ----")
    f.Println(slice2)
    CheckType(slice2)

    // rang
    var RangeCheck[4]string

    RangeCheck[0] = "a"
    RangeCheck[1] = "b"
    RangeCheck[2] = "c"
    RangeCheck[3] = "d"

    f.Println("---- Slice range ----")
    for i, s :=  range RangeCheck {
        f.Println(i, s)
    }

    // 値の切り出し
    f.Println("---- Slice 値の切り出し ----")
    slice3 := []int64{0,1,2,3,4,5,6}
    f.Println(slice3[0:1])
    f.Println(slice3[0:2])
    f.Println(slice3[0:3])
    f.Println(slice3[:5])
    f.Println(slice3[5:])
    f.Println(slice3[:])
    f.Println(slice3[0:len(slice3)])


    f.Println(sum1(1,2,3))
}

// Check Type
func CheckType(t interface{}) {
    f.Println(reflect.TypeOf(t))
}

// function
func RetunrsFunc(i, j int) (int, int) {
	i = i + 1
	j = j + 1
	return i, j
}


//  slice
    func sum1(nums ...int)(result int) {
        for _, n := range nums {
            result += n
        }
        return
    }