# GoChallenge

# Paso 1 - Clonar el repositorio

        https://github.com/lcospina/GoChallenge.git

# Paso 2 - Instalar dependencias requeridas

        go get -u github.com/go-sql-driver/mysql

        go get -u gorm.io/gorm

        go get -u gorm.io/driver/mysql

        go get -u github.com/gin-gonic/gin

# Paso 3 - Ejecutar el main

Se debe ingresar a la carpeta raiz del proyecto clonado y ejecutar el siguiente comando
        
        go run main.go


# Paso 4 - End Point primer punto

POST
http://localhost:9090/Numbers/orderNumbers

        curl -X POST http://localhost:9090/Numbers/orderNumbers -H "Content-Type: application/json" -d "{ \"unsorted\": [3,3,3,3,3,5,5,6,8,3,20,4,4,7,7,1,1,2,20,20] }"



# Paso 5 - End Point segundo punto


GET
http://localhost:9090/Users/findById/:id
        
        curl -X GET http://localhost:9090/Users/findById/1 

GET
http://localhost:9090/Users/getAllUsers

        curl -X GET http://localhost:9090/Users/getAllUsers 
        
# Paso 6 - End Points adicionales al segundo punto

GET
http://localhost:9090/Users/createSeedsDB

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
