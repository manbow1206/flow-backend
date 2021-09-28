package main

import (
	// "encoding/json"
	// "encoding/json"
	"encoding/json"
	f "fmt" // 別名
	"log"
	"os"
	"reflect"
	// "golang.org/x/text/unicode/norm"
)

// type
type ID int
type Priority int

type Task struct {
	ID     int
	Detail string
	done   bool
}

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
	var RangeCheck [4]string

	RangeCheck[0] = "a"
	RangeCheck[1] = "b"
	RangeCheck[2] = "c"
	RangeCheck[3] = "d"

	f.Println("---- Slice range ----")
	for i, s := range RangeCheck {
		f.Println(i, s)
	}

	// 値の切り出し
	f.Println("---- Slice 値の切り出し ----")
	slice3 := []int64{0, 1, 2, 3, 4, 5, 6}
	f.Println(slice3[0:1])
	f.Println(slice3[0:2])
	f.Println(slice3[0:3])
	f.Println(slice3[:5])
	f.Println(slice3[5:])
	f.Println(slice3[:])
	f.Println(slice3[0:len(slice3)])

	f.Println(sum1(1, 2, 3))

	// Map
	f.Println("------------- Map syntax -------------")
	// 宣言と初期化
	f.Println("---- Map declaration  ----")
	var month map[int]string = map[int]string{}

	month[1] = "January"
	month[2] = "February"
	f.Println(month)
	CheckType(month)

	month1 := map[int]string{
		1: "January",
		2: "February",
	}

	f.Println(month1)
	CheckType(month1)

	// 操作
	f.Println("---- Map operation  ----")
	jan := month[1]
	f.Println(jan)

	test, ok := month[1]
	f.Println(test)
	if ok {
		f.Println("OK")
	}

	// 削除
	f.Println("---- Map delete  ----")
	delete(month, 1)
	f.Println(month)

	f.Println("---- Map for syntax  ----")
	for key, value := range month1 {
		f.Println("%d %s\n", key, value)
	}

	// ポインター
	f.Println("------------- Pointer syntax -------------")

	var PointerInt int = 10
	callByValue(PointerInt)
	f.Println(PointerInt)

	callByRef(&PointerInt)
	f.Println(PointerInt)

	// defer
	f.Println("------------- Defer syntax -------------")
	// file, err := os.Open("./error.go")
	// if err != nil {
	//     // エラー処理
	// }
	// defer file.Close()

	// パニック
	f.Println("------------- Panic syntax -------------")
	// エラーを戻り値として表現できない場合や、回復不可能なシステムエラーで大域脱出が必要な場合に使用
	// 通常は関数の戻り値として、呼び出し側に返す
	// defer func() {
	//     err := recover() // パニックが起きた際のエラーを取得
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	// }()

	// a := []int{1,2,3,4}
	// f.Println(a) //パニックが発生

	// panic()
	// a := []int{1,2,3}
	// for i := 0; i < 10; i++ {
	//     panic(errors,New("Index out of range"))
	// }
	// f.Println(a[i])

	// type
	f.Println("------------- Type syntax -------------")

	var id ID = 3
	var priority Priority = 5

	// カスタム型指定通りなのでOK
	ProcessTask(id, priority)

	// カスタム型指定通りではないのでNG
	// ProcessTask(priority, id)

	// struct
	f.Println("------------- Struct syntax -------------")
	// Rubyのクラスに近い役割

	// 構造体型の宣言
	// 大文字で始まる場合はpublic、小文字で始まる場合はprivate(モジュール内)
	type Task struct {
		ID     int
		Detail string
		done   bool
	}

	// 通常の値を生成
	f.Println("---- Struct 通常使用 ----")
	var task Task = Task{
		ID:     1,
		Detail: "まんぼう",
		done:   true,
	}
	f.Println(task.ID)
	f.Println(task.Detail)
	f.Println(task.done)

	// フィールド名省略
	f.Println("---- Struct フィールド名省略 ----")
	var task1 Task = Task{
		2, "ジンベイザメ", false,
	}

	f.Println(task1.ID)
	f.Println(task1.Detail)
	f.Println(task1.done)

	// ゼロ値
	f.Println("---- Struct ゼロ値 ----")
	var task2 Task = Task{}
	f.Println(task2.ID)
	f.Println(task2.Detail)
	f.Println(task2.done)

	// ポイント型
	f.Println("---- Struct Pointer ----")
	// var task3 Task = Task{}
	// var task3 *Task = &Task{}

	// var task3 Task = Task{done :true}
	// Finish(task3)
	// f.Println(task3.done) //falseのまま

	// 構造体の名前の前に&を付けると、構造体の値ではなく、メモリアドレスが変数に格納される。ポインタ型。
	// var task4 Task = &Task{done :true}
	// Finish(task4)
	// f.Println(task4.done) //true

	// コンストラクタ
	f.Println("---- Struct Constructer ----")
	// コンストラクタにあたる構文がないので、Newで始まる関数を定義し、その内部で構造体を生成するのが通例。
	// func NewTask(id int, detail string) *Task {
	// 	task := &Task{
	// 		ID: id,
	// 		Detail: detail,
	// 		done: false,
	// 	}
	// 	return task
	// }

	// task := NewTask(1, "コンストマンボウ")
	// f.Println("%+v", task)

	// メソッド
	f.Println("---- Struct Method ----")
	// Task Structに紐づいた関数(メソッド)
	// 	func (task Task) String() string {
	//   str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	//   return str
	// }

	// interface
	f.Println("------------- Interface syntax -------------")
	// 宣言
	// typr Stringer interface {
	// 	String() string
	// }

	// 実装
	// func Print(stringer Stringer) {
	// 	f.Println(stringer.String())
	// }

	// Print(task)

	// interface{}
	// どんな型でも格納できる特殊な型。int,string,boolなどと同じ、golangの型名。
	// どんな型でも格納できるので、受け取りが分からない場合に活用可能。
	// 演算はできない
	var interfaceX, interfaceY interface{}

	f.Printf("%#v", interfaceX) // -> nilinterfaceY
	interfaceX = 1
	interfaceX = 2.1
	interfaceY = []int{1, 2, 3}
	interfaceY = "hello"
	interfaceY = 2
	f.Printf("%#v", interfaceY) // -> nilinterfaceY

	// // 型の埋め込み
	f.Println("------------- Embed syntax -------------")

	// // TaskにUserを埋め込む
	// // User 構造体
	// type User struct {
	// 	FirstName string
	// 	LastName string
	// }

	// func (u *User) FullName() string {
	// 	fullname := f.Sprintf(
	// 		"%s %s", u.FirstName, u.LastName)
	// 	return fullname
	// }

	// func NewUser(firatName, lastName) *User {
	// 	return &User {
	// 		FirstName: sirstName,
	// 		LastName: lastName,
	// 	}
	// }

	// // Task 構造体
	// type Task struct {
	// 	ID int
	//   Detail string
	//   done bool
	//   *User // Userを埋め込む
	// }

	// func NewTask(id int, detail, firstName, lastName string) *Task{
	// 	task := Task{
	// 		ID: id,
	// 		Detail: detail,
	// 		done: falss,
	// 		User: NewUser(firstName, lastName)
	// 	}
	// 	return task
	// }

	// taskEmbed := NewTask(1, "まんぼう", "まんぼう", "太郎")
	// f.Println(task.FirstName) //task構造体からUser構造体のFrisNameを取得
	// f.Println(task.lastName) //task構造体からUser構造体のLastNameを取得
	// f.Println(task.FullName()) //task構造体からUser構造体のFullName()メソッドえを取得
	// f.Println(task.User) //task構造体からUser構造体そのものを取得

	// インターフェース自体の埋め込み
	f.Println("---- Interface  in Interface ----")
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	type ReadWriter interface {
		Reader
		Writer
	}

	// Type conversion
	f.Println("------------- Type conversion -------------")

	// キャスト
	f.Println("---- Type conversion Cast ----")
	var castInt8 uint8 = 3
	var castInt32 = uint32(castInt8)
	f.Println(castInt32)

	var castString string = "abc"
	var castByte []byte = []byte(castString)
	f.Println(castByte)

	// NG
	// a := int("a")

	// // Type Assertion
	f.Println("---- Type conversion Type Assertion----")
	// func Print(value interface{}) {
	// 	s, ok := value.(string)
	//    if ok {
	//     fmt.Printf("value is string: %s\n", s)
	// } else {
	//     fmt.Printf("value is not string\n")
	// }
	// }

	// Print("abc")
	// Print(12)

	// 			type Stringer interface {
	//     String() string
	// }

	// 	func Print(value interface{}) {
	//     switch v := value.(type) {
	//     case string:
	//         fmt.Printf("value is string: %s\n", v)
	//     case int:
	//         fmt.Printf("value is int: %d\n", v)
	//     case Stringer:
	//         fmt.Printf("value is Stringer: %s\n", v)
	//     }
	// }

	// JSONパッケージ
	f.Println("------------- JSON パッケージ -------------")
	type Person struct {
		ID      int8
		Name    string
		Email   string
		Age     int
		Address string
		memo    string
	}

	// person := &Person{
	// 	        ID: 1,
	//       Name: "Gopher",
	//       Email: "gopher@example.org",
	//       Age: 5,
	//       Address: "",
	//       memo: "golang lover",
	// }

	// b, err := json.Marshal(person)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// f.Println(string(b))

	// タグ付け
	f.Println("---- JSON Tags ----")
	// type Person struct {
	// ID:int `json:"id"`
	//   Name string `json:"name"`
	//   Email string `json:"-"`
	//   Age int `json:"age"`
	//   Address string
	//   memo string
	// }

	// Unmarshal()
	f.Println("---- Unmarshal ----")
	// var person Person
	// b := []byte(`{"id":1,"name":"Gopher","age":5}`)
	// err := json.Unmarshal(b, &person)
	// タグと紐づいて、構造体に代入
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(person) // {1 Gopher 5 }

	// os,ioパッケージ
	f.Println("------------- os,io syntax -------------")
	// ファイルの生成
	f.Println("---- os,io ファイル作成----")
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// ファイルへの書き込み
	f.Println("---- os,io ファイルへの書き込み----")
	fileWritten := []byte("hello world\n")

	_, err = file.Write(fileWritten)
	if err != nil {
		log.Fatal(err)
	}

	// ファイルからの読み出し
	// f.Println("---- os,io ファイルからの読み出し----")
	// openFile, err := os.Open("./file.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer openFile.Close()

	// fileReading := make([]byte, 12)

	// _, err1 := file.Read(fileReading)
	// if err1 != nil {
	// 	log.Fatal(err)
	// }

	// f.Print(string(fileReading))

	// JSONファイルの読み出し
	f.Println("---- os,io JSONファイルからの読み出し----")

	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}

	jsonFile, err3 := os.Create("./person.json")

	if err3 != nil {
		log.Fatal(err3)
	}

	defer file.Close()

	encoder := json.NewEncoder(jsonFile)

	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
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
func sum1(nums ...int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}

// Ponter
func callByValue(i int) {
	i = 20
}

func callByRef(i *int) {
	*i = 20
}

//type
func ProcessTask(id ID, priority Priority) {
}

// 引数にポイント型を指定
func Finish(task *Task) {
	task.done = true
}

// メソッド
func (task Task) String() string {
	str := f.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}
