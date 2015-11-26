package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)       //hello word
	http.HandleFunc("/json", ResponseJson) //返回json格式例子
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
	fmt.Println("webServer startup...")
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello World")
}

/*
	参数在请求body中传递
	返回参数用json格式
	@return   RetCode   0=成功，1=失败
*/
func ResponseJson(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})
	params := make(map[string]interface{})
	robots, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ret["RetCode"] = 1
		ret["Msg"] = "参数格式化错误"
		chs, _ := json.Marshal(ret)
		fmt.Fprintln(w, string(chs))
		return
	}
	json.Unmarshal(robots, &params)
	coinNum := int(params["CoinNum"].(float64))
	phone := params["Phone"].(string)
	ret, err = actor.PlayerManager.Call(config.P_M_PLAYER_ADDCOIN, map[string]interface{}{"Num": coinNum, "Phone": phone})
	if err != nil {
		ret["RetCode"] = 1
	}
	chs, _ := json.Marshal(ret)
	fmt.Fprintln(w, string(chs))
}
