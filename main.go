package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
)

var port = flag.String("port", "80", "port number")
var tpl *template.Template

type indexPage struct {
	Menu    []os.FileInfo
	Index   []os.FileInfo
	Name    string
	Article bool
}

type articlePage struct {
	Menu    []os.FileInfo
	Article template.HTML
}

type byDate []os.FileInfo
func (f byDate) Len() int { return len(f) }
func (f byDate) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f byDate) Less(i, j int) bool {
	return f[i].ModTime().After(f[j].ModTime())
}

func getSortedDir(name string) []os.FileInfo {
	dir, err := ioutil.ReadDir(name)
	if err != nil {
		log.Print(err)
		return []os.FileInfo{}
	}

	sort.Sort(byDate(dir))
	return dir
}

func showIndex(w http.ResponseWriter, r *http.Request) {
	index := indexPage{
		Menu:  getSortedDir("."),
		Index: getSortedDir(r.URL.Path[len("/show/"):]),
		Name:  path.Base(r.URL.Path),
	}

	tpl.Execute(w, index)
}

func showArticle(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile(r.URL.Path[len("/read/"):])
	if err != nil {
		fmt.Fprint(w, "no such article")
		return
	}

	article := articlePage{
		Menu:    getSortedDir("."),
		Article: template.HTML(file),
	}

	tpl.Execute(w, article)
}

func redirToMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/show/articles", http.StatusFound)
}

func cutext(title string) string {
	if i := strings.LastIndex(title, "."); i > 0 {
		title = title[:i]
	}
	return title
}

func main() {
	tpl = template.Must(template.New("index.html").
		Funcs(template.FuncMap{"title": strings.Title, "cutext": cutext}).
		ParseFiles("index.html"))
	flag.Parse()
	http.HandleFunc("/", redirToMain)
	http.HandleFunc("/show/", showIndex)
	http.HandleFunc("/read/", showArticle)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
