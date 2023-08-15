.PHONY : test run cov cov-report help 

help:
	@echo "test        - run unit tests."
	@echo "cov         - run test coverage."
	@echo "cov-report  - run test coverage with HTML report."
	@echo "run         - run the main package."

test:
	go test ./... -v

cov:
	go test -cover ./... 

cov-report:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

run:
	CSN_PAGES="~/Projects/cheatsheets-navigator/test/pages" go run main.go $(filter-out $@,$(MAKECMDGOALS))
