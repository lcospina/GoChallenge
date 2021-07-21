# GoChallenge

Dependencias requeridas

go get -u github.com/go-sql-driver/mysql

go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

go get -u github.com/gin-gonic/gin

Servicios Disponibles

GET
http://localhost:9090/Users/createSeedsDB

GET
http://localhost:9090/Users/getAllUsers

GET
http://localhost:9090/Users/findById/:id

DELETE
http://localhost:9090/Users/deleteUser/:id

POST
http://localhost:9090/Users/createUser

curl -X POST http://localhost:9090/Users/createUser -H "Content-Type: application/json" --data-binary @- <<DATA
{
"FirstName": "Pedro",
"LastName": "Perez",
"Email": "pp@gmail.com",
"NumberPhone": "3207767839",
"RoleID": 2,
"Role": {
"ID": 2,
"Description": "Cliente",
"CreatedAt": "2021-07-20T18:43:46.193Z"
},
"CreatedAt": "2021-07-20T18:43:46.194Z"
}
DATA
