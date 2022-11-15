package service

import "cympati/Golang-GDSC-Tutorial/repository"

// Adapter (Service)
type userService struct {
	repository repository.UserRepository // interface หรือ "Port" ของ repository
}

func NewUserService(userRepository repository.UserRepository) userService { // รับ "Adapter" ของ Repository
	return userService{repository: userRepository}
}

func (s userService) SignUp(email string, password string) (*string, *string, error) {
	// implement me
	s.repository.CreateUser(email, password, "")
	return nil, nil, nil
}

func (s userService) SignIn(email string, password string) (*User, error) {
	// implement me
	s.repository.CheckUser(email)
	//user, err := s.repository.CheckUser(email)
	//if err != nil {
	//	return nil, err
	//}
	//payload := &User{
	//	Id:    user.Id,
	//	Email: user.Email,
	//}
	//return payload, nil
	return nil, nil
}

func (s userService) ListUsers() ([]*User, error) {
	users, err := s.repository.GetUsers()
	if err != nil {
		return nil, err
	}
	var userResponse []*User
	for _, user := range users {
		userResponse = append(userResponse, &User{Id: user.Id, Email: user.Email})
	}
	return userResponse, nil
}
