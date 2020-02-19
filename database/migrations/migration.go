package migrations

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/source/file"
	cmd "gitlab.com/scalent/goxpert/services/user/cmd/service"
)

//Run Function to Run Migrations
func init() {

	// configs := config.Load()
	//db, err := sql.Open(configs.DatabaseDialect, configs.DatabaseUsername+":"+configs.DatabasePassword+"@/"+configs.DatabaseName+"?multiStatements=true")

	db, err := sql.Open("mysql", cmd.CreateConnectionString())
	if err != nil {
		fmt.Println(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println(err)
	}

	//Path for Migration Files
	fsrc, err := (&file.File{}).Open("file://migrations/schema")

	if err != nil {
		fmt.Println(err)
	}

	m, err := migrate.NewWithInstance(
		"file",
		fsrc,
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println(err)
	}

	err = m.Up()

	if err != nil {
		log.Println("Migration Status:", err)
	}
	log.Println("Migration checked successfully!")
}
