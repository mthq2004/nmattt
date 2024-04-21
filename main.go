package main

import (
	"fmt"
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/routes"
)

func main() {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	// if err != nil {
	// 	log.Fatal("DB_PORT must be a valid integer")
	// }

	// sql := &database.Sql{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     port,
	// 	Password: os.Getenv("DB_PASS"),
	// 	UserName: os.Getenv("DB_USER"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// 	DbName:   os.Getenv("DB_NAME"),
	// }

	// db, err := sql.Connect()
	// if err != nil {
	// 	fmt.Println("Failed to connect to the database:", err)
	// 	return
	// }
	// defer sql.Close()

	// repo := &database.Repository{
	// 	DB: db,
	// }

	fmt.Println("Successfully")

	routes.AESRoutes()
	routes.RSARoutes()
	http.ListenAndServe(":8080", nil)

}
