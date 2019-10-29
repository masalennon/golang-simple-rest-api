package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetArticles(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder() //レスポンスの情報を格納する
	handler := http.HandlerFunc(getArticles)
	handler.ServeHTTP(rr, req) //ここで↑のgetArticlesを呼んでいる。rrに返ってきた内容が入る。出力して確認済み。
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":"12","title":"Introduction to golang","description":"Go言語の基本"},{"id":"13","title":"Introduction to Algorithm","description":"アルゴリズムの基本"},{"id":"14","title":"Introduction to Programming","description":"プログラミングの基本"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetArticleByID(t *testing.T) {

	req, err := http.NewRequest("GET", "13", nil)
	fmt.Println("req: ", req)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(getArticle)//castしている
	// handler.ServeHTTP(rr, req)//getArticleを呼んでいる
	getArticle(rr, req)
	fmt.Println("rr: ", rr)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	actual := rr.Body.String()
	expected := `{"ID":"13","Title":"Introduction to Algorithm","Description":"アルゴリズムの基本"}`
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

func TestGetArticleNotFound(t *testing.T) {
	req, err := http.NewRequest("POST", "13", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
