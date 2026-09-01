package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func selectByID(db *sql.DB, id int64) (*User, error) {
	stmt := `SELECT * FROM "users" WHERE id=$1`
	var u User
	err := db.QueryRow(stmt, id).
		Scan(&u.ID, &u.Name, &u.Login, &u.Password,
			&u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ResponseWriter is an interface, and interfaces are
// already represented by two pointers, so passing as pointer
// would just mainly increase GC workload. Also, the handler 
// function signature must be equivalent to Handler.ServeHTTP
// for setting Chi routers.
func (h *handler) GetByID(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		
		return
	}

	u, err := selectByID(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	
	response := UserResponse {
		Name:  u.Name,
		Login: u.Login,
	}
	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		
		return
	}
}
