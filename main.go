package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koki120/table-spec-gen/config"
	"github.com/koki120/table-spec-gen/filter/consumer"
	"github.com/koki120/table-spec-gen/filter/producer"
	"github.com/koki120/table-spec-gen/pipe"
)

func main() {
	db, err := sql.Open("mysql", config.DSN())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	source, err := producer.ProduceFromDB(db, config.DBName())
	if err != nil {
		log.Fatal(err)
	}

	columns, err := pipe.RowsToColumns(source)
	if err != nil {
		log.Fatal(err)
	}

	tables := pipe.ColumnsToTables(columns)

	file, err := os.Create(config.OutputFileName())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	switch config.OutputFileFormat() {
	case config.HTML:
		err = consumer.ExportToHTML(file, tables)
	case config.MARKDOWN:
		err = consumer.ExportToMarkdown(file, tables)
	}
	if err != nil {
		log.Fatal(err)
	}

}
