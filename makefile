dev:
	docker-compose -f compose.dev.yaml --env-file .env.dev up --build
prod:
	docker-compose -f compose.prod.yaml --env-file .env.prod up --build
