package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// 路由的本质是路径到处理函数的映射
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/", snippet)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/view1", snippetView1)
	mux.HandleFunc("/snippet/json", snippetCreateJSON)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/create1", snippetCreate1)
	mux.HandleFunc("/snippet/create2", snippetCreate2)

	log.Print("serve on 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// 根
func home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	// 上面这个是将非/的请求打回去
	w.Write([]byte("这是/，你可能写错了"))
}

func snippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/snippet/" {
		w.Write([]byte("现在是/snippet/"))
		return
	}
	http.NotFound(w, r)
}

// 专门加了了一层这个，目的是为了展示下面这个点
// 固定路径与子树路径
// 固定路径(fixed path) 比如/snippet/view ，特点就是必须要完全匹配
// 子树路径(subtree path) 当访问/snippet/view/ ，会去返回去找最长的或者说上一级，这个会返回/snippet/，
// 当这个都没有的时候会返回/

// 0为初始，1加上了读取url里面参数
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly a specific snippet..."))
}

func snippetView1(w http.ResponseWriter, r *http.Request) { //让那个这个父亲url里面的参数
	// 查询字符串中id参数的值，并尝试用strconv.Atoi()函数将其转化为整数，
	// 如果失败或者转化完小于1，就返回404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "显示一个具有指定 id 的文本片段 %d..", id)
}

// 1将0里面直接写http的那些请求头改为http包自带的方法，2将1里面直接的状态码和请求方法替换为预定义的 HTTP 状态码常量
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { //r.Method表示当前请求方法,比如GET,POST,PUT,DELETE
		w.Header().Set("Allow", "POST")       //在调用w.WriteHeader()或w.Write()后再修改响应头映射就不会有改变了
		w.WriteHeader(405)                    //给响应写状态码,响应只能有效写一次，
		w.Write([]byte("Method Not Allowed")) //给响应写具体内容，要先写响应头再写body，否则头是默认的200
		return
	}
	w.Write([]byte("creat a new snippet..."))
}

func snippetCreate1(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("creat a new snippet..."))
}

func snippetCreate2(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) //预定义的 HTTP 状态码常量
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//如果你没手动设置 Content-Type,第一次 Write() 时，Go 会根据最开始写出的数据自动推断 Content-Type
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("creat a new snippet..."))
}

// 这个是为了解决当前go里面的JSON 可能被当成纯文本
type snippetResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

func snippetCreateJSON(w http.ResponseWriter, r *http.Request) {
	resp := snippetResponse{
		Message: "snippet created successfully",
		ID:      123,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	// json.NewEncoder(w)是一个函数调用,返回的是一个 *json.Encoder,
	// Encode(resp) 是这个返回对象的方法调用
}
