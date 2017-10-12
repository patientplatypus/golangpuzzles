package dungeon

import (
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func DeleteAllDungeonTypes(w http.ResponseWriter, r *http.Request) {
	db := config.Sql_connect()
	log.Print("inside DeleteAllDungeonClasses() in dungeon")
	db.Query("TRUNCATE dungeon_classes")
	db.Query("TRUNCATE dungeon_races")
}
