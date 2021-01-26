# ETCD

curl -XPUT http://127.0.0.1:2379/v2/keys/foo  -d value="bar"
```
{"action":"set","node":{"key":"/foo","value":"bar","modifiedIndex":8,"createdIndex":8}}
```

curl http://127.0.0.1:2379/v2/keys/foo
```
{"action":"get","node":{"key":"/foo","value":"bar","modifiedIndex":8,"createdIndex":8}}
```

## ETCD Read Index

./etcd_read_index -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=256 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	256
	Total calls:	1000000
	Total time:	36.22s
	Requests per second:	27608.45
	Fastest time for request:	0.38ms
	Average time per request:	9.27ms
	Slowest time for request:	74.68ms

Time:
	0.1%	time for request:	0.86ms
	1%	time for request:	1.65ms
	5%	time for request:	2.77ms
	10%	time for request:	3.56ms
	25%	time for request:	5.14ms
	50%	time for request:	7.41ms
	75%	time for request:	11.60ms
	90%	time for request:	17.75ms
	95%	time for request:	22.04ms
	99%	time for request:	30.42ms
	99.9%	time for request:	42.64ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	82825.36 Byte/s (0.08 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	82825.36 Byte/s (0.08 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./etcd_read_index -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=512 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	512
	Total calls:	1000000
	Total time:	32.27s
	Requests per second:	30987.85
	Fastest time for request:	0.36ms
	Average time per request:	16.51ms
	Slowest time for request:	99.08ms

Time:
	0.1%	time for request:	1.08ms
	1%	time for request:	2.30ms
	5%	time for request:	4.15ms
	10%	time for request:	5.56ms
	25%	time for request:	8.85ms
	50%	time for request:	13.81ms
	75%	time for request:	21.60ms
	90%	time for request:	31.03ms
	95%	time for request:	37.75ms
	99%	time for request:	52.55ms
	99.9%	time for request:	67.47ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	92963.56 Byte/s (0.09 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	92963.56 Byte/s (0.09 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./etcd_read_index -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=768 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	768
	Total calls:	1000000
	Total time:	30.75s
	Requests per second:	32522.09
	Fastest time for request:	0.38ms
	Average time per request:	23.61ms
	Slowest time for request:	120.09ms

Time:
	0.1%	time for request:	1.40ms
	1%	time for request:	3.15ms
	5%	time for request:	5.94ms
	10%	time for request:	8.08ms
	25%	time for request:	13.04ms
	50%	time for request:	20.24ms
	75%	time for request:	30.88ms
	90%	time for request:	43.47ms
	95%	time for request:	52.60ms
	99%	time for request:	71.34ms
	99.9%	time for request:	93.02ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	97566.26 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	97566.26 Byte/s (0.10 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./etcd_read_index -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=1024 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	1024
	Total calls:	1000000
	Total time:	30.03s
	Requests per second:	33295.11
	Fastest time for request:	0.40ms
	Average time per request:	30.75ms
	Slowest time for request:	164.04ms

Time:
	0.1%	time for request:	1.97ms
	1%	time for request:	4.38ms
	5%	time for request:	8.27ms
	10%	time for request:	11.27ms
	25%	time for request:	17.45ms
	50%	time for request:	27.45ms
	75%	time for request:	40.13ms
	90%	time for request:	54.45ms
	95%	time for request:	64.87ms
	99%	time for request:	87.62ms
	99.9%	time for request:	115.31ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	99885.34 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	99885.34 Byte/s (0.10 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

## ETCD Write

./client -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=512 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	512
	Total calls:	1000000
	Total time:	172.02s
	Requests per second:	5813.37
	Fastest time for request:	18.92ms
	Average time per request:	88.05ms
	Slowest time for request:	379.56ms

Time:
	0.1%	time for request:	28.33ms
	1%	time for request:	36.59ms
	5%	time for request:	45.57ms
	10%	time for request:	51.24ms
	25%	time for request:	65.88ms
	50%	time for request:	88.53ms
	75%	time for request:	105.11ms
	90%	time for request:	122.61ms
	95%	time for request:	133.62ms
	99%	time for request:	166.44ms
	99.9%	time for request:	254.37ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	209281.22 Byte/s (0.21 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./client -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=1024 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	1024
	Total calls:	1000000
	Total time:	103.69s
	Requests per second:	9643.77
	Fastest time for request:	20.00ms
	Average time per request:	106.13ms
	Slowest time for request:	475.02ms

Time:
	0.1%	time for request:	28.44ms
	1%	time for request:	36.01ms
	5%	time for request:	46.58ms
	10%	time for request:	54.22ms
	25%	time for request:	72.76ms
	50%	time for request:	97.57ms
	75%	time for request:	124.12ms
	90%	time for request:	161.42ms
	95%	time for request:	211.98ms
	99%	time for request:	297.15ms
	99.9%	time for request:	398.05ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	347175.74 Byte/s (0.35 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./client -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=1536 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	1536
	Total calls:	1000000
	Total time:	85.31s
	Requests per second:	11722.62
	Fastest time for request:	19.50ms
	Average time per request:	130.97ms
	Slowest time for request:	612.36ms

Time:
	0.1%	time for request:	27.60ms
	1%	time for request:	36.75ms
	5%	time for request:	49.48ms
	10%	time for request:	59.38ms
	25%	time for request:	80.58ms
	50%	time for request:	111.83ms
	75%	time for request:	149.24ms
	90%	time for request:	240.12ms
	95%	time for request:	308.71ms
	99%	time for request:	408.42ms
	99.9%	time for request:	511.51ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	422014.38 Byte/s (0.42 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./client -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=2048 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	2048
	Total calls:	1000000
	Total time:	73.66s
	Requests per second:	13575.50
	Fastest time for request:	16.99ms
	Average time per request:	150.72ms
	Slowest time for request:	652.18ms

Time:
	0.1%	time for request:	28.68ms
	1%	time for request:	39.27ms
	5%	time for request:	54.02ms
	10%	time for request:	64.29ms
	25%	time for request:	88.33ms
	50%	time for request:	123.89ms
	75%	time for request:	178.61ms
	90%	time for request:	297.37ms
	95%	time for request:	359.36ms
	99%	time for request:	450.76ms
	99.9%	time for request:	539.55ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	488717.96 Byte/s (0.49 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```