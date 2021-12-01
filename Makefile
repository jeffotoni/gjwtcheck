# Makefile
.EXPORT_ALL_VARIABLES:	

GOPRIVATE=github.com/jeffotoni/gjwtcheck

build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gjwtcheck main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/gjwtcheck
	go mod tidy -compat=1.17 -go=1.17
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gjwtcheck main.go
	@echo "buid update completo..."
	@echo "fim"


tests:
	go test github.com/jeffotoni/gjwtcheck/gjwtcheck/controller/handler -v
	
deploy.aws:
	@make build
	@echo "########## Compilando nossa API ... "
	sh deploy-aws.sh
	@echo "fim"
