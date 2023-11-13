package meituan

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type HttpMethod string

const PostMethod = "POST"
const GetMethod = "GET"

type Request struct {
	HttpMethod HttpMethod
	Url        string
	Query      *url.Values
	Body       *url.Values
}

func (m MT) processRquest(url string, method HttpMethod, query interface{}, postBody interface{}, result Response) (errorResp *ErrorResponse, err error) {
	var (
		resp        []byte
		queryValues = parseDataToUrlValues(query)
		postValues  = parseDataToUrlValues(postBody)
	)

	resp, err = m.apiRequest(Request{
		HttpMethod: method,
		Url:        url,
		Query:      &queryValues,
		Body:       &postValues,
	})

	if err != nil {
		return
	}

	// 尝试解析错误返回值
	if err = json.Unmarshal(resp, &errorResp); err == nil && errorResp.GetResultCode() > 1 {
		return
	} else {
		errorResp = nil
	}

	if err = json.Unmarshal(resp, result); err != nil {
		return
	}

	return
}

func (m MT) apiRequest(req Request) ([]byte, error) {
	var (
		response *http.Response
		client   = http.Client{}
		result   []byte
		err      error
	)

	reqUrl, sigSource, sig := generateUrl(ApiUrl+req.Url, req.HttpMethod, m.AppID, m.AppSecret, *req.Query, *req.Body, 0)

	switch req.HttpMethod {
	case http.MethodPost:
		fmt.Printf("\n========http request debug========\n $Method: %v\n $Url: %v\n $body: %v\n $reqBody: %s\n $sigSource: %v\n $sig: %v\n", req.HttpMethod, reqUrl, req.Body, req.Body, sigSource, sig)
		response, err = client.Post(reqUrl, "application/x-www-form-urlencoded", strings.NewReader(req.Body.Encode()))
	case http.MethodGet:
		fmt.Printf("\n========http request debug========\n $Method: %v\n $Url: %v\n $sigSource: %v\n $sig: %v\n", req.HttpMethod, reqUrl, sigSource, sig)
		response, err = client.Get(reqUrl)
	}

	if err != nil {
		return nil, err
	}

	result, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error, error: %v", err)
	}

	fmt.Printf(" $response: %v\n", string(result))
	fmt.Printf("========finish========")

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetResponse code: %v body: %s", response.StatusCode, string(result))
	}

	return result, nil
}

// 请求query编码，获取sig，输出最终完整的请求url
func generateUrl(reqUrl string, method HttpMethod, appID, appSecret string, values, postValue url.Values, timestamp int64) (finallUrl, sigSource, sig string) {
	if timestamp == 0 {
		timestamp = time.Now().Local().Unix()
	}

	// 添加当前时间戳，服务器时间必须是准确的北京时间
	values.Add("timestamp", fmt.Sprintf("%d%s", timestamp, appSecret))
	values.Add("app_id", appID)

	for key, vs := range postValue {
		for _, value := range vs {
			values.Add(key, fmt.Sprintf("%v", value))
		}
	}

	// 获取所有的键
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}

	// 对键进行排序
	sort.Strings(keys)

	// 构建排序后的查询字符串
	var queryString string
	for _, key := range keys {
		valuesForKey := values[key]
		for _, value := range valuesForKey {
			if queryString != "" {
				queryString += "&"
			}
			if key == "timestamp" {
				value = strconv.FormatInt(timestamp, 10)
			}
			queryString += fmt.Sprintf("%s=%s", key, value)
		}
	}

	// queryString = fmt.Sprintf("%s?%s", reqUrl, queryString)

	// 生成sig
	sigSource = fmt.Sprintf("%s?%s%s", reqUrl, queryString, appSecret)
	sig = md5String(sigSource)

	if method == PostMethod {
		finallUrl = fmt.Sprintf("%s?app_id=%s&timestamp=%d&sig=%s", reqUrl, appID, timestamp, sig)
	} else {
		finallUrl = fmt.Sprintf("%s?%s&sig=%s", reqUrl, queryString, sig)
	}

	return
}

func md5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func parseDataToUrlValues(data interface{}) (values url.Values) {
	var tmp map[string]interface{}
	t, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(t, &tmp)
	if err != nil {
		return
	}

	values = url.Values{}
	for k, v := range tmp {
		if v == nil {
			continue
		}
		t := reflect.TypeOf(v)
		switch t.Kind() {
		case reflect.Slice:
			fallthrough
		case reflect.Map:
			fallthrough
		case reflect.Struct:
			if bytes, err := json.Marshal(v); err == nil {
				values.Add(k, string(bytes))
			}
		default:
			values.Add(k, fmt.Sprintf("%v", v))
		}
	}

	return
}

// 美团部分接口返回的json字符串中有嵌套结构体，但是有多余的"和\导致Go不能正常解析，必须先替换掉
func escapeJsonCharactor(data []byte) []byte {
	str := string(data)
	str = strings.ReplaceAll(str, "\\\"", "\"")
	str = strings.ReplaceAll(str, "\":\"{", "\":{")
	str = strings.ReplaceAll(str, "}\",", "},")
	str = strings.ReplaceAll(str, "\":\"[", "\":[")
	str = strings.ReplaceAll(str, "]\",", "],")
	return []byte(str)
}

func CallbackError() CallbackResponse {
	return CallbackResponse{Data: "error"}
}

func CallbackSuccess() CallbackResponse {
	return CallbackResponse{Data: "success"}
}

func CheckSig(url, sig string, m MT, body interface{}, timestamp int64) bool {
	// 生成sig
	_, _, genSig := generateUrl(url, PostMethod, m.AppID, m.AppSecret, parseDataToUrlValues(nil), parseDataToUrlValues(body), timestamp)
	return genSig == sig
}
