package main
 
import (
"net/http"
"io"
"os"
"io/ioutil"
"log"
"runtime/debug"
"html/template"
"path"
)

const(
    ListDir = "0x0001"
	UPLOAD_DIR = "./upload"
    TEMPLATE_DIR = "./views"
)

var templates map[string] *template.Template

func init(){
    templates = make(map[string] *template.Template)
    fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
    check(err)
    var templateName, templatePath string
    for _, fileInfo := range fileInfoArr {
        templateName = fileInfo.Name()
        if ext := path.Ext(templateName); ext != ".html" {
            continue
        }
        templatePath = TEMPLATE_DIR + "/" + templateName
        log.Println("Loading template:", templatePath)
        t := template.Must(template.ParseFiles(templatePath))
        templates[templatePath] = t
    }
}


//请在本文件同级目录下创建“upload”文件夹
func main() {
    mux := http.NewServeMux()
    staticDirHandler(mux, "/assets/", "./public", 0)
    http.HandleFunc("/" , safeHandler(listHandler))
    http.HandleFunc("/view" , safeHandler(viewHandler))
    http.HandleFunc("/hello" , safeHandler(hello))
    err := http.ListenAndServe(":80" , mux)
    if err != nil {
        log.Fatal("ListenAndServer: " , err.Error())
    }
}

func hello(w http.ResponseWriter , r *http.Request){
    if r.Method == "GET" {
        renderHtml(w, "upload", nil)
    }
    if r.Method == "POST" {
        f , h , err := r.FormFile("image")
        check(err)
        fileName := h.Filename
        defer f.Close()
        t , err := ioutil.TempFile(UPLOAD_DIR, filename)
        check(err)
        defer t.Close()
        _ , err = io.Copy(t , f)
        check(err)

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
    check(err)
    locals := make(map[string]interface{})
    images := []string{}
    for _, fileInfo := range fileInfoArr {
        images = append(images, fileInfo.Name())
    }
    locals["images"] = images
    renderHtml(w, "list", locals)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
    err =  templates[tmpl].Execute(w, locals)
    check(err)
}

func check(err error){
    if err != nil {
        panic(err)
    }
}

func safeHandler(fn http.HandlerFunc, ) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request){
        defer func(){
            if e, ok := recover().(error); ok {
                http.Error(w, e.Error(), http.StatusInternalServerError)
                //或者输出自定义的50X错误页面
                //renderHtml(w, "error", e)
                log.Println("WARN: panic in %v - %v", fn, e)
                log.Println(string(debug.Stack()))
            }
        }()
        fn(w, r)
    }
}

func staticDirHandler(mux *http.ServerMux, prefix string, staticDir string, flags int){
    mux.HandleFunc(prefix, func(w, http.ResponseWriter, r *http.Request){
        file := staticDir + r.URL.Path[len(prefix)-1:]
        if (flags & ListDir) == 0 {
            if exists := isExists(file); !exists {
                http.NotFound(w, r)
                return
            }
        }
        http.ServeFile(w, r, file)
    })
}
