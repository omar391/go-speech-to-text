# Speech to Text service using GO-Lang (db: `SQLite3`)
⭐ For API we are using `GIN` http server.
<br>⭐ We have demonstrated how to use `security related headers` for the endpoints.
<br>⭐ Set-up `docker-compose` for easier deployments.
<br>⭐ Set-up `GitHub actions` for CI purposes.
<br>⭐ Included `TEST` suites
<br>⭐ Included `POSTMAN` collections for easier testing.

# NOTE: This project is still in active development. Many features are still incomlete and buggy.

<br>

### `POSTMAN` API testing:


<br>

## How to config
- HTTP configurations (port, dabase-path) are in the toml file: `./conf.d/app.toml`

## API Security
We are taking following measures to ensure API security.
- Setting up CORS security headers.
- Setting up the 'attack-vector' headers

### API security
```
Token based JWT security check: Coming soon

```

### Web gneral security
```
	//setting-up cors headers
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//setting-up security headers
	r.Use(secure.New(secure.Config{
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		//SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))
```

## How to run

```
go build -o stt-service
./stt-service
```
## How to run test

```
go clean -testcache  
go test ./...
```

## API endpoint: 

- Use the POSTMAN collection JSON config and test there or you can use CURL.
- Login and audio text searching is available.

### Login and get an access token
```
curl http://127.0.0.1:5000/login?user=a&pass=a

TODO: convert it to POST method
```

### Get all audio transcribed JSON  data
```
curl http://127.0.0.1:5000/all-data?token=<token>
```

### Search AUDIO text with criteria: name
```
curl http://127.0.0.1:5000/search?token=<token>&text=beautiful

```


## How to build docker images

```
docker-compose up
```

## How to push to Heroku

```
coming soon
```

## References

```
coming soon
```