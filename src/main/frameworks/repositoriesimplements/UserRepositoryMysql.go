package repositoriesimplements

import (
	"GoChallenge/src/domains/models"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
