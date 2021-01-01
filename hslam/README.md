# HSLAM

curl -XPOST http://localhost:7001/db/foo -d 'bar'

curl http://localhost:7001/db/foo
```
bar
```

## HSLAM Read Index

./rpc_read_index -addrs=:8001,:8002,:8003 -total=1000000 -parallel=256 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	256
	Total calls:	1000000
	Total time:	8.17s
	Requests per second:	122403.55
	Fastest time for request:	0.46ms
	Average time per request:	2.09ms
	Slowest time for request:	10.26ms

Time:
	0.1%	time for request:	0.86ms
	1%	time for request:	1.08ms
	5%	time for request:	1.30ms
	10%	time for request:	1.43ms
	25%	time for request:	1.67ms
	50%	time for request:	1.97ms
	75%	time for request:	2.36ms
	90%	time for request:	2.87ms
	95%	time for request:	3.32ms
	99%	time for request:	4.33ms
	99.9%	time for request:	5.83ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	367210.64 Byte/s (0.37 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	367210.64 Byte/s (0.37 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_read_index -addrs=:8001,:8002,:8003 -total=1000000 -parallel=512 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	512
	Total calls:	1000000
	Total time:	6.91s
	Requests per second:	144720.50
	Fastest time for request:	0.65ms
	Average time per request:	3.53ms
	Slowest time for request:	16.75ms

Time:
	0.1%	time for request:	1.41ms
	1%	time for request:	1.85ms
	5%	time for request:	2.23ms
	10%	time for request:	2.46ms
	25%	time for request:	2.84ms
	50%	time for request:	3.31ms
	75%	time for request:	3.93ms
	90%	time for request:	4.86ms
	95%	time for request:	5.60ms
	99%	time for request:	7.56ms
	99.9%	time for request:	10.95ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	434161.51 Byte/s (0.43 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	434161.51 Byte/s (0.43 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_read_index -addrs=:8001,:8002,:8003 -total=1000000 -parallel=768 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	768
	Total calls:	1000000
	Total time:	6.18s
	Requests per second:	161938.63
	Fastest time for request:	0.91ms
	Average time per request:	4.74ms
	Slowest time for request:	14.63ms

Time:
	0.1%	time for request:	1.98ms
	1%	time for request:	2.55ms
	5%	time for request:	3.07ms
	10%	time for request:	3.36ms
	25%	time for request:	3.88ms
	50%	time for request:	4.51ms
	75%	time for request:	5.33ms
	90%	time for request:	6.43ms
	95%	time for request:	7.22ms
	99%	time for request:	8.89ms
	99.9%	time for request:	11.15ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	485815.88 Byte/s (0.49 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	485815.88 Byte/s (0.49 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_read_index -addrs=:8001,:8002,:8003 -total=1000000 -parallel=1024 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	1024
	Total calls:	1000000
	Total time:	6.02s
	Requests per second:	166126.31
	Fastest time for request:	1.39ms
	Average time per request:	6.16ms
	Slowest time for request:	23.11ms

Time:
	0.1%	time for request:	2.46ms
	1%	time for request:	3.27ms
	5%	time for request:	3.92ms
	10%	time for request:	4.29ms
	25%	time for request:	4.99ms
	50%	time for request:	5.85ms
	75%	time for request:	6.93ms
	90%	time for request:	8.44ms
	95%	time for request:	9.51ms
	99%	time for request:	11.88ms
	99.9%	time for request:	15.06ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	498378.94 Byte/s (0.50 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	498378.94 Byte/s (0.50 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

## HSLAM Write

./rpc_write -addrs=:8001,:8002,:8003 -total=1000000 -parallel=512 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	512
	Total calls:	1000000
	Total time:	55.48s
	Requests per second:	18023.89
	Fastest time for request:	7.58ms
	Average time per request:	28.39ms
	Slowest time for request:	82.70ms

Time:
	0.1%	time for request:	12.95ms
	1%	time for request:	15.90ms
	5%	time for request:	18.82ms
	10%	time for request:	20.73ms
	25%	time for request:	23.95ms
	50%	time for request:	27.18ms
	75%	time for request:	31.30ms
	90%	time for request:	37.34ms
	95%	time for request:	42.10ms
	99%	time for request:	53.99ms
	99.9%	time for request:	71.91ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	648860.07 Byte/s (0.65 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_write -addrs=:8001,:8002,:8003 -total=1000000 -parallel=1024 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	1024
	Total calls:	1000000
	Total time:	33.66s
	Requests per second:	29704.91
	Fastest time for request:	9.97ms
	Average time per request:	34.45ms
	Slowest time for request:	102.04ms

Time:
	0.1%	time for request:	16.51ms
	1%	time for request:	20.53ms
	5%	time for request:	24.92ms
	10%	time for request:	26.95ms
	25%	time for request:	29.69ms
	50%	time for request:	33.20ms
	75%	time for request:	37.39ms
	90%	time for request:	42.91ms
	95%	time for request:	48.02ms
	99%	time for request:	63.83ms
	99.9%	time for request:	84.64ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1069376.61 Byte/s (1.07 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_write -addrs=:8001,:8002,:8003 -total=1000000 -parallel=1536 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	1536
	Total calls:	1000000
	Total time:	27.05s
	Requests per second:	36967.28
	Fastest time for request:	10.68ms
	Average time per request:	41.51ms
	Slowest time for request:	124.12ms

Time:
	0.1%	time for request:	19.90ms
	1%	time for request:	24.33ms
	5%	time for request:	29.93ms
	10%	time for request:	32.19ms
	25%	time for request:	35.58ms
	50%	time for request:	39.77ms
	75%	time for request:	44.65ms
	90%	time for request:	51.81ms
	95%	time for request:	59.68ms
	99%	time for request:	81.13ms
	99.9%	time for request:	115.24ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1330822.17 Byte/s (1.33 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_write -addrs=:8001,:8002,:8003 -total=1000000 -parallel=2048 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	2048
	Total calls:	1000000
	Total time:	23.44s
	Requests per second:	42667.57
	Fastest time for request:	14.47ms
	Average time per request:	47.93ms
	Slowest time for request:	139.54ms

Time:
	0.1%	time for request:	24.11ms
	1%	time for request:	28.48ms
	5%	time for request:	35.15ms
	10%	time for request:	37.77ms
	25%	time for request:	41.01ms
	50%	time for request:	45.57ms
	75%	time for request:	51.98ms
	90%	time for request:	60.71ms
	95%	time for request:	68.91ms
	99%	time for request:	91.02ms
	99.9%	time for request:	118.53ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1536032.57 Byte/s (1.54 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
