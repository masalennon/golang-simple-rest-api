package main

import (
	"net/http"
	"fmt"
	"encoding/json"

)

func main() {
	// mux := http.NewServeMux() //デフォルトのマルチプレクサを生成するコード、リクエストをハンドラにリダイレクトする。
    // mux.HandleFunc("/", homeLink)
    // mux.HandleFunc("/events", getAllEvents)
    // mux.Handle("/events/", http.StripPrefix("/events/", http.HandlerFunc(getOneEvent)))

	http.HandleFunc("/trust", handleTrustRequest) //第一引数のURLへのリクエストが到着すると、第二引数に指定されているハンドラ関数にリダイレクトされる。全てのハンドラ関数が第一引数にResponseWriterをとり、第二引数にRequestへのポインタを取るのでハンドラ関数に改めて引数を渡す必要はない
	//ハンドラ関数とは第一引数にResponseWriterをとり、第二引数にRequestへのポインタを取るGoの関数に過ぎない
	//HTTPリクエスト発生時に呼び出される具体的な処理が書かれています。
	http.HandleFunc("/belief", handleBeliefRequest)
    http.ListenAndServe(":8080", nil) //なぜnilでいい？
}

func handleTrustRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleTrustGet(w, r)
	case "POST":
		
	}
}

func handleBeliefRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleBeliefGet(w, r)
	case "POST":
		
	}
}

type trust struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type belief struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type allTrusts []trust

type allBeliefs []belief

var trusts = allTrusts{
	{
		ID:          "11",
		Title:       "trust",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

var beliefs = allBeliefs{
	{
		ID:          "12",
		Title:       "belief",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func handleTrustGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	fmt.Println(r.URL.Path)
	json.NewEncoder(w).Encode(trusts) //JSONにエンコードしたtrustsをレスポンスボディ（w）に書き込む
}

func handleBeliefGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	json.NewEncoder(w).Encode(beliefs)
}

func handlePost(w http.ResponseWriter, r *http.Request) {

}