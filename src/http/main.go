package main

import (
	"fmt"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Keep-Alive", "timeout=3")
	http.SetCookie(w, &http.Cookie{
		Name:     "lzh",
		Value:    "haha",
		Expires:  time.Now().Add(time.Hour*24*365 + 8*time.Hour),
		MaxAge:   10,
		//HttpOnly: true,
		//SameSite: http.SameSiteStrictMode,
	})
	w.Write([]byte(r.URL.Query().Get("q")))
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", home)
	server.HandleFunc("/home", home)
	err := http.ListenAndServe(":1234", server)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("服务启动成功")
	}
}
