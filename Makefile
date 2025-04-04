compose-up:
	docker compose up

compose-down:
	docker compose down
	docker image rm sas-golang-template-frontend -f 
	docker image rm sas-golang-template-backend -f
