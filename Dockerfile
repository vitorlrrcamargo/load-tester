# Utiliza uma imagem base do Go
FROM golang:1.23.2-alpine

# Define o diretório de trabalho
WORKDIR /app

# Copia o código fonte para o contêiner
COPY . .

# Instala as dependências e compila o código
RUN go mod tidy
RUN go build -o load-tester .

# Define o comando padrão para rodar o aplicativo
ENTRYPOINT ["./load-tester"]