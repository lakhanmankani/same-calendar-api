# same-calendar-api

REST API that gives you the years that have the same calendar.

## Installation:
```bash
go get -u github.com/lakhanmankani/same-calendar-api
```

## API Usage:
### Register
Request:
```bash
curl -X POST localhost:8080/api/register
```
Response:
```json
{
"key": "9842824a468b7e0a1a01c9d802eb164c568ce3d86c1b7cfa0b80a7eab0379f3e"
}
```

### Get Same Calendar Years
Request:
```bash
curl -X GET localhost:8080/api/same-calendar?key=<apikey>;&year=2020&n=5
```
Response:
```json
[2020, 2048, 2076, 2116, 2144]
```