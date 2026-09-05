package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func insert(db *sql.DB, f *Folder) (id int64, err error) {
	stmt := `INSERT INTO "folders" ("name", "parent_id", "modified_at")
				VALUES ($1, $2, $3) RETURNING id`

	err = db.QueryRow(stmt, f.Name, f.ParentID, f.ModifiedAt).
		Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {

	f := new(Folder)
	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)

		return
	}

	err = f.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)

		return
	}

	id, err := insert(h.db, f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)

		return
	}

	f.ID = id

	rw.Header().Set("Content-Type", "application-json")
	rw.WriteHeader(http.StatusCreated)

	response := FolderResponse{f.Name}
	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}
