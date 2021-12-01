# JWT Check - Mock para tests 

This repo is just a simple example of JWT token generation using RS256 algorithm and HS256 algorithm. 
This api is responsible for generating the token and using it to validate it.

## Install Local gjwtcheck with Docker
You can install gjwtcheck with docker.
```bash
$  docker run --rm --name gjwtcheck -it \
-p 8080:8080 jeffotoni/gjwtcheck:latest
```
## Install Local gjwtcheck
You can clone it and run it on your local machine. Remembering that you have to have Go installed.
```bash
$ git clone github.com/jeffotoni/gjwtcheck.git
$ cd gjwtcheck
$ make build 
$ ./gjwtcheck
```
## Ping
```bash
$ curl -i -XGET -H "Content-type:application/json" 
localhost:8080/ping
```

**Out:**
```bash
{
    "pong": "🏓"
}
```

## Token
```bash
$ curl -i -XPOST -H "Content-type:application/json" \
localhost:8080/token
```

```bash
$ curl -i -XGET -H "Content-type:application/json" \
localhost:8080/token
```
**Out:**
```bash
{
   "user":"TestUser",
   "token":"<token-jwt-rs256-here>",
   "public":"<public-rsa-pem>",
   "expires":"2021-12-01 05:12:05"
}
```

## hs256
```bash
$ curl -i -XPOST -H "Content-type:application/json" \
localhost:8080/hs256
-d '
{
    "user": "<your@email.com>"
}
'
```

**Out:**
```bash
{
    "user": "<your@email.com>",
    "token": "<token-jwt>",
    "secret": "<secret-hs256-here>",
    "expires": "<2021-05-26 10:26:36>"
}
```

## hs256/user
```bash
$ curl -i -XPOST -H "Content-type:application/json" \
-H "Authorization: Bearer $token" 
localhost:8080/hs256/user \
-d '
{
    "user": "<your@email.com>",
    "password": "<yourpassword>"
}
'
```

**Out:**
```bash
{
   "name":"HS256",
   "user":"your-user",
   "id":"447b22be-2d17-4253-9f4a-77a8501ef168",
   "iss":"gjwtcheck - created in:2021-12-01 00:01:42 
   expires:2021-12-01 00:05:42",
   "avatar":"https://logodix.com/logo/1989600.png",
   "message":"Welcome JWT HS256"
}
```

## rs256
```bash
$ curl -i -XPOST -H "Content-type:application/json" \
localhost:8080/rs256
-d '
{
    "user": "<your@email.com>"
}
'
```

**Out:**
```bash
{
    "user": "<your@email.com>",
    "token": "<token-jwt>",
    "public": "<rsa-public-here>"
    "expires": "<2021-05-26 10:26:36>"
}
```

## rs256/user
```bash
$ curl -i -XPOST -H "Content-type:application/json" \
-H "Authorization: Bearer $token" 
localhost:8080/rs256/user \
-d '
{
    "user": "<your@email.com>",
    "password": "<yourpassword>"
}
'
```

**Out:**
```bash
{
   "name":"RS256",
   "user":"your-user",
   "id":"447b22be-2d17-4253-9f4a-77a8501ef168",
   "iss":"gjwtcheck - created in:2021-12-01 00:01:42 
   expires:2021-12-01 00:05:42",
   "avatar":"https://logodix.com/logo/1989600.png",
   "message":"Welcome JWT RS256"
}
```
