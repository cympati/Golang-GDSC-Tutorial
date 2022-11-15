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
	fmt.Print("User Inserted")
	return nil, nil
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
