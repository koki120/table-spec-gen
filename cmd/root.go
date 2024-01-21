package cmd

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koki120/table-spec-gen/filter/consumer"
	"github.com/koki120/table-spec-gen/filter/producer"
	"github.com/koki120/table-spec-gen/filter/transformer"
	"github.com/spf13/cobra"
)

const (
	HTML     = "html"
	MARKDOWN = "md"
	STDOUT   = "stdout"
)

var (
	outputFileFormat string
	outputFileName   string
	dbUser           string
	dbPassword       string
	dbHost           string
	dbPort           string
	dbName           string
)

var rootCmd = &cobra.Command{
	Use:   "open-mysql",
	Short: "Generates documentation for the MySQL tables.",
	Long:  "This is a command-line application written in Go that connects to a MySQL database, extracts table information, and generates a file documenting the database tables.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("mysql", INFORMATION_SCHEMA_DSN())
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		source, err := producer.FetchColumnMetadata(db, DBName())
		if err != nil {
			log.Fatal(err)
		}

		tables := transformer.ConvertColumnMetadataToTableMetaData(source)

		file, err := os.Create(OutputFileName())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		switch OutputFileFormat() {
		case HTML:
			err = consumer.ExportToHTML(file, tables)
		case MARKDOWN:
			err = consumer.ExportToMarkdown(file, tables)
		}
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&outputFileFormat, "format", "f", MARKDOWN, "output file format. Choose either md, html or stdout.")
	if outputFileFormat != HTML && outputFileFormat != MARKDOWN {
		log.Fatalf("%s is unavailable", outputFileFormat)
	}
	rootCmd.Flags().StringVarP(&outputFileName, "filename", "o", "output", "output file name")

	rootCmd.Flags().StringVarP(&dbUser, "user", "u", "", "db user")
	rootCmd.MarkFlagRequired("user")

	rootCmd.Flags().StringVarP(&dbPassword, "password", "p", "", "db password")
	rootCmd.MarkFlagRequired("password")

	rootCmd.Flags().StringVarP(&dbHost, "host", "s", "", "db host")
	rootCmd.MarkFlagRequired("host")

	rootCmd.Flags().StringVarP(&dbPort, "port", "r", "", "db port")
	rootCmd.MarkFlagRequired("port")

	rootCmd.Flags().StringVarP(&dbName, "dbname", "n", "", "db name")
	rootCmd.MarkFlagRequired("dbname")

}
