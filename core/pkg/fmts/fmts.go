package fmts

import (
	"io"
	"log"
	"os"
	"strings"

	gcat "github.com/jeffotoni/gconcat"
)

//Stdout func
func Stdout(strs ...interface{}) {
	str := gcat.Build(strs...)
	_, err := io.Copy(os.Stdout, strings.NewReader(str))
	if err != nil {
		log.Println(err)
	}
}

//Concat contaquena
func Concat(strs ...interface{}) string {
	return gcat.Concat(strs...)
}

//Concat contaquena
func ConcatStr(strs ...string) string {
	return gcat.ConcatStr(strs...)
}
