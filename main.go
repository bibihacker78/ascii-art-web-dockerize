package main

import (
	asciiartfs "ascii-art-web-stylize/banners"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var out string

var tmpl *template.Template

var indexPre []byte

const port = ":5000" // pour lancer notre serveur web

func main() {

	templateDir := "./template"

	tmpl = template.Must(template.ParseGlob(filepath.Join(templateDir, "*.html")))

	fs := http.FileServer(http.Dir("banners"))
	http.Handle("/banners/", http.StripPrefix("/banners/", fs))

	fd := http.FileServer(http.Dir("style"))

	http.Handle("/style/", http.StripPrefix("/style", fd))

	indexPre, _ = ioutil.ReadFile("./banners/style/indexpre.txt") // creation du fichier indexpre.txt qui contient le texte preremplir sur notre page

	fmt.Println("(http://localhost:5000) - server started on port", port)

	http.HandleFunc("/", asciiwebHandler) // le  pattern ici est la page "/" et la fonction asciiwebHandler

	http.ListenAndServe(port, nil) // le handler va etre géré ici par les function handler des lignes 8 & 9 donc nil
}

func asciiwebHandler(w http.ResponseWriter, r *http.Request) { // gestion de la page d'accueil

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "template/404.html")
		return
	}
	if r.Method == "GET" {

		out = string(indexPre)
		if err := tmpl.ExecuteTemplate(w, "index.html", out); err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {

		text := r.FormValue("character")
		font := r.FormValue("banner")

		str := ""

		if strings.Contains(text, "\r\n") {
			str = strings.ReplaceAll(text, "\r\n", "\\n")

		} else {
			str = text
		}
		if str == "" || font == "" {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "template/400.html")
			return

		}

		out := asciiartfs.Asciifs(str, font)

		if r.FormValue("execution") == "voir" {
			tmpl.ExecuteTemplate(w, "index.html", out)
		} else if r.FormValue("execution") == "telecharger" {
			file := strings.NewReader(out)
			w.Header().Set("Content-Disposition", "attachment; filename=fileascii.txt")
			w.Header().Set("Content-Length", strconv.Itoa(len(out)))
			io.Copy(w, file)
		}

		if out == "Error" {
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "template/500.html")
			return
		} else {
			val := tmpl.ExecuteTemplate(w, "index.html", out)
			if val != nil {
				http.Error(w, val.Error(), http.StatusInternalServerError)
			}
		}
		//fmt.Fprintf(w, "mot: %v\n", text)
		//fmt.Fprintf(w, "mot: %v\n", font)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "template/400.html")
		return
	}

}
