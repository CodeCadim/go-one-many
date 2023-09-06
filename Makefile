.SILENT:

.PHONY: reset-db
reset-db:
	cat init-db.sql | docker-compose exec -T db mysql demo

.PHONY: connect-db
connect-db:
	docker-compose exec db mysql demo
