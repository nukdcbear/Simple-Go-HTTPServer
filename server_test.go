package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	if strings.Contains(string(data), "Hello") == false {
		t.Errorf("expected hello got %v", string(data))
	}
}

func TestHelloWorld(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/world", nil)
	w := httptest.NewRecorder()

	world(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if strings.Contains(string(data), "Hello world!") == false {
		t.Errorf("expected Hello world got %v", string(data))
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
	if strings.Contains(string(data), "Hi") == false {
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

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	health(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if strings.Contains(string(data), "uptime") == false {
		t.Errorf("expected uptime got %v", string(data))
	}
}

func TestReadConfBadFile(t *testing.T) {

	c, err := readConf("configX.yaml")
	_ = c

	if err == nil {
		t.Errorf("expected error to be not nil")
	}

	if err != nil && !strings.Contains(err.Error(), "no such file") {
		t.Errorf("expected error \"no such file\" got %v", err.Error())
	}
}

func TestReadConf(t *testing.T) {

	c, err := readConf("config.yaml")

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if !strings.Contains(strconv.Itoa(c.Listenport), "3000") {
		t.Errorf("expected Listenport 3000 got %v", c.Listenport)
	}
}
