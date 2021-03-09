/*
 * @Description  :
 * @Author       : jsmjsm
 * @Github       : https://github.com/jsmjsm
 * @Date         : 2021-03-09 10:43:34
 * @LastEditors  : jsmjsm
 * @LastEditTime : 2021-03-09 11:07:39
 * @FilePath     : /goto/goto_v4/main.go
 */
package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	listenAddr = flag.String("http", ":8080", "http listen address")
	dataFile   = flag.String("file", "store.json", "data store file name")
	hostname   = flag.String("host", "localhost:8080", "http host name")
)

var store *URLStore

func main() {
	flag.Parse()
	store = NewURLStore(*dataFile)
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(*listenAddr, nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html") // 教程里面缺少这一行，程序无法正常运行
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://%s/%s", *hostname, key)
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
