package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// 测试http
// 1.通过使用假的Request/Response
// 2.通过起服务器

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h apiHandle
		code int
		message string
	}{
		{errPanic, 200, ""},
	}

	for _, test := range tests {
		f := errWrapper(test.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8888/list/fib.txt", nil)
		f(response, request)
		verifyResponse(response.Result(), test.code, test.message, t)
	}
}

func TestErrWrapperServer(t *testing.T) {
	tests := []struct {
		h apiHandle
		code int
		message string
	}{
		{errPanic, 200, ""},
	}

	for _, test := range tests {
		f := errWrapper(test.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		verifyResponse(response, test.code, test.message, t)
	}
}

func verifyResponse(response *http.Response, expectCode int, expectMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectCode || body != expectMsg  {
		t.Errorf("expect (%d, %s); got (%d, %s)", expectCode, expectMsg, response.StatusCode, body)
	}
}