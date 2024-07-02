package main

import (
	"github.com/mrifkiw/belajar_gin/initializer"
	"github.com/mrifkiw/belajar_gin/model"
)

func init() {
	initializer.ConnectToDB()
	initializer.LoadEnvVariables()
}

func main() {
	initializer.DB.AutoMigrate(&model.Todo{})
}
