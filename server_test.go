package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "hello\n" {
		t.Errorf("expected hello got %v", string(data))
	}
}

func TestHi(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/hi", nil)
	w := httptest.NewRecorder()

	hi(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "Hi\n" {
		t.Errorf("expected Hi got %v", string(data))
	}
}

func TestDefaultresp(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	defaultresp(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "Well then hello, \"/\"" {
		t.Errorf("expected Well then hello, \"/\" got %v", string(data))
	}
}

func TestHealth(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	health(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if strings.Contains(string(data), "uptime") == false {
		t.Errorf("expected hello got %v", string(data))
	}
}
