# Prerequisites

Only pre-req for building the application in local environment is installtion of gcc. Application is using sqlite3 for persistence which uses gcc for build.

# Configuration

For configuration of list of banned countries, use env file with comma separated 2 digit alphanumeric code of the country. Find list of code for all countries here. https://github.com/srcagency/iso-3166-1-codes/blob/master/data.json

# Information

This application exposes a http server with routes to add, list, get cards.
Cards are added to the system in pending state.
Cron jobs runs every 5 minute and process 5 cards for country validation.
Country validation is done through https://lookup.binlist.net/ API.
If country is banned, card is removed from the database.

# RUN

`$ go run main.go`

