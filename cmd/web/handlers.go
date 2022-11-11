package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// Инициализируем срез содержащий пути к двум файлам.
	file := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Используем функцию template.ParseFiles() для чтения файла шаблона.
	ts, err := template.ParseFiles(file...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, nil)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Отображение определенной заметки с ID %d...", id)
}
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Создание новой заметки..."))
}
