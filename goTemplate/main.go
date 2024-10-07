package main

import (
	"fmt"
	"go-demo/entity"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe fail!")
		return
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {

	kua := func(str string) string {
		return str + "真帅!"
	}

	// New中的名字，必须和文件名字相同，否则会解析不到。
	t := template.New("hello.gohtml")

	t.Funcs(map[string]any{
		"kua": kua,
	})

	// 嵌套模板使用时，要把嵌套的模板一起放入解析，并且要放在母模板后面
	files, err := t.ParseFiles("./goTemplate/hello.gohtml", "./goTemplate/ul.gohtml")
	if err != nil {
		fmt.Printf("ParseFiles fail! err :%v", err)
		return
	}

	// 1. 使用结构体
	localBook := entity.LocalBook{
		Username: "John",
		Gender:   "M",
		Age:      12,
	}

	// 2. 使用map也可以
	localBookM := map[string]interface{}{
		"username": "Lucy",
		"gender":   "W",
		"age":      22,
	}

	// 3. 数组
	hobbies := []string{
		"basketball",
		"football",
		"rap",
	}

	// 组合到一个大的map中
	field := map[string]interface{}{
		"bs": localBook,
		"bm": localBookM,
		"hb": hobbies,
	}

	err = files.Execute(writer, field)
	if err != nil {
		fmt.Printf("files.Execute fail! err :%v", err)
		return
	}

}
