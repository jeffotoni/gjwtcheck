# Makefile
.EXPORT_ALL_VARIABLES:

GOPRIVATE=github.com/jeffotoni/gjwtcheck

build:
	@echo "########## build API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gjwtcheck main.go
	@echo "buid success..."

update:
	@echo "########## build update API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/gjwtcheck
	go mod tidy -compat=1.17 -go=1.17
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gjwtcheck main.go
	@echo "buid update success..."
	@echo "the end"