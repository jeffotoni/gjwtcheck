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
    "key": "<token-jwt>",
    "expires": "<2021-05-26 10:26:36>"
}

```

## hs256/user
```bash

$ curl -i -XPOST -H "Content-type:application/json" \
-H "Authorization: Bearer $token2" 
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
   "user":"jeff@gmail.com",
   "id":"447b22be-2d17-4253-9f4a-77a8501ef168",
   "iss":"gjwtcheck - created in:2021-12-01 00:01:42 expires:2021-12-01 00:05:42",
   "user_avatar":"https://logodix.com/logo/1989600.png",
   "message":"seja bem JWT HS256"
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
    "key": "<token-jwt>",
    "expires": "<2021-05-26 10:26:36>"
}

```

## rs256/user
```bash

$ curl -i -XPOST -H "Content-type:application/json" \
-H "Authorization: Bearer $token2" 
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
   "user":"jeff@gmail.com",
   "id":"447b22be-2d17-4253-9f4a-77a8501ef168",
   "iss":"gjwtcheck - created in:2021-12-01 00:01:42 expires:2021-12-01 00:05:42",
   "user_avatar":"https://logodix.com/logo/1989600.png",
   "message":"seja bem JWT RS256"
}

```
