package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
)

func main20() {
	path := "d:/"

	fi, err := ioutil.ReadDir(path)
	execute(&path, &fi, &err)
}

func execute(path *string, f *[]os.FileInfo, err *error) {
	if *err != nil {
		return
	}

	for _, file := range *f {

		if file.IsDir() {
			p := *path + file.Name() + "/"
			fi, e := ioutil.ReadDir(p)
			execute(&p, &fi, &e)

		} else {
			print(path, &file)
		}
	}
}

func print(path *string, f *os.FileInfo) {
	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	wr.Write([]byte(*path + (*f).Name()))
	wr.Close()

	fmt.Println(bufs.String())
}
