module github.com/rocker-crm/assembly

replace (
	github.com/rocker-crm/platform/ => ../platform
	github.com/rocker-crm/shared/ => ../shared
)

go 1.24.6

require github.com/caarlos0/env/v11 v11.3.1 // indirect
