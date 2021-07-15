# same-calendar-api

REST API server that calculates years that have the same calendar.

## Installation:
```shell
go get -u github.com/lakhanmankani/same-calendar-api
```

## API Usage:
### Register
Request:
```shell
curl -X POST localhost:8080/api/register
```
Response:
```json
{
  "key": "9842824a468b7e0a1a01c9d802eb164c568ce3d86c1b7cfa0b80a7eab0379f3e"
}
```

### Get same calendar years in the future
Request:
```shell
curl -X GET localhost:8080/api/same-calendar?key=<apikey>&year=2020&n=5&forward=true
```
Response:
```json
{
  "years": [2020,2048,2076,2116,2144]
}
```

### Get same calendar years in the past
Request:
```shell
curl -X GET localhost:8080/api/same-calendar?key=<apikey>&year=2020&n=5&forward=false
```
Response:
```json
{
  "years": [2020,1992,1964,1936,1908]
}
```

### Unregister
Request:
```shell
curl -X DELETE localhost:8080/api/unregister?key=<apikey>
```
