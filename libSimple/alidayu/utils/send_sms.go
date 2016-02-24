package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	URL               string = "http://gw.api.taobao.com/router/rest"   //正式环境
	Method_SendSMS    string = "alibaba.aliqin.fc.sms.num.send"         //发送短信
	Method_CallTTS    string = "alibaba.aliqin.fc.tts.num.singlecall"   //
	Method_CallVoice  string = "alibaba.aliqin.fc.voice.num.singlecall" //
	Method_CallDouble string = "alibaba.aliqin.fc.voice.num.doublecall" //
)

var AppKey string
var AppSecret string

/*
   发送文字短信
   @rec_num               string    目标手机号
   @sms_free_sign_name    string    短信名称
   @sms_template_code     string    短信模板名称
   @sms_param             string
*/
func SendSMS(rec_num, sms_free_sign_name, sms_template_code, sms_param string) (success bool, response string) {
	if rec_num == "" || sms_free_sign_name == "" || sms_template_code == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_SendSMS
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["sms_type"] = "normal"
	params["sms_free_sign_name"] = sms_free_sign_name
	params["rec_num"] = rec_num
	params["sms_template_code"] = sms_template_code
	params["sms_param"] = sms_param

	return DoPost(params)

}

/*
	检查返回编号
	@return    int    返回错误编号，0=成功,-1=解析参数失败
*/
func CheckSMSResult(result string) int {
	fmt.Println(result)
	rep := make(map[string]interface{})
	err := json.Unmarshal([]byte(result), &rep)
	if err != nil {
		return -1
	}
	if v, ok := rep["error_response"]; ok {
		fmt.Println(v)
		codeMap := v.(map[string]interface{})
		fmt.Println(codeMap)
		code := int(codeMap["code"].(float64))
		return code
	}
	if v, ok := rep["alibaba_aliqin_fc_sms_num_send_response"]; ok {
		codeMap := v.(map[string]interface{})
		codeStr := codeMap["result"].(map[string]interface{})["err_code"].(string)
		code, err := strconv.Atoi(codeStr)
		if err != nil {
			return -1
		}
		return code
	}
	return -1
}

func getRequestBody(m map[string]string) (reader io.Reader, size int64) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	v := url.Values{}

	signString := AppSecret
	for _, k := range keys {
		v.Set(k, m[k])
		signString += k + m[k]
	}
	signString += AppSecret

	signByte := md5.Sum([]byte(signString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	v.Set("sign", sign)

	return ioutil.NopCloser(strings.NewReader(v.Encode())), int64(len(v.Encode()))
}

func CallDouble(caller_num, caller_show_num, called_num, called_show_num string) (success bool, response string) {
	if caller_num == "" || caller_show_num == "" || called_num == "" || called_show_num == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_CallVoice
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["caller_num"] = caller_num
	params["caller_show_num"] = caller_show_num
	params["called_num"] = called_num
	params["called_show_num"] = called_show_num

	return DoPost(params)
}

func CallTTS(called_num, called_show_num, tts_code, tts_param string) (success bool, response string) {
	if called_num == "" || called_show_num == "" || tts_code == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_CallTTS
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["called_show_num"] = called_show_num
	params["called_num"] = called_num
	params["tts_code"] = tts_code
	params["tts_param"] = tts_param

	return DoPost(params)
}

func CallVoice(called_num, called_show_num, voice_code string) (success bool, response string) {
	if called_num == "" || called_show_num == "" || voice_code == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_CallVoice
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["called_show_num"] = called_show_num
	params["called_num"] = called_num
	params["voice_code"] = voice_code

	return DoPost(params)
}

func DoPost(m map[string]string) (success bool, response string) {
	if AppKey == "" || AppSecret == "" {
		return false, "AppKey or AppSecret is requierd!"
	}

	body, size := getRequestBody(m)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", URL, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = size

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return false, err.Error()
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return true, string(data)
}
