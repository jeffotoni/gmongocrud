# gmongocrud
This api is developed in the framework macaron and the Mercurius, is an abstraction to register curriculum with upload.
It's just a simple example using the framework to do an insert, list and update curriculum in a mongodb database.

### Install Mongo and Postgresql

```
Install [https://docs.mongodb.com/v3.4/tutorial/install-mongodb-on-ubuntu/]

Create Database Mongo and configure

$ mongo 

$ use gmongocrud

$ db

$ db.createCollection("curriculum")

$ show collections

```

### Install Dependencies

```

$ go get github.com/satori/go.uuid

$ go get gopkg.in/macaron.v1

$ go get -u github.com/go-macaron/macaron

$ go get github.com/go-macaron/binding

$ go get github.com/go-macaron/cache

$ go get github.com/go-macaron/cache/memcache

$ go get github.com/go-macaron/cache/redis

$ go get github.com/go-macaron/gzip

$ go get github.com/go-macaron/i18n

$ go get github.com/go-macaron/jade

$ go get github.com/go-macaron/session

$ go get github.com/go-macaron/toolbox

$ go get github.com/jmoiron/sqlx

$ go get -u github.com/didip/tollbooth

$ go get -v gopkg.in/mgo.v2/bson

$ go get -v github.com/satori/go.uuid

$ go get -u golang.org/x/crypto/bcrypt

$ go get -u github.com/dgrijalva/jwt-go

$ go get -u github.com/felipeweb/gopher-utils

$ go get -u github.com/go-sql-driver/mysql

$ go get -u github.com/lib/pq

```

### Start App with Run or Compile

```
go run main.go 

go build main.go

./main

```

# Docker

docker [Installing] (Docker https://docs.docker.com/engine/installation)

# Pull in the image

There is another possibility, which is of the pull in a complete image

```sh

$ sudo docker pull jeffotoni/gmongocrud:latest

```

# Start image and creating container

```sh

$ sudo docker run -e mongo_uri='mongodb://192.168.0.13:27017' -it --rm --name gmongocrud -p 8181:8181 -v /tmp/upload:/go/upload  jeffotoni/gmongocrud:latest

```

# Reverse proxy nginx

```

server {

    listen 80;
    server_name avaliacao.com;
    return 301 https://$host$request_uri;
}

server {

	listen 443;
	server_name avaliacao.com;

	ssl_certificate          /etc/nginx/ssl/avaiacao/2017/avaliacao.crt;
	ssl_certificate_key      /etc/nginx/ssl/avaliacao/2017/avaliacao.key;

	ssl on;
	ssl_session_cache  builtin:1000  shared:SSL:10m;
	ssl_protocols  TLSv1 TLSv1.1 TLSv1.2;
	ssl_ciphers HIGH:!aNULL:!eNULL:!EXPORT:!CAMELLIA:!DES:!MD5:!PSK:!RC4;
	ssl_prefer_server_ciphers on;

	access_log /var/log/nginx/avaliacao-acesso.log;
	error_log /var/log/nginx/avaliacao-acesso-error.log;

	location / {

	  proxy_set_header        Host $host;
	  proxy_set_header        X-Real-IP $remote_addr;
	  proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
	  proxy_set_header        X-Forwarded-Proto $scheme;

	  # Fix the “It appears that your reverse proxy set up is broken" error.
	  proxy_pass          http://localhost:8181;

	  proxy_read_timeout  90;

	}
}

```

### Generate api documentation swagger

$ sudo apt install python3-pip

$ sudo pip install pyaml

$ python swagger-yaml-to-html.py < swagger.yalm> index.html

### Structure of a Project

```
/docapi
Api documentation

/conf 
Application configuration including environment-specific configs

/conf/app
Middlewares and routes configuration

/conf/certs
Ssl keys for token generation

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

/upload
Where to save the upload files

```


### Routes 

```

# test the server

- POST /v1/public/ping - Test API Rest

# Token User

- POST /v1/user/token - Logar 

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

### Example Curl - POST /v1/user/token

```

curl -v -X POST localhost:8181/v1/user/token \
-H "X-Key: NWUzOWU3MzY3ZDU4OWRhOWYyY2U0ZGQ1OTRhY2UyNTU=" \
-d "user=avaliacao"
```

# Resposta

```
{
	"status":"o",
	"msg":"success",
	"token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.......",
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

