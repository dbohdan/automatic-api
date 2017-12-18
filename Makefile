SRC = $(wildcard data/*.yml)

README.md: README.md.template render-template.go $(SRC)
	go run render-template.go > README.md
