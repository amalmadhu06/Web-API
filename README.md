# JWT Authentication in Go Lang

Packages used are :
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
