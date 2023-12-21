package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

// Page 分页
func Page(Limit, Page int) (limit, offset int) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

// Sort 排序
// 默认 created_at desc
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}
	return sort
}

const TimeLayout = "2006-01-02 15:04:05"

var (
	Local = time.FixedZone("CST", 8*3600)
)

func GetNow() string {
	now := time.Now().In(Local).Format(TimeLayout)
	return now
}

func TimeFormat(s string) string {
	result, err := time.ParseInLocation(TimeLayout, s, time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return result.In(Local).Format(TimeLayout)

}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

// CreateToken 创建token
func CreateToken(userId string, expiresAt time.Time) string {
	// 生成随机字符串
	tokenBytes := make([]byte, 32)
	rand.Read(tokenBytes)

	// 将随机字符串编码为 base64
	token := base64.StdEncoding.EncodeToString(tokenBytes)

	// 设置过期时间
	expiresAtInSecs := expiresAt.Unix()

	// 生成 token 字符串
	tokenStr := fmt.Sprintf("%s:%d", token, expiresAtInSecs)

	return tokenStr
}

func ArrayToString(values []string) string {
	return strings.Join(values, ",")
}

func StringToArray(value string) []string {
	return strings.Split(value, ",")
}
