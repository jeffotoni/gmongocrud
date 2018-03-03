/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package handler

import (
	"github.com/jeffotoni/mercuriuscrud/lib/context"
	"net/http"
)

func Hello(ctx *context.Context) {

	msgJson := `{"msg":"Hello, Handler Works!"}`
	ctx.JSON(http.StatusOK, msgJson)
}
