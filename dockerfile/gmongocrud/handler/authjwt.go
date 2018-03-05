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
	//"net/http"
)

type fa func(ctx *context.Context) bool
type fn2 func(ctx *context.Context)

// Function responsible for abstraction and receive the
// authentication function and the handler that will execute if it is true
func HandlerFuncAuth(authJwt fa, handlerContext fn2) fn2 {

	return func(ctx *context.Context) {

		if authJwt(ctx) {

			handlerContext(ctx)

		} else {

			HandlerError(ctx)
		}
	}
}
