run:
	go run main.go
refresh_db:
	psql -U postgres -d postgres -f ./sql/init.sql
create_tables:
	psql -U postgres -d shop_db -f ./sql/create.sql
insert_data:
	psql -U postgres -d shop_db -f ./sql/insert.sql