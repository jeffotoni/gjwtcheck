# JWT Check - Mock para tests 

Este repo é somente um exemplo simples de geração de token **JWT** usando **algorítimo RS256** e **algoritimo HS256**.
Esta api é responsável por gerar o token e usa-lo para valida-lo.

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

## Check
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
    "key": "<token-jwt>",
    "expires": "<2021-05-26 10:26:36>"
}

```

## User
```bash

$ curl -i -XPOST -H "Content-type:application/json" \
-H "Authorization: Bearer $token2" 
localhost:8080/auth/user \
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
    "user_name": "name",
    "user_email": "<your@email.com>",
    "avatar_url": "",
    "message": "Welcome",
    "token": "eyJhbGciOiJSU...",
    "expires": "2021-05-28 18:50:08"
}

```

## Token
```bash

$ curl -i -XGET -H "Content-type:application/json" \
-H "X-Authorization:0c768ad97c01cc31a0f7a93550611cd7d28c60e743262b132286325aa05a500f" \
localhost:8080/auth/token

```

**Out:**
```bash
{
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiMGM3NjhhZDk3YzAxY2MzMWEwZjdhOTM1NTA2MTFjZDxxxxxx",
    "expires": "2021-05-28 18:50:08",
    "message": "Welcome"
}

```
