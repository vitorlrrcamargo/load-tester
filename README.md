# Load Tester

O **Load Tester** é uma ferramenta de linha de comando (CLI) desenvolvida em **Go** para realizar testes de carga em um serviço web. Ele permite enviar múltiplas requisições HTTP para testar a performance de um serviço sob carga.

### Instalação

Clone o repositório e compile o projeto:
```bash
git clone https://github.com/seu-usuario/load-tester.git cd load-tester go build -o load-tester .
```

### Executando o Teste

Para compilar o projeto, execute:
```bash
go build .
```
Isso gerará o binário load-tester que você pode usar para rodar os testes de carga.

Para rodar o teste de carga, execute o comando abaixo, substituindo a URL pelo serviço que deseja testar, e defina os parâmetros de requisições e concorrência:
```bash
./load-tester --url=http://google.com --requests=100 --concurrency=10
```

### Executando com Docker

Para rodar com Docker, primeiro crie a imagem:
```bash
docker build -t load-tester .
```

Execute o contêiner com a URL do serviço que deseja testar, o número de requisições e a quantidade de concorrência:
```bash
docker run load-tester --url=http://google.com --requests=1000 --concurrency=10
```

### Parâmetros

- `--url`: A URL do serviço que será testado.
- `--requests`: Número total de requisições.
- `--concurrency`: Número de requisições simultâneas.

### Relatório

Após a execução, será gerado um relatório com o tempo total de execução, a quantidade de requisições realizadas, o número de requisições com status 200, e a distribuição de outros códigos de status HTTP.

### Autor
Feito com 💛 por vitorlrrcamargo
