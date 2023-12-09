run-db:
	docker-compose up --build mysql
db-migrate:
	migrate -path build/db/migration/ddl/ -database 'mysql://root:@tcp(localhost:3306)/dramas?parseTime=true&loc=Local' up
# TODO: 以下が動かない所から確認
run-server:
	docker-compose up --build server