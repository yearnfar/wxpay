package wxpay

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
)

// 发送Xml请求
func sendXmlRequest(method, path string, xmlString string, tlsConfig *tls.Config, timeout time.Duration) (body []byte, err error) {
	req, err := http.NewRequest(method, path, bytes.NewBufferString(xmlString))
	if err != nil {
		return
	}

	client := http.Client{}

	if timeout > 0 {
		client.Timeout = timeout * time.Second
	}

	if tlsConfig != nil {
		client.Transport = &http.Transport{TLSClientConfig: tlsConfig}
	}

	resp, err := client.Do(req)
	if err != nil {
		err = errors.New("request fail")
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// map转xml
func map2Xml(params map[string]string) string {
	xmlString := "<xml>"

	for k, v := range params {
		xmlString += fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xmlString += "</xml>"
	return xmlString
}

// xml转map
func xml2Map(in interface{}) (map[string]string, error) {
	xmlMap := make(map[string]string)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("xml2Map only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		tagv := fi.Tag.Get("xml")

		if strings.Contains(tagv, ",") {
			tagvs := strings.Split(tagv, ",")

			switch tagvs[1] {
			case "innerXml":
				innerXmlMap, err := xml2Map(v.Field(i).Interface())
				if err != nil {
					return nil, err
				}
				for k, v := range innerXmlMap {
					if _, ok := xmlMap[k]; !ok {
						xmlMap[k] = v
					}
				}
			}
		} else if tagv != "" && tagv != "xml" {
			xmlMap[tagv] = v.Field(i).String()
		}
	}
	return xmlMap, nil
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

	str := strings.Join(sorted, "&")
	str += "&key=" + key

	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}

// 产生随机字符串
func getNonceStr(n int) string {
	chars := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	value := []byte{}
	m := len(chars)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		value = append(value, chars[r.Intn(m)])
	}

	return string(value)
}
