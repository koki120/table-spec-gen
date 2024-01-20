# Table Specification Generator

This is a command-line application written in Go that connects to a MySQL database, extracts table information, and generates a file documenting the database schema.

## Usage 
To run the application, you need to set the following environment variables:

- DB_USER: Database user
- DB_PASSWORD: Database password
- DB_HOST: Database host
- DB_PORT: Database port
- DB_NAME: Database name

Once the environment variables are set, execute the following command:

```bash
go run main.go -outputfilename=output -outputformat=md
```
- outputfilename: Specify the name of the output file. In the example above, it's set to "output," but you can choose a different name. The default value is "table_spec"

- outputformat: Specify the desired output format. Currently, the supported formats are Markdown (md) and HTML. Choose either md or html. The default value is "md"
