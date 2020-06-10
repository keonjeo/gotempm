package statements

import (
	"fmt"
	"goTemp/globalUtils"
	"log"
)

var language = globalUtils.LangEN

func SetLanguage(newLanguage globalUtils.Languages) {
	language = newLanguage
}

type UserErr string

var errTxtEn = map[string]UserErr{
	"internalError":          "Internal error. Error: %v\n",
	"insertError":            "Unable to create user. Error: %v\n",
	"updateError":            "Unable to update user. Error: %v \n",
	"deleteError":            "Unable to delete user %v. Error: %v\n",
	"deleteRowNotFoundError": "row with id %d not found. Unable to delete the row",
	"selectReadError":        "Unable to get rows from the DB. Error: %v \n",
	"selectScanError":        "Unable to read the user rows returned from the Db. Error: %v\n",
	"selectRowReadError":     "Unable to get row from the DB. Error: %v \n",
}

var errTxtES = map[string]UserErr{
	"internalError":          "Error interno. Error: %v\n",
	"insertError":            "No se pudo crear la auditoria. Error: %v\n",
	"updateError":            "No se pudo actualizar la auditoria. Error: %v \n",
	"deleteError":            "No se pudo borrar la auditoria %v. Error: %v\n",
	"deleteRowNotFoundError": "La auditoria %d no se pudo encontrar. No se pudo borrar el usuario",
	"selectReadError":        "No su pudo leer datos de la base de datos. Error: %v \n",
	"selectScanError":        "No se pudo leer los datos recibidos de la base de datos. Error: %v\n",
	"selectRowReadError":     "No se pudo leer la auditoria de la base de datos. Error: %v \n",
}

func (ge *UserErr) getSqlTxt(errKey string, myLanguage globalUtils.Languages) string {
	var returnstr string
	switch myLanguage {
	case globalUtils.LangEN:
		returnstr = string(errTxtEn[errKey])
	case globalUtils.LangES:
		returnstr = string(errTxtES[errKey])
	case globalUtils.LangFR:
		log.Fatalf("%s language not implemented for users", myLanguage)
	default:
		log.Fatalf("%s language not implemented for users", myLanguage)
	}
	return returnstr
}

func (ge *UserErr) internalError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("internalError", language), err)
}

func (ge *UserErr) InsertError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("insertError", language), err)
}

func (ge *UserErr) UpdateError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("updateError", language), err)
}

func (ge *UserErr) DeleteError(Id int64, err error) string {
	return fmt.Sprintf(ge.getSqlTxt("deleteError", language), Id, err)
}

func (ge *UserErr) DeleteRowNotFoundError(id int64) string {
	return fmt.Sprintf(ge.getSqlTxt("selectRowReadError", language), id)
}

func (ge *UserErr) SelectReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectReadError", language), err)
}

func (ge *UserErr) SelectScanError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectScanError", language), err)
}

func (ge *UserErr) SelectRowReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectRowReadError", language), err)
}
