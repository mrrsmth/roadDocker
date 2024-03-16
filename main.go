package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

type Roadmap struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Technology string `json:"technology"`
	Theme      string `json:"theme"`
	Bool       bool   `json:"bool"`
}

func main() {
	db, err := gorm.Open("mysql", "root:secret@tcp(localhost:3306)/road?parseTime=true") // Исправлено на правильный порт и добавлено /road в строке подключения
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&Roadmap{})

	roadmap := []Roadmap{
		{Technology: "HTML", Theme: "Элементы в HTML", Bool: false},
		{Technology: "", Theme: "Формы, валидация форм", Bool: false},
		{Technology: "", Theme: "Семантическая верстка", Bool: false},
		{Technology: "CSS", Theme: "Селекторы", Bool: false},
		{Technology: "", Theme: "Свойства", Bool: false},
		{Technology: "", Theme: "Позиционирование элементов, Flexbox, Grid", Bool: false},
		{Technology: "", Theme: "Трансформации, переходы, анимации", Bool: false},
		{Technology: "", Theme: "Адаптивный дизайн и медиазапросы", Bool: false},
		{Technology: "", Theme: "CSS-препроцессоры(sass, scss, less)", Bool: false},
		{Technology: "", Theme: "БЭМ", Bool: false},
		{Technology: "JavaScript", Theme: "Типы данных, преобразования типов", Bool: false},
		{Technology: "", Theme: "Условное ветвление, логические операторы, циклы", Bool: false},
		{Technology: "", Theme: "Функции, функциональные выражения, стрелочные функции, поднятие", Bool: false},
		{Technology: "", Theme: "Замыкание, IIFE", Bool: false},
		{Technology: "", Theme: "Строки, шаблонные строки, регулярные выражения", Bool: false},
		{Technology: "", Theme: "Массивы, методы массивов, перебор массивов", Bool: false},
		{Technology: "", Theme: "Объекты, методы объектов, сравнение объектов, ссылки", Bool: false},
		{Technology: "", Theme: "Классы, наследование, статические свойства и методы, защищенные свойства и методы", Bool: false},
		{Technology: "", Theme: "Колбэки, промисы, обработка ошибок, микротаски, азупс/амаи, event loop", Bool: false},
		{Technology: "", Theme: "Взаимодействие DOM (создание, добавление, изменение и удаление элементов веб-станиць), браузерные события, распространение событий", Bool: false},
		{Technology: "", Theme: "Хранение данных - сооке, session storage, local storage", Bool: false},
		{Technology: "", Theme: "Дебаг в Chrome DevTools", Bool: false},
		{Technology: "", Theme: "XMLHttpRequest, FetchApi, WebSocket", Bool: false},
		{Technology: "Angular", Theme: "Компоненты, модули, загрузка приложения", Bool: false},
		{Technology: "", Theme: "Привязка данных", Bool: false},
		{Technology: "", Theme: "Привязка к событиям дочернего компонента", Bool: false},
		{Technology: "", Theme: "Жизненный цикл компонента", Bool: false},
		{Technology: "", Theme: "Атрибутивные и структурные директивы, создание директив", Bool: false},
		{Technology: "", Theme: "Сервисы и dependency injection", Bool: false},
		{Technology: "", Theme: "Работа с формами", Bool: false},
		{Technology: "", Theme: "НТТР и взаимодействие с сервером", Bool: false},
		{Technology: "", Theme: "Маршрутизация", Bool: false},
		{Technology: "", Theme: "Pipes", Bool: false},
	}

	for _, data := range roadmap {
		db.Create(&data)
	}

	// Создаем новый роутер Chi
	router := chi.NewRouter()

	// Настройка CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},                   // Укажите разрешенные источники (оригинаы запроса) для CORS
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Укажите разрешенные методы HTTP
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsOptions.Handler)

	// Обрабатываем GET-запрос для эндпоинта "/roadmaps"
	router.Get("/roadmaps", func(w http.ResponseWriter, r *http.Request) {
		// Получаем все элементы roadmaps из базы данных
		var roadmaps []Roadmap
		if err := db.Find(&roadmaps).Error; err != nil {
			log.Fatalf("Не удалось получить данные из базы данных: %v", err)
		}

		// Преобразуем данные roadmaps в формат JSON
		jsonData, err := json.Marshal(roadmaps)
		if err != nil {
			log.Fatalf("Не удалось преобразовать данные в формат JSON: %v", err)
		}

		// Устанавливаем заголовки ответа и записываем JSON-данные в ответ
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(jsonData))
	})

	// Запускаем HTTP-сервер и слушаем порт 8082
	log.Fatal(http.ListenAndServe(":8082", router))

	var roadmaps []Roadmap
	db.Find(&roadmaps)

	for _, roadmap := range roadmaps {
		fmt.Printf("ID: %d, Technology: %s, Theme: %s, Bool: %t\n", roadmap.ID, roadmap.Technology, roadmap.Theme, roadmap.Bool)
	}

	fmt.Println("Roadmap успешно создан")
}
