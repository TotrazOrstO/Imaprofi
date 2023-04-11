package rest

import (
	"database/"
	"net/http"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Получение данных пользователя из запроса
	// ...

	// Использование компонента PostgreSQL для сохранения пользователя в базе данных
	err := database.SaveUser(postgresDB, user)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	// Использование компонента RabbitMQ для отправки уведомления о создании пользователя
	err = messaging.PublishUserCreated(rabbitMQConn, user)
	if err != nil {
		http.Error(w, "Error sending user created notification", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
