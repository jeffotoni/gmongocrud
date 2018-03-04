/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package handler

import (
	"github.com/jeffotoni/gmongocrud/conf"
	"github.com/jeffotoni/gmongocrud/lib/context"
	"github.com/jeffotoni/gmongocrud/repo"
	"log"
	"net/http"
	"strings"
)

var msgerror string

// inserindo perguntas na base de dados
func CurriculumCreate(ctx *context.Context) {

	// bytes json body
	var byteJson []byte

	// define
	// var error
	var err error

	// msg json
	var msgJson, Uuid string

	// tipo aceito de content-type
	cTypeAceito := "application/json"

	// capturando content-type
	cType := ctx.Req.Header.Get("Content-Type")

	// validando content-type
	if strings.ToLower(strings.TrimSpace(cType)) == cTypeAceito {

		// capturando json findo da requisicao
		// estamos pegando em bytes para ser
		// usado no Unmarshal que recebe bytes
		byteJson, err = ctx.Req.Body().Bytes()

		// fechando Req.Body
		defer ctx.Req.Body().ReadCloser()

		// caso ocorra
		// erro ao ler
		// envia msg
		// de error
		if err != nil {
			msgerror = "[CurriculumCreate] Erro ao capturar Json: " + err.Error()
			log.Println(msgerror)
			msgJson = `{"status":"error","msg":"` + msgerror + `}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return

		} else {

			// se nao tiver nem um valor
			// no json, emite um erro
			// para o usuario
			if string(byteJson) != "" {

				// Criar registro no banco de dados
				Uuid, err = repo.AddCurriculum(byteJson)

				// tratando o erro
				if err != nil {
					log.Println(err.Error())
					msgJson = `{"status":"error","msg":"` + err.Error() + `"}`
					ctx.JSON(http.StatusUnauthorized, msgJson)
					return
				} else {
					// sucesso
					msgJson = `{"status":"ok","msg":"seus dados foram cadastrados com sucesso!", "uuid":"` + Uuid + `"}`
					ctx.JSON(http.StatusOK, msgJson)
				}
			} else {
				log.Println("[CurriculumCreate] Erro em sua string json nao pode ser vazia!")
				msgJson = `{"status":"error","msg":"Erro em sua string json"}`
				ctx.JSON(http.StatusUnauthorized, msgJson)
				return
			}
		} // fim else
	} else {

		vetCot := strings.Split(cType, ";")
		Content := strings.ToLower(strings.TrimSpace(vetCot[0]))

		if Content == "multipart/form-data" {

			log.Println("teste nome:: ", ctx.Req.Form.Get("nome"))
			log.Println("teste cpf:: ", ctx.Req.Form.Get("cpf"))

			jsonl := `{` +
				`"nome":"` + ctx.Req.Form.Get("nome") + `",` +
				`"cpf":"` + ctx.Req.Form.Get("cpf") + `",` +
				`"rg":"` + ctx.Req.Form.Get("rg") + `",` +
				`"idade":"` + ctx.Req.Form.Get("idade") + `",` +
				`"bio":"` + ctx.Req.Form.Get("bio") + `",` +
				`"skill":"` + ctx.Req.Form.Get("skill") + `"` + `}`

			// log.Println(jsonl)

			byteJson := []byte(jsonl)

			// adicionado curriculo
			Uuid, err = repo.AddCurriculum(byteJson)

			// tratando o erro
			if err != nil {
				log.Println(err.Error())
				msgJson = `{"status":"error","msg":"` + err.Error() + `"}`
				ctx.JSON(http.StatusUnauthorized, msgJson)
				return
			} else {

				// path upload
				pathUpload := "./" + conf.PATH_UPLOAD + "/" + Uuid

				// salvando arquivo em disco
				ctx.SaveToFile(conf.NAME_FORM_FILE, pathUpload)

				// sucesso
				msgJson = `{"status":"ok","msg":"seus dados foram cadastrados com sucesso!", "uuid":"` + Uuid + `"}`
				ctx.JSON(http.StatusOK, msgJson)
			}

		} else {

			log.Println("[CurriculumCreate] Erro Content-Type: aceitamos somente " + cTypeAceito)
			msgJson = `{"status":"error","msg":"error no Content-Type: ` + cType + `, aceitamos somente [Content-Type: ` + cTypeAceito + `]"}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		}
	}
}

// Busca Curriculum especifico na base de dados
func CurriculumFind(ctx *context.Context) {

	// msg json
	var msgJson string

	// chave para atualizacao
	Uuid := ctx.Params(":id")

	if Uuid != "" {

		// para atualizacao temos o nome do collection a chave para efetuar o update e
		// os campose que sera feita o set update
		strJson, err := repo.GetCurriculum(Uuid)

		// testando se tudo
		// correu bem
		if err == nil {

			// Uuid
			msgJson = `{"status":"ok","msg":"Encontrou o id na base de dados!", "data":"` + strJson + `"}`
			// send write to client
			ctx.JSON(http.StatusOK, msgJson)

		} else {
			msgerror = "[CurriculumFind] " + err.Error()
			log.Println(msgerror)
			msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		}

	} else {

		msgerror = "[CurriculumFind] Uuid é obrigatorio!"
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}
}

// Update na base de dados
func CurriculumUpdate(ctx *context.Context) {

	// define err
	var err error

	// mensagem json
	var msgJson string

	// byjson
	var byteJson []byte

	// Chave unica
	Uuid := ctx.Params(":id")

	// testando
	// o Uuid
	if Uuid != "" {

		// capturando json findo da requisicao
		// estamos pegando em bytes para ser
		// usado no Unmarshal que recebe bytes
		byteJson, err = ctx.Req.Body().Bytes()

		if err != nil {
			msgJson = `{"status":"error","msg":"` + err.Error() + `"}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		}

		// fechando Req.Body
		defer ctx.Req.Body().ReadCloser()

		// del question
		err := repo.UpCurriculum(Uuid, byteJson)

		// if tudo correr bem
		// registro foi atualizado
		// com sucesso
		if err == nil {
			// Uuid
			msgJson = `{"status":"ok","msg":"atualizado com sucesso seu Uuid: ` + Uuid + `!"}`
			ctx.JSON(http.StatusOK, msgJson)

		} else {

			log.Println(err.Error())
			msgJson = `{"status":"error","msg":"` + err.Error() + `]"}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		}

	} else {
		msgerror = "[QuestionsUpdate] Uuid é obrigatório!"
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}
}
