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
	"github.com/jeffotoni/gmongocrud/repo"
	"log"
	"net/http"
)

func Ping(ctx *context.Context) {

	pathNewOrg := repo.GetWdLocal(0)

	log.Println("pring path: ", pathNewOrg)

	msgJson := `{"msg":"pong..."}`
	ctx.JSON(http.StatusOK, msgJson)
}
