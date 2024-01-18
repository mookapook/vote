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
