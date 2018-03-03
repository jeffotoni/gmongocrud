# gmongocrud

* Responda para o endereço avaliacao.com. ( local )
* Esse site deve responder em HTTPS e se chamado em HTTP, ser redirecionado para HTTPS. ( nativo e com nginx)
* A funcionalidade do sistema deve ser uma API para cadastrar e consultar os dados do seu currículo. (insert e listar)
* As ações com a API deve ser feita de forma autenticada. (JWT, usando login e senha ou uma chave)
* Você deve incluir a configuração para criar um servidor virtual local, que irá atender essa API. (ok local)
* um pequeno doc da aplicação
* Utilize uma imagem padrão, só com o sistema operacional do tipo Linux, que deve ser baixada no momento de configuração. (como deseja fazer isto ((colocar a api em um container docker)


### Install Mongo and Postgresql

```
Install [https://docs.mongodb.com/v3.4/tutorial/install-mongodb-on-ubuntu/]

Create Database Mongo and configure

$ mongo 

$ use gmongocrud

$ db

$ db.createCollection("tuser")

$ db.createCollection("tdadosuser")

$ show collections

```

### Install Dependencies

```
go get -v gopkg.in/mgo.v2/bson

go get -v github.com/satori/go.uuid

go get -v github.com/jeffotoni/gmongocrud

```

### Start App with Run or Compile

```
go run main.go 

go build main.go

./main

```

### Structure of a Project

```
/conf 
Application configuration including environment-specific configs

/conf/app
Middlewares and routes configuration

/handler
HTTP handlers

/locale
Language specific content bundles

/lib
Common libraries to be used across your app

/model
Models

/public
Web resources that are publicly available

/public/templates
Jade templates

/repository
Database comunication following repository pattern

main.go
Application entry
```


### Routes 

```

# test the server

- POST /v1/public/ping - Test API Rest

# Login User

- POST /v1/user/login - Logar 

# Data Base using Mongo

- POST /v1/curriculum 	- Adiciona novo curriculo

- GET /v1/curriculum/12 	- Recupera os dados do seu curriculo

- PUT /v1/curriculum/1 	- Atualiza os dados do curriculo

```

### Example Curl - POST /v1/public/ping

```

curl -X POST localhost:8080/v1/public/ping

```

### Response

A successful request returns pong

```
pong

```

### Example Curl - POST /v1/user/login

```

curl -v -X POST localhost:8080/v1/user/login \
-H "X-Key: NWUzOWU3MzY3ZDU4OWRhOWYyY2U0ZGQ1OTRhY2UyNTU=" \
-d "user=email@server.com" \
-d "password=1234"

```

### Example Curl - POST /v1/curriculum

```

curl -v -X POST localhost:8080/v1/curriculum \
-H "Content-Type: application/json" \
-d @curriculums.json

```
### OR

```

curl -X POST localhost:8080/v1/curriculum \
-H "Content-Type: application/json" \
-d '{"nome":"nome aqui","cpf":"xxxx-xxxxx","rg":"xxx.xxx","idade":"xx.xx.xx","bio":"xxxxx.xxx..xx","skill":"xxx.xx.x.xx.x"}'

```

### Sample Response

A successful request returns the HTTP 200 OK status code and a JSON response body.

```

{
	"status'":"ok",
	"msg":"seus dados foram inseridos com sucesso!", 
	"uuid":"238d377b-5db7-45be-ba71-858bed7b4cfc"
}

```

### Example Curl - GET /v1/curriculum/1235

```

curl -X GET localhost:8080/v1/curriculum/1235

```

### Sample Response

A successful request returns the HTTP 200 OK status code and a JSON response body.

```

{
	"status":"ok",
	"msg":"Encontrou o id na base de dados!", 
	"data":
		"{
			"nome":"nome aqui",
			"cpf":"xx-xx-xxx",
			"rg":"xxx-x-xx-",
			"idade":"xx.xxx.x",
			"bio":"xxxx.x.xx",
			"skill":"xxxx"
		}"
}

```
