package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)


var DB *sql.DB

func ConnectToDB() {
	workingDir, err := os.Getwd();
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Working directory: ", workingDir)

	db, err := sql.Open("sqlite", workingDir + "/database.db")

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if err != nil {
		log.Fatal(err)
	}
}
