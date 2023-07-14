package main
import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	cst_name   string
	cst_email string
	cst_dob string
	kewarganegaraan string
	telepon string
}

func main() {
	// Read environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	// Create connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase)

	// Open database connection
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Read all users
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	// Display the result
	fmt.Println("Daftar Pengguna:")
	for _, user := range users {
		fmt.Println(user)
	}
}

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.cst_name, &user.cst_email,&user.kewarganegaraan,&user.telepon,&user.cst_dob)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
