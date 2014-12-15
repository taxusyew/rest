package rest

import (
	"bufio"
	"io"
	"mime/multipart"
	"os"
)

type FormFile struct {
	FileName    string
	ContentType string
	File        multipart.File
	limit       int64
}

func (this *FormFile) Limit(len int64) *FormFile {
	this.limit = len
	return this
}

func (this *FormFile) Save(path string, limit int) int64 {
	defer this.File.Close()
	of, e := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if nil != e {
		panic(e.Error())
	}
	defer of.Close()
	bof := bufio.NewWriterSize(of, 8*1024)
	writeLen, e := io.Copy(bof, this.File)
	if nil != e {
		panic(e.Error())
	}
	return writeLen
}
