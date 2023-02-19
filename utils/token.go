package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// BuildToken 根据user_id和username生成token
func BuildToken(id int64, username string) string {
	token := fmt.Sprintf("%v", id) + "$" + username
	return token
}

// ParseToken 解析token从而获取user_id
func ParseToken(token string) int64 {
	parseToken := strings.Split(token, "$")
	id, err := strconv.ParseInt(parseToken[0], 10, 64)
	ResolveError(err)
	return id
}
