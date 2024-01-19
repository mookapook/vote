# Use Test
```
username =  Alice
password = 123456

username =  Lamda
password = 123456x

username =  mookrob
password = 123456x7
```


# Open Terminal
```
cd path/project

copy file docker-compose ออกไป 1 ชั้น

cd .. // cd go path up level path project

docker-compose up -d
```
# Document API
```
<https://documenter.getpostman.com/view/4032383/2s9YsRc9Em>
```

# API Endpoint
```

GET /v1/item
GET /v1/export
GET /v1/exporvoteitem/:id

POST /v1/login
POST /v1/itemcreate
POST /v1/itemvote/:id //vote

PUT /v1/itemvote/:id  //Unvote
PUT  /v1/itemcreate/:id
PUT  /v1/itemopenclose/:id
PUT  /v1/itemclearbyid/:id
PUT  /v1/itemclear

DELETE /v1/itemcreate/:id

```

# Performace Test
```
Use Tool autocannon

install npm i autocannon -g

-c | --connections
Number of concurrent connections to use
By default, its value is 10
-p | --pipeline
Number of pipelined requests to use
By default, its value is 1
-d | --duration
Number of seconds to run the autocannon
By default, its value is 10
-w | --workers
Number of worker threads to fire requests
-m | --method
HTTP method to use
By default, its value is 'GET'
-t | --timeout
Number of seconds before timing out and resetting a connection
By default, its value is 10
-j | --json
Print the output as newline delimited JSON
By default, its value is false
-f | --forever
Run the benchmark forever

```

# Login Test
```
autocannon -c 10 -d 10 -m POST -H "Content-Type: application/json" -b '{"username": "Alice","password":"123456"}' "http://localhost:8080/v1/login"
Running 10s test @ http://localhost:8080/v1/login
10 connections


┌─────────┬──────┬──────┬───────┬───────┬─────────┬─────────┬───────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%   │ Avg     │ Stdev   │ Max   │
├─────────┼──────┼──────┼───────┼───────┼─────────┼─────────┼───────┤
│ Latency │ 1 ms │ 3 ms │ 12 ms │ 16 ms │ 4.19 ms │ 3.07 ms │ 56 ms │
└─────────┴──────┴──────┴───────┴───────┴─────────┴─────────┴───────┘
┌───────────┬────────┬────────┬────────┬────────┬─────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg     │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼─────────┼─────────┼────────┤
│ Req/Sec   │ 1,705  │ 1,705  │ 2,137  │ 2,361  │ 2,137.4 │ 184.58  │ 1,705  │
├───────────┼────────┼────────┼────────┼────────┼─────────┼─────────┼────────┤
│ Bytes/Sec │ 544 kB │ 544 kB │ 681 kB │ 753 kB │ 682 kB  │ 58.7 kB │ 544 kB │
└───────────┴────────┴────────┴────────┴────────┴─────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.
# of samples: 10

21k requests in 10.09s, 6.82 MB read
```


# Item Create
```
autocannon -c 10 -d 10 -m POST  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzMjA4NzgsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0.7Axm0XyX2UgwFC_27RjOXAClYoX5saVWq88InYwgEm8"   -H "Content-Type: application/json" -b '{"name": "test Autuconn","description":"test Autocannon"}' "http://localhost:8080/v1/itemcreate"
Running 10s test @ http://localhost:8080/v1/itemcreate
10 connections


┌─────────┬──────┬──────┬───────┬───────┬────────┬─────────┬───────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%   │ Avg    │ Stdev   │ Max   │
├─────────┼──────┼──────┼───────┼───────┼────────┼─────────┼───────┤
│ Latency │ 2 ms │ 4 ms │ 12 ms │ 15 ms │ 4.9 ms │ 2.65 ms │ 39 ms │
└─────────┴──────┴──────┴───────┴───────┴────────┴─────────┴───────┘
┌───────────┬────────┬────────┬────────┬────────┬─────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg     │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼─────────┼─────────┼────────┤
│ Req/Sec   │ 1,606  │ 1,606  │ 1,880  │ 1,993  │ 1,847.5 │ 126.88  │ 1,606  │
├───────────┼────────┼────────┼────────┼────────┼─────────┼─────────┼────────┤
│ Bytes/Sec │ 492 kB │ 492 kB │ 575 kB │ 610 kB │ 565 kB  │ 38.9 kB │ 491 kB │
└───────────┴────────┴────────┴────────┴────────┴─────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.
# of samples: 10

18k requests in 10.02s, 5.65 MB read
```

# Item Update
```
autocannon -c 10 -d 10 -m PUT  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzMjA4NzgsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0.7Axm0XyX2UgwFC_27RjOXAClYoX5saVWq88InYwgEm8"   -H "Content-Type: application/json" -b '{"name": "test Autuconn update","description":"test Autocannon"}' "http://localhost:8080/v1/itemcreate/65a89b41ff1dc45c844c523d"
Running 10s test @ http://localhost:8080/v1/itemcreate/65a89b41ff1dc45c844c523d
10 connections


┌─────────┬──────┬──────┬───────┬───────┬─────────┬─────────┬───────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%   │ Avg     │ Stdev   │ Max   │
├─────────┼──────┼──────┼───────┼───────┼─────────┼─────────┼───────┤
│ Latency │ 2 ms │ 6 ms │ 22 ms │ 27 ms │ 7.25 ms │ 5.45 ms │ 66 ms │
└─────────┴──────┴──────┴───────┴───────┴─────────┴─────────┴───────┘
┌───────────┬────────┬────────┬────────┬────────┬──────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg      │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼──────────┼─────────┼────────┤
│ Req/Sec   │ 982    │ 982    │ 1,316  │ 1,541  │ 1,290.41 │ 187.56  │ 982    │
├───────────┼────────┼────────┼────────┼────────┼──────────┼─────────┼────────┤
│ Bytes/Sec │ 176 kB │ 176 kB │ 236 kB │ 276 kB │ 231 kB   │ 33.6 kB │ 176 kB │
└───────────┴────────┴────────┴────────┴────────┴──────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.
# of samples: 10

13k requests in 10.04s, 2.31 MB read
```

# Item Get
```
autocannon -c 10 -d 10 -m GET  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzMjA4NzgsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0.7Axm0XyX2UgwFC_27RjOXAClYoX5saVWq88InYwgEm8"    "http://localhost:8080/v1/item"
Running 10s test @ http://localhost:8080/v1/item
10 connections


┌─────────┬───────┬───────┬───────┬───────┬──────────┬──────────┬────────┐
│ Stat    │ 2.5%  │ 50%   │ 97.5% │ 99%   │ Avg      │ Stdev    │ Max    │
├─────────┼───────┼───────┼───────┼───────┼──────────┼──────────┼────────┤
│ Latency │ 15 ms │ 35 ms │ 72 ms │ 81 ms │ 37.01 ms │ 15.45 ms │ 176 ms │
└─────────┴───────┴───────┴───────┴───────┴──────────┴──────────┴────────┘
┌───────────┬────────┬────────┬────────┬────────┬────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg    │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Req/Sec   │ 245    │ 245    │ 264    │ 283    │ 265.9  │ 13.3    │ 245    │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Bytes/Sec │ 365 kB │ 365 kB │ 393 kB │ 422 kB │ 396 kB │ 19.8 kB │ 365 kB │
└───────────┴────────┴────────┴────────┴────────┴────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.
# of samples: 10

3k requests in 10.04s, 3.96 MB read
```

# Vote By Item
```
autocannon -c 10 -d 10 -m POST  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzMjA4NzgsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0.7Axm0XyX2UgwFC_27RjOXAClYoX5saVWq88InYwgEm8"   -H "Content-Type: application/json"  "http://localhost:8080/v1/itemvote/65a7887f166933dd2b7a834f"
Running 10s test @ http://localhost:8080/v1/itemvote/65a7887f166933dd2b7a834f
10 connections


┌─────────┬──────┬──────┬───────┬───────┬─────────┬─────────┬───────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%   │ Avg     │ Stdev   │ Max   │
├─────────┼──────┼──────┼───────┼───────┼─────────┼─────────┼───────┤
│ Latency │ 1 ms │ 5 ms │ 17 ms │ 21 ms │ 5.95 ms │ 4.31 ms │ 41 ms │
└─────────┴──────┴──────┴───────┴───────┴─────────┴─────────┴───────┘
┌───────────┬────────┬────────┬────────┬────────┬─────────┬───────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg     │ Stdev │ Min    │
├───────────┼────────┼────────┼────────┼────────┼─────────┼───────┼────────┤
│ Req/Sec   │ 1,425  │ 1,425  │ 1,523  │ 1,708  │ 1,549.3 │ 87.41 │ 1,425  │
├───────────┼────────┼────────┼────────┼────────┼─────────┼───────┼────────┤
│ Bytes/Sec │ 261 kB │ 261 kB │ 279 kB │ 313 kB │ 284 kB  │ 16 kB │ 261 kB │
└───────────┴────────┴────────┴────────┴────────┴─────────┴───────┴────────┘

Req/Bytes counts sampled once per second.
# of samples: 10

16k requests in 10.03s, 2.84 MB read
```


