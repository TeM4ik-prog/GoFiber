package services

import (
	"fiber-app/database"
	"fiber-app/database/models"
	"log"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (cs *UserService) GetUsers() []models.User {
	var users []models.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Printf("Error retrieving users: %v", result.Error)
		return []models.User{}
	}
	return users
}

func (cs *UserService) CreateUser(user models.User) models.User {
	// existingUser, err := cs.findUserById(int64(user.ID))
	// if err != nil {
    //     return existingUser
    // }

	
	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Fatalf("Failed to create user: %v", result.Error)
		return models.User{}
	}

	return user
	

}


func (cs *UserService) findUserById(id int64) (models.User, error){
	var user models.User
    result := database.DB.Where("id =?", id).First(&user)

    if result.Error!= nil {
        log.Printf("Error retrieving user with id %d: %v", id, result.Error)
        return models.User{}, result.Error
    }
    return user, nil
}
