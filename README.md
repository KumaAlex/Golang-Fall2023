# Golang-Fall2023

Commands:

1. Migrations
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up

2. CRUD operations
POST
BODY='{"brand":"Urban lifestyle","color":"black","weight":"1.6 kg", "Dimensions":[10,16,41.6]}'
curl -i -d "$BODY" localhost:4000/v1/laptopBags
BODY='{"brand":"Pandec","color":"black and white","weight":"4 kg", "Dimensions":[40, 40, 40]}'
curl -i -d "$BODY" localhost:4000/v1/laptopBags
BODY='{"brand":"Nig","color":"black","weight":"0.1 kg", "Dimensions":[4,5,6]}'
curl -i -d "$BODY" localhost:4000/v1/laptopBags

UPDATE
BODY='{"brand":"Urban lifestyle","color":"black","weight":"69 kg", "Dimensions":[10,16,41.6]}'
curl -X PUT -d "$BODY" localhost:4000/v1/laptopBags/1

GET
curl localhost:4000/v1/laptopBags/1

DELETE
curl -X DELETE localhost:4000/v1/laptopBags/3

