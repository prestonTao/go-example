package main
 
import (
"net/http"
"io"
"io/ioutil"
"log"
"os"
)

const(
	UPLOAD_DIR = "./upload"
)


//请在本文件同级目录下创建“upload”文件夹
func main() {
    http.HandleFunc("/" , listHandler)
    http.HandleFunc("/view" , viewHandler)
    http.HandleFunc("/hello" , hello)
    err := http.ListenAndServe(":80" , nil)
    if err != nil {
        log.Fatal("ListenAndServer: " , err.Error())
    }
}

func hello(w http.ResponseWriter , r *http.Request){
    if r.Method == "GET" {
        io.WriteString(w , "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">"+
            "<html xmlns=\"http://www.w3.org/1999/xhtml\">"+
            "<head>"+
            "<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />"+
            "<title>无标题文档</title>"+
            "</head>"+
            "<body>"+
            "<form id=\"form1\"  enctype=\"multipart/form-data\" method=\"post\" action=\"/hello\">"+
              "Choose an image to upload:"+
              "<input name=\"image\" type=\"file\"  /><br/>"+
                "<input type=\"submit\" name=\"button\" id=\"button\" value=\"提交\" />"+
            "</form>"+
            "</body>"+
            "</html>")
        return
    }
    if r.Method == "POST" {
        f , h , err := r.FormFile("image")
        if err != nil {
            http.Error(w , err.Error() , http.StatusInternalServerError)
            return
        }
        fileName := h.Filename
        defer f.Close()
        
        t , err := os.Create(UPLOAD_DIR + "/" + fileName)
        if err != nil {
            http.Error(w , err.Error() , http.StatusInternalServerError)
            return
        }
        defer t.Close()

        if _ , err := io.Copy(t , f); err != nil {
            http.Error(w , err.Error() , http.StatusInternalServerError)
            return
        }

        http.Redirect(w , r , "/view?id=" + fileName , http.StatusFound)
    }

    
}

func viewHandler(w http.ResponseWriter , r *http.Request){
    imageId := r.FormValue("id")
    imagePath := UPLOAD_DIR + "/" + imageId
    if exists := isExists(imagePath); !exists {
        http.NotFound(w, r)
        return
    }

    w.Header().Set("Content-Type" , "image")
    http.ServeFile(w , r , imagePath)

}

func isExists(path string) bool {
    _ , err := os.Stat(path)
    if err == nil {
        return true
    }
    return os.IsExist(err)
}

func listHandler(w http.ResponseWriter , r *http.Request){
    fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var listHtml string
    for _, fileInfo := range fileInfoArr {
        imgid := fileInfo.Name()
        listHtml += "<li><a href=\"/view?id=" + imgid + "\">"+ imgid +"</a></li>"
    }
    io.WriteString(w, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">"+
            "<html xmlns=\"http://www.w3.org/1999/xhtml\">"+
            "<head>"+
            "<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />"+
            "<title>无标题文档</title>"+
            "</head>"+
            "<body>"+
            "<ol>" + listHtml + "</ol>"+
            "</body>"+
            "</html>" )
    
}


