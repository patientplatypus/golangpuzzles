package character

import (
	"fmt"
	"net/http"
)

func ClassType(w http.ResponseWriter, r *http.Request) {
	class := r.FormValue("class")
	fmt.Fprintln(w, "Your Class: ", class)
}
