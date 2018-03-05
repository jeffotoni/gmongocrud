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
	"github.com/jeffotoni/gmongocrud/lib/gjwt"
	"github.com/jeffotoni/gmongocrud/lib/models"
	"log"
	"net/http"
	"time"
)

// Testing whether the service is online
//
func Login(ctx *context.Context) {

	TokenLocal := ""

	t1 := time.Now()

	statusOk := false

	ok, emailUser, uidUser, uidWks := gjwt.ValidateUser(ctx)

	//log.Println("ok: ", ok)

	if ok {

		// do it now
		var model models.User

		model.Login = emailUser

		model.Uid = uidUser

		model.Uidwks = uidWks

		model.Password = ""

		model.Role = "user-default"

		token, expires := gjwt.GenerateJWT(model)

		if token == "" || expires == "" {

			jsonString := `{"status":"error","msg":"Token error generating!"}`
			//jsonByte := []byte()

			//w.WriteHeader(http.StatusOK)
			ctx.Resp.WriteHeader(http.StatusOK)

			//w.Header().Set("Content-Type", "application/json")
			ctx.Resp.Header().Set("Content-Type", "application/json")
			//ctx.Resp.Write(jsonByte)
			ctx.JSON(http.StatusOK, jsonString)
			//w.Write(jsonByte)

		} else {

			//send email emailUser
			// write client
			TokenLocal = token
			statusOk = true

			jsonString := `{"status":"ok","msg":"success","token":"` + token + `","expires":"` + expires + `"}`

			//jsonByte := []byte(msgJsonx)

			//w.WriteHeader(http.StatusOK)
			ctx.Resp.WriteHeader(http.StatusOK)
			ctx.Resp.Header().Set("Content-Type", "application/json")
			//ctx.Resp.Write(jsonByte)
			ctx.JSON(http.StatusOK, jsonString)

			//w.Header().Set("Content-Type", "application/json")
			//w.Write(jsonByte)
		}

	} else {

		// Can contain a json as a message when
		// it gives some kind of error
		//jsonByte := []byte(emailUser)

		//w.WriteHeader(http.StatusOK)
		//w.Header().Set("Content-Type", "application/json")
		//w.Write(jsonByte)

		ctx.Resp.WriteHeader(http.StatusOK)
		ctx.Resp.Header().Set("Content-Type", "application/json")
		//ctx.Resp.Write(jsonByte)
		ctx.JSON(http.StatusOK, emailUser)
	}

	//
	//
	//
	t2 := time.Now()

	if statusOk {

		msg := "login success"
		//LogHandlerLoginOn(w, r, msg, t1, t2, TokenLocal)
		log.Println(msg, t1, t2, TokenLocal)

	} else {

		msg := "Error while trying to login"
		log.Println(msg, t1, t2, TokenLocal)
		//LogHandlerOff(w, r, msg, t1, t2)
	}
}
