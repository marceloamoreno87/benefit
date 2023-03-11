# Sistema para capturar beneficios de usuários usando o CPF

## Baixando e Instalando pacotes

1) Na raiz do projeto rode o comando: go mod tidy

## Configurando variáveis de ambiente

1) Vá até a pasta cmd/api e crie o arquivo .env com base no .env.example
2) Vá até a pasta cmd/msbenefit e crie o arquivo .env com base no .env.example

## Requisitos

1) Docker (https://docs.docker.com/engine/install)
2) Docker compose (https://docs.docker.com/compose/)

## Como iniciar o projeto via docker

1) Vá até onde o arquivo docker-compose.yml está (Raiz do projeto)
2) Rode o comando: docker compose up

## Requisitos

1) Go (https://go.dev/doc/install)

## Como iniciar o projeto manualmente (sem docker)

1) Vá até a pasta "api" e rode o comando go run main.go
2) Vá até a pasta "msbenefit" e rode o comando go run main.go

## Link do swagger

URL: http://localhost:8080/docs/index.html

## Como rodar os testes

1) Inicie o projeto manualmente (sem docker)
2) Rode o comando na pasta "api": go test ./tests -v
