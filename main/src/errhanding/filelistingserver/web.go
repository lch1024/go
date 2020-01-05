package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"net/http"
	"os"
)

func fserver (writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

type apiHandle func(http.ResponseWriter, *http.Request) error

func errWrapper(handler apiHandle) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request){
		defer func() {
			if r := recover(); r != nil {
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(fserver))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}