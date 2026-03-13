package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	w.Write([]byte("hello from snippeetbox"))
}

func snippet(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path == "/snippet/" {
	// 	w.Write([]byte("Display all snippets..."))
	// 	return
	// }

	// http.NotFound(w, r)
	w.Write([]byte("hello12313123 from snippeetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly a specific snippet..."))
}

func snippetView111(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly other specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly a specific snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/", snippet)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/view111", snippetView111)
	mux.HandleFunc("/snippet/creat", snippetCreate)

	log.Print("serve on 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// 路由的本质是路径到处理函数的映射

// 固定路径与子树路径
// 固定路径(fixed path) 比如/snippet/view ，特点就是必须要完全匹配
// 子树路径(subtree path) 当访问/snippet/view/ ，会去返回去找最长的或者说上一级，这个会返回/snippet/，
// 当这个都没有的时候会返回/
