/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package handler

import (
	"github.com/jeffotoni/gmongocrud/lib/context"
	"log"
	"net/http"
)

// inserindo perguntas na base de dados
func HandlerError(ctx *context.Context) {

	log.Println("Error ...")

	jsonString := `{"status":"error","msg":"Error ao validar o handler, confira seu token de acesso!"}`
	ctx.Resp.Header().Set("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, jsonString)
}
