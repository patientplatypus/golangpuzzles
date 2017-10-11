package users

import (
	"log"

	"github.com/patientplatypus/gorest/config"
)

func DeleteAll() {
	db := config.Sql_connect()
	log.Print("inside DeleteAll() in users")
	db.Query("TRUNCATE users")
}
