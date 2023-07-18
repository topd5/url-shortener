# URL shortener
Run:
```bash
docker compose up
```

POST request to:
```
http://localhost:11200/api/short
```
with payload:
```json
{
	"url": "https://example.com"
}
```
201 answer example:
```json
{
	"url": "https://example.com",
	"short": "Q4KCcAP"
}
```

To get loopback url from shorten value:
GET request to:
```
http://localhost:11200/api/Q4KCcAP
```
200 answer example:
```json
{
	"short": "Q4KCcAP",
	"url": "https://example.com"
}
```

All data is stored in Redis.
