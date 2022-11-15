package repository

// Adapter (Repository)
import (
	"database/sql"
	"fmt"
)

// Private
type userRepositoryDB struct {
	db *sql.DB
}

func NewRepositoryDB(db *sql.DB) userRepositoryDB { // รับ instance database ที่เราจะใช้มาแล้วออกมาเป็น "Adapter" ของ repository
	return userRepositoryDB{db: db}
}

func (u userRepositoryDB) CreateUser(email string, password string, secret string) (*User, error) {
	// implement me
	//fmt.Print("User Inserted")
	//return nil, nil

	// Insert document into database
	insert, err := u.db.Exec("INSERT INTO users (email, password, secret) VALUES (?, ?, ?)", email, password, secret)
	if err != nil {
		return nil, err
	}
	userId, err := insert.LastInsertId()

	// Create user object
	var user = User{
		Id:       userId,
		Email:    email,
		Password: password,
		Secret:   secret,
	}
	return &user, nil
}
func (u userRepositoryDB) CheckUser(email string) (*User, error) {
	// implement me
	fmt.Print("User Fetched")
	return nil, nil
}

func (u userRepositoryDB) GetUsers() ([]*User, error) {
	// implement me
	fmt.Print("Users Fetched")
	return nil, nil
}
