package main

import (
	"github.com/mrifkiw/belajar_gin/initializer"
	"github.com/mrifkiw/belajar_gin/model"
)

func main() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
	initializer.DB.AutoMigrate(&model.Todo{})
}
