module github.com/rocket-crm/order

go 1.24.6

replace github.com/rocket-crm/inventory => ../inventory

replace github.com/rocker-crm/shared/ => ../shared

require (
	github.com/go-chi/chi/v5 v5.2.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
