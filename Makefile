makemigration:
	migrate create -ext sql -dir migrations -format 20060102150405 $(NAME)

migrate-up:
	migrate -source "file://migrations" -database "mysql://root@tcp(localhost:3306)/orm_test" up

migrate-down:
	migrate -source "file://migrations" -database "mysql://root@tcp(localhost:3306)/orm_test" down

generate-orm:
	sqlboiler --wipe --no-tests -o ./internal/pkg/comparison/sqlboilerrepo/models mysql