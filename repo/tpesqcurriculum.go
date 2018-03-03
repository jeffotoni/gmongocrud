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
	"github.com/jeffotoni/gmongocrud/model"
	bson "gopkg.in/mgo.v2/bson"
	"log"
)

// criando Curriculum
func AddCurriculum(byteJson []byte) (Uuid string, err error) {

	// struct model data
	var Tp model.TPesqCurriculum

	//
	var exist bool

	// convert json to struct
	err = json.Unmarshal(byteJson, &Tp)

	//checar o erro
	if err != nil {
		log.Println("[AddCurriculum] Erro nao consegui converter sua string em json: " + err.Error())
		return
	} else {

		//
		// campos obrigatorios
		//

		if Tp.Ppr_ordem <= 0 {
			msgerror = "[AddCurriculum] Erro A coluna ppr_ordem obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_per_cod <= 0 {
			msgerror = "[AddCurriculum] Erro A coluna Ppr_per_cod obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_cod <= 0 {
			msgerror = "[AddCurriculum] Erro A coluna Ppr_cod obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_ppq_cod <= 0 {
			msgerror = "[AddCurriculum] Erro A coluna Ppr_ppq_cod obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_dtcadastro == "" {
			msgerror = "[AddCurriculum] Erro A coluna Ppr_dtcadastro obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_dtaltera == "" { // campo obrigatorio string
			msgerror = "[AddCurriculum] Erro A coluna Ppr_dtaltera obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else {

			// validar se existe o dado no banco
			// if exist o dado no mongo nao faca
			exist, err = FindExist(tpep_collection, bson.M{"ppr_cod": Tp.Ppr_cod})

			// validar
			if err != nil {
				log.Println("[AddCurriculum] A collection pode esta errada ou sua condicao error: !" + err.Error())
				return
			} else {

				// se existe
				// nao faca
				// o insert
				if exist {
					msgerror = "[AddCurriculum] Error estes dados ja existe na base de dados!"
					err = errors.New(msgerror)
					log.Println(msgerror)
					return
				} else {

					// faca o insert no banco de dados
					// retorne o uuid do salvar
					// caso tenha erro informe ao cliente
					Uuid, err = Insert(tpep_collection, Tp)
					if err != nil {
						log.Println("[AddCurriculum] ocorreu algum erro ao salvar seus dados error: !" + err.Error())
						return
					} else {

						return
					}
				}
			}
		}
	}
}

// deletando Curriculum
func DelCurriculum(Uuid string) (err error) {

	if Uuid == "" {

		msgerror = "[DelCurriculum] Uuid obrigatorio!"
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}

	// chamando o metodo para remover o registro do banco
	err = Remove(tpep_collection, bson.M{"ppr_uuid": Uuid})

	if err != nil {
		msgerror = "[DelCurriculum] Algo errado ocorreu ao remover registro: " + err.Error()
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}

	return
}

// atualizando Curriculum
func UpCurriculum(Uuid string, byteJson []byte) (err error) {

	// struct model data
	var Tp model.TPesqCurriculum

	if Uuid != "" {

		err = json.Unmarshal(byteJson, &Tp)
		if err != nil {
			msgerror = "[UpCurriculum] Algo estranho ocorreu quando tentamos ler seu json!"
			log.Println(msgerror)
			return
		}

		// chave para fazer update
		KeyUp := bson.M{"ppr_uuid": Uuid}

		// campos a serem realizado o update
		SetFields := bson.M{"$set": bson.M{"ppr_cod": Tp.Ppr_cod,
			"ppr_ppq_cod": Tp.Ppr_ppq_cod, "ppr_per_cod": Tp.Ppr_per_cod,
			"ppr_ordem": Tp.Ppr_ordem, "ppr_dtcadastro": Tp.Ppr_dtcadastro,
			"ppr_dtaltera": Tp.Ppr_dtaltera,
		}}

		// para atualizacao temos o nome do collection a chave para efetuar o update e
		// os campose que sera feita o set update
		err = Update(tpep_collection, KeyUp, SetFields)

		// testando se tudo
		// correu bem
		if err == nil {

			//sucesso
			return

		} else {
			msgerror = "[UpCurriculum] " + err.Error()
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		}
	} else {
		msgerror = "[UpCurriculum] Uuid é obrigatório!"
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}
}

// Buscando Curriculum especifico
func GetCurriculum(Uuid string) (strJson string, err error) {

	// para atualizacao temos o nome do collection a chave para efetuar o update e
	// os campose que sera feita o set update
	strJson, err = Find(tpep_collection, Uuid)

	if err != nil {

		msgerror = "[UpCurriculum] Erro ao efetuar a busca: " + err.Error()
		log.Println(msgerror)
		return
	}

	return
}

// Buscando Curriculum especifico
func GetAllCurriculum() (strJson string, err error) {

	// para atualizacao temos o nome do collection a chave para efetuar o update e
	// os campose que sera feita o set update
	strJson, err = FindAll(tpep_collection)

	if err != nil {

		msgerror = "[UpCurriculum] Erro ao efetuar a busca: " + err.Error()
		log.Println(msgerror)
		return
	}

	return
}
