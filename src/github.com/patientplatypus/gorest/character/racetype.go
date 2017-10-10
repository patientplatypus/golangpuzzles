package character

import (
	"fmt"
	"net/http"
)

func RaceType(w http.ResponseWriter, r *http.Request) {
	race := r.FormValue("race")
	fmt.Fprintln(w, "Your Race: ", race)
	// ...
}
