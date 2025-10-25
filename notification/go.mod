module github.com/rocker-crm/notifacation

replace (
	github.com/rocker-crm/platform/ => ../platform
	github.com/rocker-crm/shared/ => ../shared
)

go 1.24.6

require github.com/go-telegram/bot v1.17.0 // indirect
