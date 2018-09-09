package wxpay

import (
	"bytes"
	"crypto/md5"
	"encoding/xml"
	"errors"
	"fmt"
	"math/rand"
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

	signStr, err := makeXMLSign(data, key)
	if err != nil {
		return
	}

	param.SetSign(signStr)

	data, err = xml.Marshal(param)
	return
}

// checkXMLSign 校验xml数据有效性
func checkXMLSign(data []byte, key string) error {
	xmlMap, err := mxj.NewMapXml(data)
	if err != nil {
		return err
	}

	dataMap, ok := xmlMap["xml"].(map[string]interface{})
	if !ok {
		return errors.New("非法xml")
	}

	signStr, ok := dataMap["sign"].(string)
	if !ok {
		return errors.New("无sign值")
	}

	sign := makeSign(dataMap, key)
	if signStr != sign {
		return errors.New("sign错误")
	}

	return nil
}

// makeXMLSign 生成xml数据的sign值
func makeXMLSign(data []byte, key string) (signStr string, err error) {
	xmlMap, err := mxj.NewMapXml(data)
	if err != nil {
		return "", err
	}

	dataMap, ok := xmlMap["xml"].(map[string]interface{})
	if !ok {
		return "", errors.New("非法xml")
	}

	return makeSign(dataMap, key), nil
}

// 生成sign值
func makeSign(m map[string]interface{}, key string) string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	buf := new(bytes.Buffer)

	for _, k := range keys {
		if k == "sign" {
			continue
		}

		v := fmt.Sprint(m[k])
		if v == "" {
			continue
		}

		buf.WriteString(k + "=" + v + "&")
	}

	buf.WriteString("key=" + key)
	return fmt.Sprintf("%X", md5.Sum(buf.Bytes()))
}
