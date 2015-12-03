package wxpay

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map转xml
func map2Xml(params map[string]string) string {
	xmlString := "<xml>"
	for k, v := range params {
		xmlString += fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xmlString += "</xml>"
	return xmlString
}

// 生成sign
func makeSign(params map[string]string, key string) string {
	var keys []string
	var sorted []string

	for k, v := range params {
		if k != "sign" && v != "" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	for _, k := range keys {
		sorted = append(sorted, fmt.Sprintf("%s=%s", k, params[k]))
	}

	str := string.Join(sorted, "&")
	str += "&key=" + key

	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}

// 产生随机字符串
func getNonceStr(n uint8) string {
	chars := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	bytes := []byte{}
	m := len(chars)
	r := rand.New(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		bytes = append(bytes, chars[r.Intn(m)])
	}

	return string(bytes)
}
