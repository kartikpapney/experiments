echo "Without threads"
go run main.go normal
echo "\n"
echo "With threads and normal lock"
go run main.go lock
echo "\n"
echo "Without threads and skip lock"
go run main.go skip-lock
echo "\n"