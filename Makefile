.PHONY:
.SILENT:
.DEFAULT_GOAL:= run

run: ## Запустить приложение
	go run ./cmd/main.go

build: ## Сборка приложения с тэгом
	go build ./cmd/main.go
