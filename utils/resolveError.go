package utils

import "log"

// ResolveError 处理错误
func ResolveError(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}
