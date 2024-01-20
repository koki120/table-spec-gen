package config

import (
	"log"
	"os"
)

const (
	HTML     = "html"
	MARKDOWN = "md"
)

var (
	dbName             string
	output_file_format string
	output_file_name   string
)

func init() {
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

	output_file_format = os.Getenv("OUTPUT_FILE_FORMAT")
	if output_file_format == "" {
		output_file_format = MARKDOWN
	}
	if output_file_format != HTML && output_file_format != MARKDOWN {
		log.Fatalf("%s is unavailable", output_file_format)
	}

	output_file_name = os.Getenv("OUTPUT_FILE_NAME")
	if output_file_name == "" {
		output_file_name = ""
	}
}

func OutputFileName() string {
	return output_file_name + "." + output_file_format
}

func OutputFileFormat() string {
	return output_file_format
}

func DBName() string {
	return dbName
}
