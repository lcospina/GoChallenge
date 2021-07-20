package repositoriesimps

import (
	"GoChallenge/src/domains/models"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryMysql struct {
}

func (this UserRepositoryMysql) ConnectGORM() *gorm.DB {
	dsn := DSNMYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (this UserRepositoryMysql) Insert(user models.User) bool {
	var db = this.ConnectGORM()
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})
	result := db.Create(&user)
	if result.Error == nil {
		return true
	} else {
		return false
	}
}

func (this UserRepositoryMysql) GetAll() []models.User {
	var db = this.ConnectGORM()
	var users []models.User
	db.Find(&users)
	return users
}

func (this UserRepositoryMysql) Update(user models.User) bool {
	return true
}

func (this UserRepositoryMysql) Delete(user models.User) bool {
	return true
}

func (this UserRepositoryMysql) SelectFindId(user models.User) models.User {
	return user
}

func (this UserRepositoryMysql) CreateSeedsUsers() {
	user1 := *new(models.User)
	user1.FirstName = "Luis"
	user1.LastName = "Ospina"
	user1.NumberPhone = "3173040456"
	user1.Email = "lcospina9@gmail.com"
	user1.CreatedAt = time.Time{}
	role1 := *new(models.Role)
	role1.ID = 1
	role1.Description = "Administrador"
	user1.Role = role1

	user2 := *new(models.User)
	user2.FirstName = "Pedro"
	user2.LastName = "Perez"
	user2.NumberPhone = "3207767839"
	user2.Email = "pp@gmail.com"
	user2.CreatedAt = time.Time{}
	role2 := *new(models.Role)
	role2.ID = 2
	role2.Description = "Cliente"
	user2.Role = role2

	this.Insert(user1)
	this.Insert(user2)
}
