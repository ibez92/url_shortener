# URL SHORTENER
Простенький сокращатель ссылок

# Local start
```bash
go run cmd/url_shortener/main.go
```

# Request examples
Create short url
```bash
curl -v -X POST -d '{"url": "https://google.com"}' -H 'content-type: application/json' localhost:3000/api/v1/shorten
```

Get info by short url
```bash
curl -v localhost:3000/api/v1/shorten/{shortCode}
```
