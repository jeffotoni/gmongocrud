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

curl -X POST localhost:8181/v1/public/ping

```

### Response

A successful request returns pong

```
pong

```

### Example Curl - POST /v1/user/login

```

curl -v -X POST localhost:8181/v1/user/login \
-H "X-Key: NWUzOWU3MzY3ZDU4OWRhOWYyY2U0ZGQ1OTRhY2UyNTU=" \
-d "user=email@server.com" \
-d "password=1234"

```

# Resposta

```
{
	"status":"o",
	"msg":"success",
	"token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiamVmZi5vdG9uaSIsInVpZCI6IjEyMzQ1NjkzOSIsInVpZHdrcyI6IjM4MzgzODM4MzM4ODMiLCJleHAiOjE1MjEwNDg4NTMsImlzcyI6Imp3dCBwcm9qZWN0In0.wo9fJ0CxBfDlZkwrJql8_3Vjzup1xJhDwYaGxTQbmyug80HPO8uRkvng8ZUKG95hL-ldZCYOe_sdkz2wb7wGcM-8mzFJpRqqgcJeyHBIjSDd4xHQPDyzKF4wZwQ7EgR6C_pUSmnvYRHOcT-FEK8gnqydl8BK0ZvcwWg3zYO7rc8",
	"expires":"14009-09-14"
}

```

### Example Curl - POST /v1/curriculum

```

curl -v -X POST localhost:8181/v1/curriculum \
-H "Content-Type: application/json" \
-H "Authorization: Bearer xxxxxxxxxxxxxxx" \
-d @curriculums.json

```
### OR

```

curl -X POST localhost:8181/v1/curriculum \
-H "Content-Type: application/json" \
-H "Authorization: Bearer xxxxxxxxxxxxxxx" \
-d '{"nome":"nome aqui","cpf":"xxxx-xxxxx","rg":"xxx.xxx","idade":"xx.xx.xx","bio":"xxxxx.xxx..xx","skill":"xxx.xx.x.xx.x"}'

```

### Example Curl Upload File - POST /v1/curriculum

```

curl -v -X POST localhost:8181/v1/curriculum \
-H "Authorization: Bearer xxxxxxxxxxxxxxx" \
--form "nome=seu nome" \
--form "cpf=xx.xxx.xxx-xx" \ 
-F "rg=xxxxxx" \
-F "idade=xx/xx/xxxx" \
-F "bio=Seu Bio aqui" \
-F "skill=Seu Skill aqui" \
-F "file=@seucurriculo.pdf"

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

curl -X GET localhost:8181/v1/curriculum/1235 \
-H "Authorization: Bearer xxxxxxxxxxxxxxx"

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

### Example Curl - PUT /v1/curriculum/1234

```

curl -X PUT localhost:8181/v1/curriculum/1234
-H "Content-Type: application/json" \
-H "Authorization: Bearer xxxxxxxxxxxxxxx" \
-d @curriculums-update.json

```

### Sample Response

A successful request returns the HTTP 200 OK status code and a JSON response body.

```

{"status":"ok","msg":"Atualizado com sucesso!"}

```

