echo "\nWithout threads"
go run main.go normal
echo "\nWith threads and normal lock"
go run main.go lock
echo "\nWith threads and skip lock"
go run main.go skip-lock