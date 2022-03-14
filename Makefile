.PHONY: local
local:
	docker compose up --build

.PHONY: compose-down
compose-down:
	docker compose down --rmi local

.PHONY: test
test:
	go test ./... -cover
