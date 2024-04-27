package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/liujunren93/bian/utils"
)

func Sign(queryPrams, bodyPrams any, key string) (queryStr string, requestBody map[string]any, err error) {
	var bodyStr string
	now := time.Now().UnixMilli()
	if queryPrams != nil {
		var queryMap map[string]any
		queryMap, err = utils.Struct2MapNoZero(queryPrams)
		if err != nil {
			return "", nil, err
		}
		queryMap["timestamp"] = now
		queryStr, err = Map2QueryString(queryMap)
		if err != nil {
			return "", nil, err
		}
	}
	if bodyPrams != nil {
		requestBody, err = utils.Struct2MapNoZero(bodyPrams)
		if err != nil {
			return "", nil, err
		}
		requestBody["timestamp"] = now
		bodyStr, err = Map2QueryString(requestBody)
		if err != nil {
			return "", nil, err
		}

	}
	var signStr = strings.Builder{}
	if len(queryStr) > 0 {
		signStr.WriteString(queryStr)
	}
	if len(bodyStr) > 0 {
		signStr.WriteString(bodyStr)
	}
	fmt.Println(signStr.String())
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(signStr.String()))
	hmacValue := h.Sum(nil)
	// signBase64Str := base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(hmacValue)))
	signBase64Str := hex.EncodeToString(hmacValue)
	if bodyPrams != nil {
		requestBody["signature"] = signBase64Str
	} else if queryPrams != nil {
		queryStr += "&signature=" + signBase64Str
	}
	return
}

func Map2QueryString(dataMap map[string]interface{}) (string, error) {
	val := url.Values{}
	for k, v := range dataMap {
		val.Add(k, fmt.Sprintf("%v", v))
	}
	return val.Encode(), nil
	// keys := make([]string, 0, len(dataMap))
	// for k := range dataMap {
	// 	keys = append(keys, k)
	// }
	// sort.Strings(keys)
	// var queryParams []string
	// for _, k := range keys {
	// 	queryParams = append(queryParams, fmt.Sprintf("%s=%v", k, dataMap[k]))
	// }
	// return strings.Join(queryParams, "&"), nil

}
