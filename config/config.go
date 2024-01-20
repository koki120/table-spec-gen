package config

import (
	"flag"
	"log"
	"os"
)

const (
	HTML     = "html"
	MARKDOWN = "md"
)

var (
	dbName           string
	outputFileFormat string
	outputFileName   string
)

func init() {
	flag.StringVar(&outputFileFormat, "outputformat", MARKDOWN, "Output file format (markdown or html)")
	flag.StringVar(&outputFileName, "outputfilename", "table_spec", "Output file name")
	flag.Parse()

	if outputFileFormat != HTML && outputFileFormat != MARKDOWN {
		log.Fatalf("%s is unavailable", outputFileFormat)
	}

	if os.Getenv("DB_USER") == "" {
		log.Fatal("DB_USER is empty")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		log.Fatal("DB_PASSWORD is empty")
	}

	if os.Getenv("DB_HOST") == "" {
		log.Fatal("DB_HOST is empty")
	}

	if os.Getenv("DB_PORT") == "" {
		log.Fatal("DB_PORT is empty")
	}

	dbName = os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME is empty")
	}

}

func OutputFileName() string {
	return outputFileName + "." + outputFileFormat
}

func OutputFileFormat() string {
	return outputFileFormat
}

func DBName() string {
	return dbName
}
