# MYSQL Tables Specification Generator

This is a command-line application written in Go that connects to a MySQL database, extracts table information, and generates a file documenting the database tables.

## Usage:
```
table-spec-gen [flags]
```

### Flags:
-  -n, --dbname string     db name
-  -o, --filename string   output file name (default "output")
-  -f, --format string     output file format. Choose either md, html or stdout. (default "md")
-  -h, --help              help for open-mysql
-  -s, --host string       db host
-  -p, --password string   db password
-  -r, --port string       db port
-  -u, --user string       db user