createnetwork:
	docker network create -d bridge dev-network

updb:
	docker run --rm -d \
		--name db-so-cheap \
		--network dev-network \
		-e POSTGRES_USER=admin \
		-e POSTGRES_PASSWORD=admin \
		-e POSTGRES_DB=so-cheap \
		-v /Users/marcon/desenvolvimento/go/databases/so-cheap/data:/var/lib/postgresql \
		-p 5432:5432 \
		postgres:15.1

migrateup:
	docker run -v $(PWD)/pkg/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database postgres://admin:admin@localhost:5432/so-cheap?sslmode=disable up

runapp:
	docker-compose up -d --build

stopapp:
	docker-compose down

migratedown:
	docker run -v $(PWD)/pkg/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database postgres://admin:admin@localhost:5432/so-cheap?sslmode=disable down 4

stopdb:
	docker stop db-so-cheap

removenetwork:
	docker network rm dev-network