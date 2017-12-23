SRC = $(wildcard data/*.yml)


.PHONY: fmt


all: fmt README.md

fmt:
	go fmt render-template.go

README.md: README.md.template render-template.go $(SRC)
	go run render-template.go > README.md
