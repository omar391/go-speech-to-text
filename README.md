

# Speech to Text service using GO-Lang (db: `SQLite3`)
⭐ For API we are using `GIN` http server.
<br>⭐ We have demonstrated how to use `security related headers` for the endpoints.
<br>⭐ Used FULL-TEXT search indexing using FTS4 module to filter audio data faster.
<br>⭐ APP ensure security with statefull HMAC SHA256 JWT tokens.
<br>⭐ Set-up `docker-compose` for easier deployments.
<br>⭐ Set-up `GitHub actions` for CI purposes.
<br>⭐ Included `TEST` suites
<br>⭐ Included `POSTMAN` collections for easier testing.

`DEMO app:` https://protected-mountain-13923.herokuapp.com/

<br>

### `APP web interface`:
![image](https://user-images.githubusercontent.com/13569609/114558100-27156480-9c8c-11eb-80e8-bf63f00abb07.png)
![image](https://user-images.githubusercontent.com/13569609/114558049-19f87580-9c8c-11eb-80c2-722e27c31a12.png)


<br>

### `POSTMAN` API testing:
![image](https://user-images.githubusercontent.com/13569609/114550302-d863cc80-9c83-11eb-846b-4d0d9cfb6fdf.png)
![image](https://user-images.githubusercontent.com/13569609/114550427-00533000-9c84-11eb-8af5-dde88fcb7e00.png)


<br>

## How to config
- HTTP configurations (port, database-path) are in the toml file: `./conf.d/app.toml`

## API Security
We are taking following measures to ensure API security.
- Setting up CORS security headers.
- Setting up the 'attack-vector' headers

### API security
```
Token based JWT security check: Coming soon

```

### Web general security
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
## Hot code reloading: 
```
nodemon -e js,go,json,html,css,toml --exec go run main.go --signal SIGTERM
```

## API endpoint: 

- Use the POSTMAN collection JSON config and test there or you can use CURL.
- Login and audio text searching is available.

### Login and get an access token
```
curl -X POST -d [options] http://127.0.0.1:5000/login
OPTIONS: user=a pass=a

```

### Register and get an access token
```
curl -X POST -d [options] http://127.0.0.1:5000/register
OPTIONS: user=a pass=a email=a@a.com

```

### Transcribe to get audio text data
```
curl -X POST -d [options] http://127.0.0.1:5000/transcribe
OPTIONS: token=<token> file=<file> is_save_file=<true/false>

```

### Get all audio transcribed JSON  data
```
curl -X POST -d [options] http://127.0.0.1:5000/all-data

OPTIONS: token=<token> page_no=1
```

### Filter AUDIO text with criteria: text
```
curl -X POST -d [options] http://127.0.0.1:5000/filter

OPTIONS: token=<token> page_no=1 query=<search_term>
```


## How to build docker images

```
docker-compose up
```

## How to push to Heroku

```
1. Install Docker
2. Install Heroku cli
3. Create a Heroku project
4. Run docker container locally. Then run these commands:

- heroku login
- heroku container:login
- heroku create
- heroku container:push web
- heroku container:release web
- heroku open

```

## TODO:

```
- Create and run unimplemented  test suites.
- Config GitHub actions to run properly.
```

## References

```
- Admin panel used in the UI: https://adminlte.io/
- GIN API documentations: https://github.com/gin-gonic/gin#readme
```
