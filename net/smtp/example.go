package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */

func SendMail(to, key string) error {
	user := "prestontao@126.com"
	password := "aaa--123"
	host := "smtp.126.com:25"
	subject := "Test send email by golang"
	body := "<html><body><h3>" + key + "</h3></body></html>"
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])

	content_type := "Content-Type: text/html" + "; charset=UTF-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {

	// to := "356211471@qq.com;taopopoo@126.com"
	to := "356211471@qq.com"

	err := SendMail(to, "123")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

}
