package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func selectAll(db *sql.DB) ([]User, error) {

	stmt := `SELECT * FROM users WHERE deleted = false`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	us := make([]User, 0)
	for rows.Next() {

		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Login, &u.Password,
			&u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)
		if err != nil {
			return nil, err
		}
		us = append(us, u)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return us, nil
}

func (h *handler) List(rw http.ResponseWriter, r *http.Request) {

	us, err := selectAll(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)

		return
	}

	rw.Header().Set("Content-Type", "application/json")

	response := make([]UserResponse, 0)
	for _, u := range us {

		response = append(response, UserResponse{
			Name:  u.Name,
			Login: u.Login,
		})

	}

	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)

		return
	}
}
