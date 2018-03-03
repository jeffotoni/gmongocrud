/*
* Go Library (C) 2017 Inc.
*
* @project    Project Globo / avaliacao.com
* @author      @jeffotoni
* @size        01/03/2018
 */

package model

// struct TPesqPerguntas
type TRespostas struct {
	Rsp_uuid       string `json:"rsp_uid"`
	Rsp_cod        int    `json:"rsp_cod"`
	Rsp_per_cod    int    `json:"rsp_per_cod"`
	Rsp_titulo     string `json:"rsp_titulo"`
	Rsp_correta    int    `json:"rsp_correta"`
	Rsp_dtcadastro string `json:"rsp_dtcadastro"`
	Rsp_datetime   string `json:"rsp_datetime"`
}
