package mysql

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
	db.Model(&models.User{}).Preload("Role").Find(&users)
	return users
}

func (this UserRepositoryMysql) Update(user models.User) bool {
	//ESTE METODO ES PROPIO DE LA CLASE PERO NO SE IMPLEMENTA POR QUE NO ES NECESARIO PARA LA PRUEBA...
	return true
}

func (this UserRepositoryMysql) Delete(user models.User) bool {
	var db = this.ConnectGORM()
	result := db.Delete(&user)
	if result.Error == nil {
		return true
	} else {
		return false
	}

}

func (this UserRepositoryMysql) FindByIdUserUseCase(user models.User) models.User {
	var db = this.ConnectGORM()
	var userTemp models.User
	db.Model(&models.User{}).Preload("Role").Find(&userTemp, user.ID)
	return userTemp
}

func (this UserRepositoryMysql) CreateSeedsUsers() bool {
	user1 := *new(models.User)
	user1.FirstName = "Luis"
	user1.LastName = "Ospina"
	user1.NumberPhone = "3178976527"
	user1.Email = "lc@gmail.com"
	user1.CreatedAt = time.Time{}
	role1 := *new(models.Role)
	role1.ID = 1
	role1.Description = "Administrador"
	user1.Role = role1

	user2 := *new(models.User)
	user2.FirstName = "Pedro"
	user2.LastName = "Perez"
	user2.NumberPhone = "3208765467"
	user2.Email = "pp@gmail.com"
	user2.CreatedAt = time.Time{}
	role2 := *new(models.Role)
	role2.ID = 2
	role2.Description = "Cliente"
	user2.Role = role2

	var resp1 = this.Insert(user1)
	var resp2 = this.Insert(user2)

	if resp1 && resp2 {
		return true
	} else {
		return false
	}
}
