# Sistema para capturar beneficios de usuários usando o CPF

## Requisitos

1) Go (https://go.dev/doc/install)
1) Docker (https://docs.docker.com/engine/install)
3) Docker compose (https://docs.docker.com/compose/)

## Como iniciar o projeto via docker

1) Vá até onde o arquivo docker-compose.yml está (Raiz do projeto)
2) Rode o comando: docker compose up

## Como iniciar o projeto manualmente (sem docker)
1) Vá até a pasta cmd/api e crie o arquivo .env com base no .env.example
2) Vá até a pasta cmd/msbenefit e crie o arquivo .env com base no .env.example
3) Na raiz do projeto rode o comando: go mod tidy
4) Vá até a pasta "api" e rode o comando go run main.go
5) Vá até a pasta "msbenefit" e rode o comando go run main.go

## Como rodar os testes

1) Inicie o projeto manualmente (sem docker)
2) Rode o comando na pasta "api": go test ./tests -v

## Endpoint e Swagger

1) Endpoint: http://localhost:8080/api/benefit?doc=927.100.938-04,351.165.848-95
2) OBS: Pode incluir quantos cpfs quiser colocando "vírgula"

3) Swagger: http://localhost:8080/docs/index.html

