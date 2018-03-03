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
	"net/http"
)

func Login(ctx *context.Context) {

	msgJson := `{"msg":"login..."}`
	ctx.JSON(http.StatusOK, msgJson)
}
