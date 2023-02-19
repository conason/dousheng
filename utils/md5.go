package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(password string) string {
	h := md5.New()
	_, err := io.WriteString(h, password)
	ResolveError(err)
	md5Password := string(h.Sum([]byte(nil)))
	md5Password = fmt.Sprintf("%x",md5Password)
	return md5Password
}
