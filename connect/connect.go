package connect 

import(
	"log"
	"https://github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
	"../structures"
)

var connection *gorm.DB

// var for connection a serve database
// variables de entorno
const engine_sql string = "mysql"

const username string = "root"

const password string = ""

const database string = "course"


// "username:password@database"

func InitializeDatabase() {
	connection = ConnectionORM(CreateString())
	log.Println("La Conexion a la base de datos fue exitosa")
}

func CloseConnection() {
	connection.Close()
	log.Println("La Conexion a la base de datos fue cerrada")
}

func ConnectionORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open( engine_sql, stringConnection)
	if err != nil {
		log.Fatal(err)
		return nill
	}
	return connection
}

func GetUSer(id string) structures.User{
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	return user
}

func CreateString() string {
	return username + ":" + password + "@/" + database
}

func CreateUser(user structures.User) structures.User{
	connection.Create(&user) // se asigna un id
	return user
}

func UpdateUSer(id string, user structures.User) {
	currentUser := structures.User{}
	connection.Where("id = ?", id).First(&currentUser)

	currentUser.Username = user.Username
	currentUser.First_name = user.First_name
	currentUser.Last_name = user.First_name

	connection.Save(&currentUser)
	return currentUser
}

func DeleteUSer(id string) {
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	connection.Delete(&user)
}