package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
	"unsafe"
)

type Signer interface {
	Add(string, interface{})
	Len() int
	String() string
}

func Sign(t1 Signer, secret string) Signer {
	timestamp := time.Now().Local().UnixMilli()
	if t1 != nil {
		t1.Add("timestamp", timestamp)

		t1.Add("signature", ComputeHmacSha256(t1.String(), secret))
	}
	return t1
}

func Sign1(t1, t2 Signer, secret string) Signer {
	timestamp := time.Now().Local().UnixMilli()
	// timestamp = 1499827319559
	if t1 != nil && t2 != nil && t1.Len() > 0 && t2.Len() > 0 {

		t1Str := t1.String()
		t2.Add("timestamp", timestamp)
		t2Str := t2.String()
		t2.Add("signature", ComputeHmacSha256(t1Str+t2Str, secret))
		return t2
	}

	t1.Add("timestamp", timestamp)
	t1.Add("signature", ComputeHmacSha256(t1.String(), secret))
	return t1
}

type QueryParams map[string]interface{}

func (q QueryParams) ParseQuery(query string) (err error) {
	if strings.Index(query, "?") == -1 {
		return nil
	}
	query = query[strings.Index(query, "?")+1:]
	for query != "" {
		var key string
		key, query, _ = strings.Cut(query, "&")
		if strings.Contains(key, ";") {
			err = fmt.Errorf("invalid semicolon separator in query")
			continue
		}
		if key == "" {
			continue
		}
		key, value, _ := strings.Cut(key, "=")
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		q[key] = value
	}
	return
}
func (q QueryParams) Len() int {
	return len(q)
}
func (q QueryParams) Add(key string, val interface{}) {
	q[key] = val
}

func (q QueryParams) String() string {
	tmp := *(*Params)(unsafe.Pointer(&q))
	return tmp.String()
}

type Params map[string]interface{}

func (m Params) String() string {
	if m == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		vs := m[k]
		keyEscaped := url.QueryEscape(k)
		buf.WriteString(keyEscaped)
		buf.WriteByte('=')

		buf.WriteString(url.QueryEscape(fmt.Sprintf("%v", vs)))

	}
	return buf.String()

	// var buf = strings.Builder{}
	// for k, v := range m {
	// 	if buf.Len() == 0 {
	// 		buf.WriteString(fmt.Sprintf("%s=%v", k, v))
	// 	} else {
	// 		buf.WriteString(fmt.Sprintf("&%s=%v", k, v))
	// 	}
	// }

	return buf.String()
}
func (m Params) Len() int {
	return len(m)
}
func (m Params) Add(key string, val interface{}) {
	m[key] = val
}
func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return hex.EncodeToString(h.Sum(nil))
}
