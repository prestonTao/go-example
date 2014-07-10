package main
 
import (
"net/http"
"io"
"io/ioutil"
"log"
"os"
"html/template"
"path"
)

const(
	UPLOAD_DIR = "./upload"
    TEMPLATE_DIR = "./views"
)

var templates map[string] *template.Template

func init(){
    templates = make(map[string] *template.Template)
    fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
    if err != nil {
        panic(err)
        return
    }
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
        if err := renderHtml(w, "upload", nil); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
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
    locals := make(map[string]interface{})
    images := []string{}
    for _, fileInfo := range fileInfoArr {
        images = append(images, fileInfo.Name())
    }
    locals["images"] = images
    if err = renderHtml(w, "list", locals); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
    err =  templates[tmpl].Execute(w, locals)
    return
}

