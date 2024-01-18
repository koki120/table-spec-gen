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
	dbName = os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME is empty")
	}
	output_file_format = os.Getenv("OUTPUT_FILE_FORMAT")
	if output_file_format == "" || (output_file_format != HTML && output_file_format == MARKDOWN) {
		output_file_format = MARKDOWN
	}
	output_file_name = os.Getenv("OUTPUT_FILE_NAME")
	if output_file_name == "" {
		output_file_name = "output"
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
