/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package repo

import (
	"os"
	"strings"
)

// criando path absoluto apartir
// de onde encontra-se a executacao
// do script
func GetWdLocal(level int) string {

	pathNow, _ := os.Getwd()

	pathV := strings.Split(pathNow, "/")

	pathNew := pathV[:len(pathV)-level]

	pathNewOrg := strings.Join(pathNew, "/")

	return pathNewOrg
}

// checando se existe o path
func GetExistFile(pathNewOrg string) string {

	_, err := os.Stat(pathNewOrg)

	if err == nil {

		return pathNewOrg

	} else if os.IsNotExist(err) {

		return ""

	} else {

		return ""
	}
}

// criando o path absoluto
func TmplSqlGwd(NameTmpl string, level int) string {

	pathNewOrg := GetWdLocal(level)

	pathNewOrg = pathNewOrg + "/sqltabletmpl/" + NameTmpl

	return GetExistFile(pathNewOrg)
}
