# JWT Authentication in Go Lang

## Packages used are :
### `gorm.io/gorm` <br>
GORM helps Go programs to work with data stored in database using the language's built in data types, rather than having to use SQL queries
### `grom.io/driver/postgres` <br>
grom.io/driver/postgres helps our webserver to connect with PostgreSQL
###  `gin-gonic/gin` <br>
Gin is a powerful web framework which allows us to create web server, handle routes, etc.
###  `crypto/bcrypt` <br>
bcrypt package provides secure way to hash and store passwords.
###  `golang-jwt/jwt` <br>
JWT helps us to use JSON web tokens instead of cookies
###  `godotenv` <br> 
 The godotenv package is a Go package that provides a way to load environment variables from a file into the environment of  Go application
 ###  `CompileDaemon` <br>
 Tracks .go files in a directory and invokes `go build` if a file is changed

 ## Endpoints

 ### POST : User Signup <br>
 `http://localhost:3000/signup`

 To create a new user.

Accepts username and password from the user. Unique email id contraint is added with GORM

Hashes the password using crypto/bcrypt and stores it in the database(postgres)

Body : raw json


`{
  "email": "test4",
  "password": "test4"
}`

### POST : User Login <br>
 `http://localhost:3000/login`

 Login receives email and password from the user and verfies it.

If the email id and password matches with the hased password stored in the database, it will create a joken using JWT and sends it as a cookie in response.

Body : raw json


`{
  "email": "test4",
  "password": "test4"
}`


