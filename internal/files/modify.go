package files

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func update(db *sql.DB, id int64, f *File) error {
	stmt := `UPDATE "folders" SET "name" = $1, "modified_at" = &2 WHERE "id" = $3`
	_, err := db.Exec(stmt, f.Name, f.ModifiedAt, id)

	return err
}

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	f := new(File)
	err = json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	err = update(h.db, int64(id), f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusNoContent)
}
