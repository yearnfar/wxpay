package wxpay

import (
	"bytes"
	"crypto/md5"
	"encoding/xml"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strings"

	"github.com/clbanning/mxj"
)

var (
	RandChar          = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	RandCharLen int32 = 36
)

// makeNonceStr 生成随机字符串
func makeNonceStr(n int) string {
	sb := new(strings.Builder)
	for i := 0; i < n; i++ {
		sb.WriteByte(RandChar[rand.Int31n(RandCharLen)])
	}
	return sb.String()
}

// 转化为xml数据
func toXML(param Parameter, key string) (data []byte, err error) {
	err = param.Valid()
	if err != nil {
		return
	}

	data, err = xml.Marshal(param)
	if err != nil {
		return
	}

	signStr, err := makeSign(data, key)
	if err != nil {
		return
	}

	param.SetSign(signStr)

	data, err = xml.Marshal(param)
	return
}

// Send 发送请求
func Send(client *http.Client, method string, urlStr string, param Parameter, resp Responser, key string) error {
	req, err := NewRequest(method, urlStr, param, key)
	if err != nil {
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	err = ParseResponse(response, resp, key)
	if err != nil {
		return err
	}

	return nil
}

// 生成sign
func makeSign(data []byte, key string) (signStr string, err error) {
	xmlMap, err := mxj.NewMapXml(data)
	if err != nil {
		return
	}

	dataMap, ok := xmlMap["xml"].(map[string]interface{})
	if !ok {
		err = errors.New("非法xml")
		return
	}

	var keys []string
	for k := range dataMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	buf := new(bytes.Buffer)

	for _, k := range keys {
		if k == "sign" {
			continue
		}

		v := fmt.Sprint(dataMap[k])
		if v == "" {
			continue
		}

		buf.WriteString(k + "=" + v + "&")
	}

	buf.WriteString("key=" + key)

	fmt.Printf("sign-text: %s\n", buf.String())

	signStr = fmt.Sprintf("%X", md5.Sum(buf.Bytes()))
	return
}
