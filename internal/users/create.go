package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func insert(db *sql.DB, u *User) (id int64, err error) {

	stmt := `INSERT INTO users ("name", "login", "password", "modified_at")
			VALUES ($1, $2, $3, $4) RETURNING id`

	err = db.QueryRow(stmt, u.Name, u.Login, u.Password, u.ModifiedAt).
		Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil

}

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {

	u := new(User)

	// Request.Body is an io.ReadCloser which has Read(). For example, a *strings.Read() (for
	// user data such as plain text), *os.File.Read() (for uploaded files), and so on
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)

		return
	}

	err = u.SetHashedOriginalPassword(u.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)

		return
	}

	err = u.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)

		return
	}

	id, err := insert(h.db, u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)

		return
	}
	u.ID = id

	rw.Header().Set("Content-Type", "application/json")

	response := UserResponse{
		Name:  u.Name,
		Login: u.Login,
	}
	// NewEncoder wires the Encoder output to rw so Encoder knows where to send, using
	// ResponseWriter.Write(), the struct fields' content converted and formatted into JSON bytes
	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)

		return
	}

}
