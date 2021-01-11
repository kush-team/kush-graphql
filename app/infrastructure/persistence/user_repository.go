package persistence

import (
	"kush-graphql/app/domain/repository/user"
	"kush-graphql/app/models"

	"github.com/jinzhu/gorm"
)

type userService struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *userService {
	return &userService{
		db,
	}
}

var _ user.UserService = &userService{}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {

	err := s.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {

	err := s.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *userService) DeleteUser(id string) error {

	usr := &models.User{}

	err := s.db.Where("id = ?", id).Delete(usr).Error
	if err != nil {
		return err
	}

	return nil

}

func (s *userService) GetUserByID(id string) (*models.User, error) {

	usr := &models.User{}

	err := s.db.Where("id = ?", id).Take(&usr).Error
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *userService) GetUserByEmail(emailAddress string) (*models.User, error) {

	usr := &models.User{}

	err := s.db.Where("email_address = ?", emailAddress).Take(&usr).Error
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *userService) GetAllUsers() ([]*models.User, error) {

	var users []*models.User

	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
