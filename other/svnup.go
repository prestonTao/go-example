//svn up with golang online
//go 调用php执行svn up.为什么不直接php调用,因为php命令行执行的是root用户,但apache执行php是apache账号
//apache执行php调用svn up存在权限问题.通过go调用php,php调用svn up避开权限问题.
//为什么不go直接调用svn up?因为php文件便于修改和扩展新功能
package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"os/exec"
	//"strings"
)

func main() {
	p := flag.String("p", "5200", "port")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                w.Header().Add("content-type","text/html")
		if _, err := os.Stat("." + r.URL.Path); err == nil {}
			println(r.URL.Path)
			cmd := exec.Command("php","-f","svnup.php",r.URL.Path)

			d, e := cmd.Output()
			if e != nil {
				println(e)
			}

			//println(string(d))
			io.WriteString(w, string(d))
		

	})
	http.ListenAndServe("0.0.0.0:"+*p, nil)

}
