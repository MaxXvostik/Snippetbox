package main

import (
	"log"
	"net/http"
)

// Создается функция-обработчик "home"
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из Snippetbox"))
}

func main() {
	
	// Используется функция http.NewServeMux() для инициализации нового рутера
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера.
	err := http.ListenAndServe(":4000", mux)

	// мы используем функцию log.Fatal() для логирования ошибок. Обратите внимание
	// что любая ошибка, возвращаемая от http.ListenAndServe(), всегда non-nil.
	log.Fatal(err)
	
	
	// вот и разобрался с gitHub
}
