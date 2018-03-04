/*
* Project Go Library (C) 2017 Inc.
*
* @project     Project
* @package     main
* @author      @jeffotoni
* @size        01/06/2017

* @description Our main auth will be responsible for validating our
* handlers, validating users and will also be in charge of creating users,
* removing them and doing their validation of access.
* We are using jwt to generate the tokens and validate our handlers.
* The logins and passwords will be in the AWS Dynamond database.
*
* $ openssl genrsa -out private.rsa 1024
* $ openssl rsa -in private.rsa -pubout > public.rsa.pub
*
 */

/**
*
Types of handlers authentication:

1 => AuthBasicKey

This method will validate the header containing the X-key, this key is generated in the application that server the front-end, html application and js.
The key is an algorithm available from the api itself to talk to each other.
Every request from the public part of the site will come to this key.
It can be a temporary key or not.

Affected Methods


2 => AuthBasicJwt

This method will validate the X-Key header and the login and password from a form of application / x-www-form-urlencoded or application / json
Validation will be done in the login database and password, which will work fine, the system generates a token with an expiration date that is sent to the user.
The user will receive the access token and the expiration date.
With this token the user can access all the handlers of the admin environment, that is after login.

Affected Methods

/login

3 => AuthBasicHandler

This method is responsible for validating internal or logged in handlers

/hello
/curriculum

The Ping method is the only open method, it does not need any type of authentication, it is to perform tests in the Availability API.
It has a rate limit of access per second like all methods.

*/

package gjwt

import (
	"crypto/rsa"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeffotoni/gmongocrud/lib/context"
	"github.com/jeffotoni/gmongocrud/lib/models"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const HoursExpires = 240

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey

	//pathPrivate = "../certs/private.rsa"
	//pathPublic  = "../certs/public.rsa.pub"

	ProjectTitle = "jwt project"

	//
	// md5(basicHash) BD
	//
	HEDER_X_KEY = "5e39e7367d589da9f2ce4dd594ace255"

	//
	// Base64 KEY
	//
	HEDER_X_KEY_B64 = "NWUzOWU3MzY3ZDU4OWRhOWYyY2U0ZGQ1OTRhY2UyNTU="

	//
	// md5(basicHash) BD
	//
	HEDER_X_KEY_RESTORE_ACCOUNT = "ffe9d2024b009238a2794fcbd5030c61"

	// Base64 KEY
	HEDER_X_KEY_RESTORE_ACCOUNT_B64 = "ZmZlOWQyMDI0YjAwOTIzOGEyNzk0ZmNiZDUwMzBjNjE="
)

// discontinued, we will not use
// this more expensive implementation.
func GetDirKeys() (string, string) {

	var pathApp string

	pwd, err := os.Getwd()

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	VetDir := strings.Split(pwd, "/")

	lenght := len(VetDir) - 2

	for i, path := range VetDir {

		if lenght != i {

			pathApp += fmt.Sprintf("%v%s", path, "/")

		} else {

			break
		}
	}

	pathApp = pathApp + "" + "conf/certs/private.rsa"
	pathApp2 := pathApp + "" + "conf/certs/public.rsa.pub"

	//fmt.Println(pathApp)

	return pathApp, pathApp2
}

//
// Structure of our server configurations
//
type JsonMsg struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

//
// jwt init
//
func init() {

	var errx error

	privateByte := []byte(RSA_PRIVATE)

	privateKey, errx = jwt.ParseRSAPrivateKeyFromPEM(privateByte)

	if errx != nil {

		// WriteJson("error", "Could not parse privatekey!")
		return
	}

	publicByte := []byte(RSA_PUBLIC)

	//
	//
	//
	publicKey, errx = jwt.ParseRSAPublicKeyFromPEM(publicByte)

	if errx != nil {

		// WriteJson("error", "ould not parse publickey!")
		return
	}
}

//
// jwt 'GenerateJWT'
//
func GenerateJWT(model models.User) (string, string) {

	//
	// Generating date validation to return to the user
	//
	Expires := time.Now().Add(time.Hour * HoursExpires).Unix()

	//
	// convert int64
	//
	ExpiresInt64, _ := strconv.ParseInt(fmt.Sprintf("%v", Expires), 10, 64)

	//
	// convert time unix to Date RFC
	//
	ExpiresDateAll := time.Unix(ExpiresInt64, 0)

	//
	// Date
	//
	ExpiresDate := ExpiresDateAll.Format("2009-09-02")

	//
	// claims Token data, the header
	//
	claims := models.Claim{

		User: model.Login,

		Uid: model.Uid,

		Uidwks: model.Uidwks,

		StandardClaims: jwt.StandardClaims{

			//
			// Expires in 24 hours * 10 days
			//
			ExpiresAt: Expires,
			Issuer:    ProjectTitle,
		},
	}

	//
	// Generating token
	//
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	//
	// Transforming into string
	//
	tokenString, err := token.SignedString(privateKey)

	if err != nil {

		return "Could not sign the token!", "2006-01-02"
	}

	//
	// return token string
	//
	return tokenString, ExpiresDate
}

//
// base64 (md5(key))
//
// login e password default in base 64
// curl -X POST -H "Content-Type: application/json"
// -H "Authorization: Basic ZTg5NjFlZDczYTQzMzE0YWYyY2NlNDdhNGY1YjY1ZGI=:ZGExMjRhMDAwNTE1MDUyYzFlNWJjNmU0NzQ4Yzc3ZTU="
// "https://localhost:8181/token"
//
func AuthBasicJwt(ctx *context.Context) bool {

	//
	// do it now
	//
	if AuthBasicKey(ctx) {

		return true

	} else {

		return false
	}
}

//
//
//
func TokenClaimsJwt(ctx *context.Context) string {

	return (GetSplitTokenJwt(ctx))
}

//
//
//
func ClaimsData(token string) (string, string, string, string) {

	if token != "" {
		// start
		parsedToken, err := jwt.ParseWithClaims(token, &models.Claim{}, func(*jwt.Token) (interface{}, error) {

			return publicKey, nil

		})

		if err != nil || !parsedToken.Valid {

			//HttpWriteJson(w, r, "error", "Your token has expired!", http.StatusAccepted)
			return "", "", "", ""
		}

		claims, ok := parsedToken.Claims.(*models.Claim)

		if !ok {

			//HttpWriteJson(w, r, "error", "Token has something wrong!", http.StatusAccepted)
			return "", "", "", ""
		}

		// review implementation, performance
		// validate user before, is there a user?
		//if !DynamodbValidUserAccess(claims.User) {
		//if !ValidateUser(claims.User) {
		if false {
			//msgTmp = `this user [` + claims.User + `] no longer exists, check out your access token`
			//return false, `{"status":"error","msg":"` + msgTmp + `"}`, msgTmp
			return "", "", "", ""
		}
		//review implementation, performance

		return claims.User, claims.Uid, claims.Uidwks, fmt.Sprintf("%v", claims.ExpiresAt)

	} else {

		return "", "", "", ""
	}
}

// TokenGlobal = token
// UserGlobal = claims.User
// ExpiGlobal = fmt.Sprintf("%d", claims.ExpiresAt)
// fmt.Println("User: ", claims.User)
// func2(w, r)
func GetSplitTokenJwt(ctx *context.Context) string {

	var Authorization string

	//Authorization = ctx.Req.Header.Get("Authorization")
	Authorization = ctx.Req.Header.Get("Authorization")

	if Authorization == "" {

		Authorization = ctx.Req.Header.Get("authorization")
	}

	// browsers
	if Authorization == "" {

		Authorization = ctx.Req.Header.Get("Access-Control-Allow-Origin")
	}

	//fmt.Println(ctx.Req.Header.Get("Access-Control-Request-Headers"))
	//fmt.Println(ctx.Req.Header.Get("bearer"))
	//fmt.Println("auth: ", Authorization)

	if Authorization != "" {

		auth := strings.SplitN(Authorization, " ", 2)

		if len(auth) != 2 || strings.TrimSpace(strings.ToLower(auth[0])) != "bearer" {

			//HttpWriteJson(w, r, "error", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return ""
		}

		token := strings.Trim(auth[1], " ")
		token = strings.TrimSpace(token)

		return token
	} else {

		return ""
	}
}

//
//
//
func TokenJwtClaimsValid(ctx *context.Context) bool {

	token := GetSplitTokenJwt(ctx)

	if token != "" {

		// fmt.Println(token)
		// start
		parsedToken, err := jwt.ParseWithClaims(token, &models.Claim{}, func(*jwt.Token) (interface{}, error) {

			return publicKey, nil

		})

		if err != nil || !parsedToken.Valid {

			// HttpWriteJson(w, r, "error", "Your token has expired!", http.StatusAccepted)
			return false
		}

		_, ok := parsedToken.Claims.(*models.Claim)

		// fmt.Println("ok", ok)
		return ok

	} else {

		return false
	}
}

//
// ValidateHandler
//
func ValidateHandler(ctx *context.Context) bool {

	if !TokenJwtClaimsValid(ctx) {

		//w.WriteHeader(http.StatusAccepted)
		// fmt.Fprintln(w, r, "There's something strange about your token!")
		//HttpWriteJson(w, r, "error", "There's something strange about your token!", http.StatusAccepted)
		return false
	}

	return true
}

//
// Returns json without typing in http
//
func WriteJson(Status string, Msg string) {

	msgJsonStruct := &JsonMsg{Status, Msg}

	msgJson, errj := json.Marshal(msgJsonStruct)

	if errj != nil {

		fmt.Println(`{"status":"error","msg":"We could not generate the json error!"}`)
		return
	}

	fmt.Println(msgJson)
}

//
//
//
func GetXKey(ctx *context.Context) string {

	//
	//
	// 'X-Key: YOUR-API-KEY-HERE'
	//
	auth := strings.SplitN(ctx.Req.Header.Get("X-Key"), " ", 2)

	//
	//
	//
	if len(auth) <= 0 {

		//HttpWriteJson(w, r, "error", "Your X-Key it's wrong!", http.StatusAccepted)
		return ""
	}

	//
	//
	//
	tokenBase64 := strings.Trim(auth[0], " ")

	//
	//
	//
	tokenBase64 = strings.TrimSpace(tokenBase64)

	//
	// User, Login byte
	//
	tokenUserDecode, _ := b64.StdEncoding.DecodeString(tokenBase64)

	//
	// User, Login string
	//
	tokenUserDecodeS := strings.TrimSpace(strings.Trim(string(tokenUserDecode), " "))

	return tokenUserDecodeS
}

//
//
//
func GetXKeyUrl(ctx *context.Context) string {

	//
	//
	// 'X-Key: YOUR-API-KEY-HERE'
	//
	//auth := r.FormValue("k")

	auth := ctx.Req.Form.Get("k")

	//
	//
	//
	if auth == "" {

		//HttpWriteJson(w, r, "error", "Your X-Key it's wrong!", http.StatusAccepted)
		return ""
	}

	//
	//
	//
	tokenBase64 := strings.Trim(auth, " ")

	//
	//
	//
	tokenBase64 = strings.TrimSpace(tokenBase64)

	//
	// User, Login byte
	//
	tokenUserDecode, _ := b64.StdEncoding.DecodeString(tokenBase64)

	//
	// User, Login string
	//
	tokenUserDecodeS := strings.TrimSpace(strings.Trim(string(tokenUserDecode), " "))

	return tokenUserDecodeS
}

//
// base64 (md5(key))
//
// login e password default in base 64
// curl -X POST -H "Content-Type: application/json"
// -H "Authorization: Basic ZTg5NjFlZDczYTQzMzE0YWYyY2NlNDdhNGY1YjY1ZGI=:ZGExMjRhMDAwNTE1MDUyYzFlNWJjNmU0NzQ4Yzc3ZTU="
// "https://localhost:9001/token"
//
func AuthBasicKey(ctx *context.Context) bool {

	tokenUserDecodeS := GetXKey(ctx)

	//
	// Validate user and password in the database
	//
	if tokenUserDecodeS == HEDER_X_KEY {

		return true

	} else {

		//HttpWriteJson(w, r, "error", "Your X-Key it's wrong ...", http.StatusAccepted)
		return false
	}

	//defer r.Body.Close()
	//HttpWriteJson(w, r, "error", "Your X-Key it's wrong ..", http.StatusAccepted)
	return false
}

//
// base64 (md5(key))
//
// login e password default in base 64
// curl -X POST -H "Content-Type: application/json"
// -H "Authorization: Basic ZTg5NjFlZDczYTQzMzE0YWYyY2NlNDdhNGY1YjY1ZGI=:ZGExMjRhMDAwNTE1MDUyYzFlNWJjNmU0NzQ4Yzc3ZTU="
// "https://localhost:9001/token"
//
func GetAuthBasicKey(ctx *context.Context) bool {

	tokenUserDecodeS := GetXKeyUrl(ctx)

	//
	// Validate user and password in the database
	//
	if tokenUserDecodeS == HEDER_X_KEY {

		return true

	} else {

		//HttpWriteJson(w, r, "error", "Your X-Key it's wrong ...", http.StatusAccepted)
		return false
	}

	//defer r.Body.Close()
	//HttpWriteJson(w, r, "error", "Your X-Key it's wrong ..", http.StatusAccepted)
	return false
}

//
// base64 (md5(key))
//
// login e password default in base 64
// curl -X POST -H "Content-Type: application/json"
// -H "Authorization: Basic ZTg5NjFlZDczYTQzMzE0YWYyY2NlNDdhNGY1YjY1ZGI=:ZGExMjRhMDAwNTE1MDUyYzFlNWJjNmU0NzQ4Yzc3ZTU="
// "https://localhost:9001/token"
//
func GetAuthBasicKeyRestoreAccount(ctx *context.Context) bool {

	tokenUserDecodeS := GetXKeyUrl(ctx)

	//
	// Validate user and password in the database
	//
	if tokenUserDecodeS == HEDER_X_KEY_RESTORE_ACCOUNT {

		return true

	} else {

		//HttpWriteJson(w, r, "error", "Your X-Key it's wrong ...", http.StatusAccepted)
		return false
	}

	//defer r.Body.Close()
	//HttpWriteJson(w, r, "error", "Your X-Key it's wrong ..", http.StatusAccepted)
	return false
}

//
// Returns json by typing on http
//
func HttpWriteJsonNew(ctx *context.Context, Status string, Msg string, httpStatus int) {

	//
	//
	//
	t1 := time.Now()

	//
	//
	//
	msg := Msg

	//
	//
	//
	msgJsonStruct := &JsonMsg{Status, Msg}

	//
	//
	//
	msgJson, errj := json.Marshal(msgJsonStruct)

	//
	//
	//
	if errj != nil {

		//ctx.Resp.WriteHeader(http.StatusUnauthorized)

		msgJson := `{"status":"error","msg":"We could not generate the json error!"}`
		ctx.JSON(http.StatusOK, msgJson)

		//fmt.Fprintln(w, `{"status":"error","msg":"We could not generate the json error!"}`)
		return
	}

	//
	//
	//
	//w.WriteHeader(httpStatus)
	ctx.Resp.WriteHeader(httpStatus)

	//
	//
	//
	// w.Header().Set("Content-Type", "application/json")
	ctx.Resp.Header().Set("Content-Type", "application/json")

	//
	//
	//
	//w.Write(msgJson)
	ctx.JSON(http.StatusOK, msgJson)

	//
	//
	//
	t2 := time.Now()

	//
	//
	//
	// LogHandlerOff(w, r, msg, t1, t2)
	fmt.Println("Error {HttpWriteJsonNew} ", msg, " ", t1, " ", t2)
}
