package handler

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Port string
}

func (e *Environment) Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("读取 `.env` 文件失败。")
	}

	e.Port = os.Getenv("PORT")
}
