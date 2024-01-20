package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koki120/table-spec-gen/config"
	"github.com/koki120/table-spec-gen/filter/consumer"
	"github.com/koki120/table-spec-gen/filter/producer"
	"github.com/koki120/table-spec-gen/filter/transformer"
)

func main() {
	db, err := sql.Open("mysql", config.INFORMATION_SCHEMA_DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	source, err := producer.FetchColumnMetadata(db, config.DBName())
	if err != nil {
		log.Fatal(err)
	}

	tables := transformer.ConvertColumnMetadataToTableMetaData(source)

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
