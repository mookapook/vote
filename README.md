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


┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬───────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max   │
├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼───────┤
│ Latency │ 0 ms │ 0 ms │ 1 ms  │ 2 ms │ 0.16 ms │ 0.62 ms │ 43 ms │
└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴───────┘
┌───────────┬─────────┬─────────┬─────────┬─────────┬─────────┬────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg     │ Stdev  │ Min     │
├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼────────┼─────────┤
│ Req/Sec   │ 13,519  │ 13,519  │ 14,335  │ 17,487  │ 14,910  │ 1,168  │ 13,513  │
├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼────────┼─────────┤
│ Bytes/Sec │ 4.31 MB │ 4.31 MB │ 4.58 MB │ 5.58 MB │ 4.76 MB │ 373 kB │ 4.31 MB │
└───────────┴─────────┴─────────┴─────────┴─────────┴─────────┴────────┴─────────┘

Req/Bytes counts sampled once per second.
# of samples: 10

149k requests in 10.06s, 47.6 MB read
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


┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬───────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max   │
├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼───────┤
│ Latency │ 0 ms │ 0 ms │ 2 ms  │ 4 ms │ 0.51 ms │ 0.82 ms │ 12 ms │
└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴───────┘
┌───────────┬─────────┬─────────┬─────────┬─────────┬─────────┬──────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg     │ Stdev    │ Min     │
├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼──────────┼─────────┤
│ Req/Sec   │ 8,099   │ 8,099   │ 9,799   │ 11,647  │ 9,684   │ 1,039.86 │ 8,097   │
├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼──────────┼─────────┤
│ Bytes/Sec │ 1.45 MB │ 1.45 MB │ 1.75 MB │ 2.08 MB │ 1.73 MB │ 186 kB   │ 1.45 MB │
└───────────┴─────────┴─────────┴─────────┴─────────┴─────────┴──────────┴─────────┘

Req/Bytes counts sampled once per second.
# of samples: 11

107k requests in 11.02s, 19.1 MB read
```

# Item Get
```
autocannon -c 10 -d 10 -m GET  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzMjA4NzgsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0.7Axm0XyX2UgwFC_27RjOXAClYoX5saVWq88InYwgEm8"    "http://localhost:8080/v1/item"

Running 10s test @ http://localhost:8080/v1/item
10 connections


┌─────────┬───────┬───────┬───────┬────────┬──────────┬──────────┬────────┐
│ Stat    │ 2.5%  │ 50%   │ 97.5% │ 99%    │ Avg      │ Stdev    │ Max    │
├─────────┼───────┼───────┼───────┼────────┼──────────┼──────────┼────────┤
│ Latency │ 12 ms │ 28 ms │ 63 ms │ 101 ms │ 30.76 ms │ 21.31 ms │ 337 ms │
└─────────┴───────┴───────┴───────┴────────┴──────────┴──────────┴────────┘
┌───────────┬────────┬────────┬────────┬────────┬────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg    │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Req/Sec   │ 223    │ 223    │ 351    │ 365    │ 319.3  │ 54.11   │ 223    │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Bytes/Sec │ 332 kB │ 332 kB │ 522 kB │ 543 kB │ 474 kB │ 80.4 kB │ 331 kB │
└───────────┴────────┴────────┴────────┴────────┴────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.
# of samples: 10

3k requests in 10.02s, 4.74 MB read
```

# Vote By Item
```
autocannon -c 1 -d 1 -m POST  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzMjA4NzgsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0.7Axm0XyX2UgwFC_27RjOXAClYoX5saVWq88InYwgEm8"   -H "Content-Type: application/json"  "http://localhost:8080/v1/itemvote/65a7887f166933dd2b7a834f"
Running 1s test @ http://localhost:8080/v1/itemvote/65a7887f166933dd2b7a834f
1 connections


┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬───────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max   │
├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼───────┤
│ Latency │ 0 ms │ 0 ms │ 0 ms  │ 0 ms │ 0.02 ms │ 0.31 ms │ 11 ms │
└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴───────┘
┌───────────┬────────┬────────┬────────┬────────┬────────┬───────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg    │ Stdev │ Min    │
├───────────┼────────┼────────┼────────┼────────┼────────┼───────┼────────┤
│ Req/Sec   │ 4,163  │ 4,163  │ 4,163  │ 4,163  │ 4,162  │ 0     │ 4,161  │
├───────────┼────────┼────────┼────────┼────────┼────────┼───────┼────────┤
│ Bytes/Sec │ 762 kB │ 762 kB │ 762 kB │ 762 kB │ 762 kB │ 0 B   │ 761 kB │
└───────────┴────────┴────────┴────────┴────────┴────────┴───────┴────────┘

Req/Bytes counts sampled once per second.
# of samples: 1

4k requests in 1.01s, 761 kB read
```


