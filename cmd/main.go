package main

import (
	"auth_sevice_microservice/internal/server"
	"gitlab.com/nevasik7/lg"
)

func main() {
	lg.Init()
	//Запуск Сервера
	if err := server.Run(); err != nil {
		lg.Fatalf("Не удалось запусить сервер")
	}

}
