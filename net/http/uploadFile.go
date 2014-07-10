package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
)

// 获取大小的借口
type Sizer interface {
    Size() int64
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
    if "POST" == r.Method {
        file, _, err := r.FormFile("userfile")
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }

        defer file.Close()
        f,err:=os.Create("filenametosaveas")
        defer f.Close()
        io.Copy(f,file)

        fmt.Fprintf(w, "上传文件的大小为: %d", file.(Sizer).Size())
        return
    }

    // 上传页面
    w.Header().Add("Content-Type", "text/html")
    w.WriteHeader(200)
    html := `
<form enctype="multipart/form-data" action="/hello" method="POST">
    Send this file: <input name="userfile" type="file" />
    <input type="submit" value="Send File" />
</form>
`
    io.WriteString(w, html)
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


客户端上传文件代码：

 func Upload() (err error) {

    // Create buffer
    buf := new(bytes.Buffer) // caveat IMO dont use this for large files, \
    // create a tmpfile and assemble your multipart from there (not tested)
    w := multipart.NewWriter(buf)
    // Create file field
    fw, err := w.CreateFormFile("file", "helloworld.go") //这里的file很重要，必须和服务器端的FormFile一致
    if err != nil {
        fmt.Println("c")
        return err
    }
    fd, err := os.Open("helloworld.go")
    if err != nil {
        fmt.Println("d")
        return err
    }
    defer fd.Close()
    // Write file field from file to upload
    _, err = io.Copy(fw, fd)
    if err != nil {
        fmt.Println("e")
        return err
    }
    // Important if you do not close the multipart writer you will not have a
    // terminating boundry
    w.Close()
    req, err := http.NewRequest("POST","http://192.168.2.127/configure.go?portId=2", buf)
    if err != nil {
        fmt.Println("f")
        return err
    }
    req.Header.Set("Content-Type", w.FormDataContentType())
      var client http.Client
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("g")
        return err
    }
    io.Copy(os.Stderr, res.Body) // Replace this with Status.Code check
    fmt.Println("h")
    return err
}
