setup:
	docker-compose --env-file ./dev.env up -d

tear-down:
	docker-compose --env-file ./dev.env down
	docker-compose --env-file ./dev.env stop
