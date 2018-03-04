/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package gjwt

import (
	"encoding/json"
	"fmt"
	"github.com/jeffotoni/gmongocrud/lib/context"
	"github.com/jeffotoni/gmongocrud/repo"
	//"io/ioutil"
	//"net/http"
	"strings"
)

type UserValid struct {
	Email string `json:"user"`
	//Password string `json:"password"`
	Password json.Number `json:"password,Number"`
}

//
// Create User Form
//
func ValidateUser(ctx *context.Context) (bool, string, string, string) {

	//
	//
	//
	var Email, Password2 string

	Password := UserValid{}

	// just scape
	fmt.Sprint("%v", Password)

	typeHeader := ValidHeader(ctx)

	//
	//
	//
	if typeHeader == "JSON" {

		//
		//
		//

		data, _ := ctx.Req.Body().Bytes()
		//fmt.Print(data)

		var L UserValid

		//
		//
		//
		err := json.Unmarshal(data, &L)

		if err != nil {

			return false, `{"status":"error","msg":"Json did not work: ` + fmt.Sprintf("%s", err) + `"}`, "", ""
		}

		Email = L.Email
		Password := L.Password
		Password2 = fmt.Sprintf("%s", Password)

	} else if typeHeader == "FORM" {

		// Email
		// Password
		Email = ctx.Req.Form.Get("user")
		Password := ctx.Req.Form.Get("password")
		Password2 = fmt.Sprintf("%s", Password)

	} else {

		//error
		return false, `{"status":"error","msg":"Invalid Content-Type"}`, "", ""
	}

	Password2 = strings.TrimSpace(strings.Trim(Password2, " "))

	//
	// emails valid
	//
	Email = strings.ToLower(strings.TrimSpace(strings.Trim(Email, " ")))

	if Email == "" {

		return false, `{"status":"error","msg":"Empty field [User]"}`, "", ""

	} else if Password2 == "" {

		return false, `{"status":"error","msg":"Empty field [Password]"}`, "", ""
	}

	if len(Email) >= 201 {

		return false, `{"status":"error","msg":"Very large size, allowed 200 characters"}`, "", ""

	} else if len(Password2) >= 101 {

		return false, `{"status":"error","msg":"Very large size, allowed 100 characters"}`, "", ""
	}

	//
	// Is the user active?
	//
	//if !DynamodbValidUserAccess(Email) {
	//if !PgUserValid(Email) {

	if false {

		return false, `{"status":"error","msg":"Wrong user is not active!"}`, "", ""
	}

	//
	// Validate if the user exists
	//
	//uidUser, uidWrks := PgAuthUser(Email, Password2)
	uidUser, uidWrks := repo.MongoAuthUser(Email, Password2)

	if uidUser == "" {

		return false, `{"status":"error","msg":"Wrong password, try again!"}`, "", ""
	}

	return true, Email, uidUser, uidWrks
}

//
//
//
func ValidHeader(ctx *context.Context) string {

	//
	// Accept
	//
	contentType := ctx.Req.Header.Get("Content-Type")

	tmpContent := strings.ToLower(strings.TrimSpace(contentType))
	tmpContentV := strings.Split(tmpContent, ";")

	if tmpContentV[0] == "application/x-www-form-urlencoded" {

		return "FORM"

	} else if strings.ToLower(strings.TrimSpace(contentType)) == "application/x-www-form-urlencoded" {

		return "FORM"

	} else if strings.ToLower(strings.TrimSpace(contentType)) == "application/json" {

		return "JSON"

	} else {

		return "Error"
	}
}

func ValidJson(bodyJson []byte) (bool, string) {

	if len(bodyJson) == 0 {

		return true, `{"status":"error","msg":"Missing Json"}`

	}

	//
	// Looking for keys in the first and last position
	//
	last_pos := len(bodyJson) - 1

	//
	//
	//
	if string(bodyJson[0]) != "{" {

		return true, `{"status":"error","msg":"Missing keys on your json '{'"}`

	} else if string(bodyJson[last_pos]) != "}" {

		return true, `{"status":"error","msg":"Missing keys on your json '}'"}`

	} else {

		return false, ""
	}
}
