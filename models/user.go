package models

func LoginCheck(username string, passwd string) (bool, int) {
	var id int
	var salt int
	var password string
	DB.QueryRow("SELECT password FROM users WHERE email = ?", username).Scan(&password)
	DB.QueryRow("SELECT id FROM users WHERE email = ?", username).Scan(&id)
	DB.QueryRow("SELECT salt FROM users WHERE email = ?", username).Scan(&salt)

	if password == Compare(passwd) {
		return true, id
	}
	DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&password)
	DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	DB.QueryRow("SELECT salt FROM users WHERE username = ?", username).Scan(&id)
	if password == Compare(passwd) {
		return true, id
	}
	return false, 0
}
