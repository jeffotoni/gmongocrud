/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jeffotoni/gmongocrud/conf"
	"github.com/jeffotoni/gmongocrud/model"
	uuid "github.com/satori/go.uuid"
	bson "gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
	"time"
)

// fazendo um insert em um collection especifica
// o metodo retorna o Uuid caso tenha sucesso
// ou error caso ocorra algum imprevisto
func Insert(namecollection string, insert model.TPesqCurriculum) (Uuid string, err error) {

	// criando session e retorno o db do collection
	col, err := conf.GetMongoCollection(namecollection)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// garantindo que sera fechado
	// a session do banco assim
	// que finalizar
	defer col.Database.Session.Close()

	// criando uuid
	// para salvar
	// na base
	u1 := uuid.NewV4()

	// convertendo uuid para string
	insert.Ppr_uuid = fmt.Sprintf("%s", u1)

	DataNow := time.Now().Format(conf.LayoutDate)
	HourNow := time.Now().Format(conf.LayoutHour)

	insert.Ppr_datetime = DataNow + " " + HourNow

	// salvando na base de
	// dados o insert do
	// collection
	err = col.Insert(insert)

	// checando
	// se tudo
	// ocorreu bem
	if err != nil {
		log.Println(err.Error())
		return
	}

	// se tudo ocorreu bem
	// retorna Uuid para o usuario
	Uuid = string(insert.Ppr_uuid)
	return
}

// find Exist
func FindExist(namecollection string, Jcondition bson.M) (exist bool, err error) {

	// criando session e retorno o db do collection
	col, err := conf.GetMongoCollection(namecollection)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//busca na base se existir os mesmos dados nao deixa inserir
	// Find and Count
	tpesqcount, err := col.Find(Jcondition).Count()

	if err != nil {

		exist = false
		log.Println(err.Error())
		return
	}

	if tpesqcount > 0 {

		exist = true
		return
	}

	exist = false
	return
}

// remove registro
func Remove(namecollection string, Jcondition bson.M) (err error) {

	// criando session e retorno o db do collection
	col, err := conf.GetMongoCollection(namecollection)

	if err != nil {
		log.Println(err.Error())
		return
	}

	// Remove
	info, err := col.RemoveAll(Jcondition)

	if err != nil {
		log.Println(err.Error())
		return
	}

	if info.Removed == 0 {
		err = errors.New("Error ao remover no mongo o registro")
		log.Println(err.Error())
		return
	}

	err = nil
	return
}

// remove registro
func Update(namecollection string, Jcondition bson.M, SetField bson.M) (err error) {

	// criando session e retorno o db do collection
	col, err := conf.GetMongoCollection(namecollection)

	// conferindo os erros
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Remove
	info, err := col.UpdateAll(Jcondition, SetField)

	if err != nil {
		log.Println(err.Error())
		return
	}

	// testando se houve
	// atualizacao
	if info.Updated == 0 {

		err = errors.New("Error ao atualizar no mongo o registro")
		log.Println(err.Error())
		return

	}

	//sucesso
	err = nil
	return
}

// find registro
func Find(namecollection, Uuid string) (jsonStr string, err error) {

	// criando session e retorno o db do collection
	col, err := conf.GetMongoCollection(namecollection)

	// collection que sera criado
	// para busca
	var collection model.TPesqCurriculum

	// fazendo uma busca para buscar um registro
	err = col.Find(bson.M{"ppr_uuid": Uuid}).Select(nil).One(&collection)

	// checando se
	// tudo correu bem
	if err != nil {

		log.Println(err.Error())
		return
	}

	// criando mapa para fazer json com
	// marshal
	mapPesq := map[string]string{
		"ppr_cod":        strconv.Itoa(collection.Ppr_cod),
		"ppr_ppq_cod":    strconv.Itoa(collection.Ppr_ppq_cod),
		"ppr_per_cod":    strconv.Itoa(collection.Ppr_per_cod),
		"ppr_ordem":      strconv.Itoa(collection.Ppr_ordem),
		"ppr_dtcadastro": collection.Ppr_dtcadastro,
		"ppr_dtaltera":   collection.Ppr_dtaltera,
	}

	// convertendo para json
	jsonPesq, err := json.Marshal(mapPesq)

	// checando se
	// tudo correu bem
	// na conversao
	// do json
	if err != nil {

		log.Println(err.Error())
		return
	}

	jsonStr = string(jsonPesq)
	return
}

// find all registro
func FindAll(namecollection string) (jsonStr string, err error) {

	// return
	var jsonPesq []byte

	// criando session e retorno o db do collection
	col, err := conf.GetMongoCollection(namecollection)

	// collection que sera criado
	// para busca
	var collection []model.TPesqCurriculum

	// Find All
	err = col.Find(nil).Sort("-start").All(&collection)

	if err != nil {
		log.Println(err.Error())
		return
	}

	// caso encontre registros
	// faz o marshal para json
	if len(collection) > 0 {

		// convertendo para json
		jsonPesq, err = json.Marshal(collection)

		// checando se
		// tudo correu bem
		// na conversao
		// do json
		if err != nil {

			log.Println(err.Error())
			return
		}
	} else {

		msgerror = "Nao encontramos o resultado"
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}

	// recebendo o json
	// caso ocorra tudo certo
	jsonStr = string(jsonPesq)
	return
}

func MongoAuthUser(user, password string) (iduser, idworks string) {

	return
}
