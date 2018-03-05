/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package model

// struct TPesqPerguntas
type TPesqCurriculum struct {
	Uuid       string `json:"uid"`
	Nome       string `json:"nome"`
	Cpf        string `json:"cpf"`
	Rg         string `json:"rg"`
	Idade      string `json:"idade"`
	Skill      string `json:"skill"`
	Bio        string `json:"bio"`
	Datetime   string `json:"datetime"`
	Dtcadastro string `json:"dtcadastro"`
	Dtaltera   string `json:"dtaltera"`
}
