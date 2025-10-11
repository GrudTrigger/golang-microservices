module github.com/rocket-crm/inventory

replace (
	github.com/rocker-crm/shared/ => ../shared
	github.com/rocker-crm/platform/ => ../platform
)


go 1.24.6

require (
	github.com/brianvoe/gofakeit/v7 v7.7.3 // indirect
	github.com/caarlos0/env/v11 v11.3.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	go.mongodb.org/mongo-driver v1.17.4 // indirect
)
