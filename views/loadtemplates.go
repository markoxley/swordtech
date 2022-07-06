package views

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var templates *template.Template
var funcs = template.FuncMap{}

func findAndParseTemplates(root *template.Template, rootDir string) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx : len(path)-5]
			t := root.New(name)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}

// LoadTemplates loads in the html templates
func LoadTemplates() {
	templates = template.New("hive")
	var err error
	AddFunc("fullDate", func(t *time.Time) string {
		ord := "th"
		switch t.Day() {
		case 1, 21, 31:
			ord = "st"
		case 2, 22:
			ord = "nd"
		}
		f := fmt.Sprintf("15:04 on Monday, 2%s January 06", ord)
		return t.Format(f)
	})
	templates = templates.Funcs(funcs)
	// templates.Funcs(template.FuncMap{
	// 	"fullDate": func(t *time.Time) string {
	// 		ord := "th"
	// 		switch t.Day() {
	// 		case 1, 21, 31:
	// 			ord = "st"
	// 		case 2, 22:
	// 			ord = "nd"
	// 		}
	// 		f := fmt.Sprintf("15:04 on Monday, 2%s January 06", ord)
	// 		return t.Format(f)
	// 	},
	// })
	cleanRoot := filepath.Clean("views")
	pfx := len(cleanRoot) + 1

	filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx : len(path)-5]
			t := templates.New(name)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalln("Unable to load templates")
	}

}

// AddFunc adds a function to the templates
func AddFunc(n string, f interface{}) {
	funcs[n] = f
}

// View returns the view specified
func View(w http.ResponseWriter, name string, data interface{}) error {
	return templates.ExecuteTemplate(w, name, data)
}
