all: build

build: 
	@go build -o main main.go

docker-up:
	@if docker compose up 2>/dev/null; then \
			; : \
	else \
			echo "Falling back to docker compose v1"; \
			docker-compose up; \
	fi

docker-down
	@if docker compose down 2>/dev/null; then \
			; : \
	else \
			echo "Falling back to docker compose v1"; \
			docker-compose down; \
	fi

clean:
	@echo "cleaning..."
	@rm -f main

watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi