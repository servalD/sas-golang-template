package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Import nécessaire pour PostgreSQL
)

func InitDatabase() (db *sql.DB, err error) {
	var dbURLvars []string
	for _, variable := range [5]string{"POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_DB"} {
		value, exists := os.LookupEnv(variable)
		if !exists {
			return nil, fmt.Errorf("%s environment variable not set", variable)
		}
		dbURLvars = append(dbURLvars, value)
	}
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbURLvars[0], dbURLvars[1], dbURLvars[2], dbURLvars[3], dbURLvars[4])
	fmt.Println("Connecting to database:", dbURL)
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	// Create all needed tables if not exists
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		role VARCHAR(20) NOT NULL DEFAULT 'user'
	);
	CREATE TABLE IF NOT EXISTS user_sessions (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		token VARCHAR(255) NOT NULL,
		expiry TIMESTAMP NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS nfts (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT,
		image_url VARCHAR(255) NOT NULL,
		owner_id INT NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS collections (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT,
		image_url VARCHAR(255) NOT NULL,
		owner_id INT NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS collection_nfts (
		collection_id INT NOT NULL,
		nft_id INT NOT NULL,
		PRIMARY KEY (collection_id, nft_id),
		FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
		FOREIGN KEY (nft_id) REFERENCES nfts(id) ON DELETE CASCADE
	);`)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}
	return db, nil
}



// func CreateUser(db *sql.DB, user User) error {
// 	_, err := db.Exec("INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)",
// 		user.Username, user.Email, user.Password, user.Role)
// 	return err
// }
func CreateUser(db *sql.DB, user *User) error {
	return db.QueryRow(
		"INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Username, user.Email, user.Password, user.Role,
	).Scan(&user.ID)
}

func UpdateUser(db *sql.DB, user User) error {
	_, err := db.Exec("UPDATE users SET username = $1, email = $2, password = $3, role = $4 WHERE id = $5",
		user.Username, user.Email, user.Password, user.Role, user.ID)
	return err
}
func DeleteUserByID(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
func GetUserByID(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email, password, role FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func GetUserByUsername(db *sql.DB, username string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email, password, role FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username, email, password, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUserSession(db *sql.DB, session UserSession) error {
	_, err := db.Exec("INSERT INTO user_sessions (user_id, token, expiry) VALUES ($1, $2, $3)",
		session.UserID, session.Token, session.Expiry)
	return err
}
func UpdateUserSession(db *sql.DB, session UserSession) error {
	_, err := db.Exec("UPDATE user_sessions SET token = $1, expiry = $2 WHERE id = $3",
		session.Token, session.Expiry, session.ID)
	return err
}
func DeleteUserSessionByID(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM user_sessions WHERE id = $1", id)
	return err
}
func GetUserSessionByID(db *sql.DB, id int) (UserSession, error) {
	var session UserSession
	err := db.QueryRow("SELECT id, user_id, token, expiry FROM user_sessions WHERE id = $1", id).Scan(&session.ID, &session.UserID, &session.Token, &session.Expiry)
	if err != nil {
		return UserSession{}, err
	}
	return session, nil
}
func GetUserSessionByToken(db *sql.DB, token string) (UserSession, error) {
	var session UserSession
	err := db.QueryRow("SELECT id, user_id, token, expiry FROM user_sessions WHERE token = $1", token).Scan(&session.ID, &session.UserID, &session.Token, &session.Expiry)
	if err != nil {
		return UserSession{}, err
	}
	return session, nil
}
