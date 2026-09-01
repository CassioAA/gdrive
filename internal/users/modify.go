package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func update(db *sql.DB, id int64, u *User) error {

	u.ModifiedAt = time.Now()

	stmt := `UPDATE "users" SET "name" = $1, "login" = $2, "modified_at" = $3 WHERE "id" = $4`
	_, err := db.Exec(stmt, u.Name, u.Login, u.ModifiedAt, id)

	return err

}

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	u := new(User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)

		return
	}

	if u.Name == "" {
		http.Error(rw, ErrNameRequired.Error(), http.StatusBadRequest)
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	err = update(h.db, int64(id), u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")

	response := UserResponse{
		Name:  u.Name,
		Login: u.Login,
	}

	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
