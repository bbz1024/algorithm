package main

import "net/http"

func main() {
	//fmt.Println(a & a)
	// 启动一个服务8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", nil)
}
