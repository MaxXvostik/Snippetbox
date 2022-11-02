package main

import (
	"log"
	"net/http"
)

// Создается функция-обработчик "home"
func home(w http.ResponseWriter, r *http.Request) {

	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	// Важно, чтобы мы завершили работу обработчика через return.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Привет из Snippetbox"))
}

// Обработчик для отображения содержимого заметки.
func showSnippert(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отоброжение заметки..."))
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для создания новой заметки..."))
}

func main() {

	// Используется функция http.NewServeMux() для инициализации нового рутера
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippert", showSnippert)
	mux.HandleFunc("/snippert/create", createSnippet)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера.
	err := http.ListenAndServe(":4000", mux)

	// мы используем функцию log.Fatal() для логирования ошибок. Обратите внимание
	// что любая ошибка, возвращаемая от http.ListenAndServe(), всегда non-nil.
	log.Fatal(err)
}
