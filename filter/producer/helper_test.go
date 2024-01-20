package producer_test

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koki120/table-spec-gen/config"
)

var (
	informationSchemaDB *sql.DB
)

func init() {
	db, err := sql.Open("mysql", config.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createUsersTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE
	);
	`

	if _, err = db.Exec(createUsersTableSQL); err != nil {
		log.Fatal(err)
	}

	createOrdersTableSQL := `
	CREATE TABLE IF NOT EXISTS orders (
		id INT AUTO_INCREMENT PRIMARY KEY,
		product_name VARCHAR(255) NOT NULL,
		user_id INT,
		quantity INT DEFAULT 1,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
	if _, err = db.Exec(createOrdersTableSQL); err != nil {
		log.Fatal(err)
	}

	informationSchemaDB, err = sql.Open("mysql", config.INFORMATION_SCHEMA_DSN())
	if err != nil {
		log.Fatal(err)
	}
}
