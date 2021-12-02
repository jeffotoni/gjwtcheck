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
## Decode Token
Our JWT token will have the following format.

HEADER:ALGORITHM & TOKEN TYPE

```bash
{
"alg": "RS256",
"typ": "JWT"
}
```

PAYLOAD:DATA
```bash
{
  "user": "jeff@gmail.com",
  "id": "81e129ad-9289-4693-962a-ab4b17592ea7",
  "exp": 1638417515,
  "iat": 1638413915,
  "iss": "gjwtcheck - created in:2021-12-02 02:58:35 expires:2021-12-02 03:58:35",
  "nbf": 1638410315
}
```
VERIFY SIGNATURE
To validate the Token, just have the public or private key, the server will always return the public key so you can use it in your gateway, or in your services so that it can be validated.

So that you can view it even better, just enter: [jwt.io](https://jwt.io/)

![jwt.io](img/jwt-1.png?raw=true "jwt")

## Token
Token is a resource that will work by GET and POST method, it returns by default the RS256 algorithm a JWT token
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

This patch is responsible for generating an HS256 token and you can spend the time you would like it to expire with the "time" field.
It will return the jwt Hs256 token, and the secret used to generate the token and expiration date.

```bash
$ curl -i -XPOST -H "Content-type:application/json" \
localhost:8080/hs256
-d '
{
    "user": "<your@email.com>",
    "time": 3600
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

This patch is responsible for receiving your Authorization: Bearer <token> and if it works it means that everything went well in the validation, if you try to use another token or it expires the server will not accept and will return an error.
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
This patch is responsible for generating an RS256 token and you can spend the time you would like it to expire with the "time" field.
It will return the jwt rs256 token, the PEM public that was used to generate the expiry date token.

```bash
$ curl -i -XPOST -H "Content-type:application/json" \
localhost:8080/rs256
-d '
{
    "user": "<your@email.com>",
    "time": 3600
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
This patch is responsible for receiving your Authorization: Bearer <token> and if it works it means that everything went well in the validation, if you try to use another token or it expires the server will not accept and will return an error.
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
