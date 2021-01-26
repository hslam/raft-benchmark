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
	Total time:	7.63s
	Requests per second:	131101.07
	Fastest time for request:	0.43ms
	Average time per request:	1.95ms
	Slowest time for request:	19.56ms

Time:
	0.1%	time for request:	0.80ms
	1%	time for request:	1.01ms
	5%	time for request:	1.21ms
	10%	time for request:	1.33ms
	25%	time for request:	1.56ms
	50%	time for request:	1.86ms
	75%	time for request:	2.22ms
	90%	time for request:	2.67ms
	95%	time for request:	3.01ms
	99%	time for request:	3.87ms
	99.9%	time for request:	5.49ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	393303.20 Byte/s (0.39 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	393303.20 Byte/s (0.39 MByte/s)

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
	Total time:	6.44s
	Requests per second:	155314.16
	Fastest time for request:	0.70ms
	Average time per request:	3.29ms
	Slowest time for request:	11.63ms

Time:
	0.1%	time for request:	1.34ms
	1%	time for request:	1.72ms
	5%	time for request:	2.09ms
	10%	time for request:	2.30ms
	25%	time for request:	2.68ms
	50%	time for request:	3.17ms
	75%	time for request:	3.75ms
	90%	time for request:	4.40ms
	95%	time for request:	4.86ms
	99%	time for request:	6.13ms
	99.9%	time for request:	8.23ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	465942.48 Byte/s (0.47 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	465942.48 Byte/s (0.47 MByte/s)

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
	Total time:	5.90s
	Requests per second:	169466.65
	Fastest time for request:	1.02ms
	Average time per request:	4.53ms
	Slowest time for request:	15.25ms

Time:
	0.1%	time for request:	1.87ms
	1%	time for request:	2.46ms
	5%	time for request:	2.98ms
	10%	time for request:	3.26ms
	25%	time for request:	3.76ms
	50%	time for request:	4.39ms
	75%	time for request:	5.13ms
	90%	time for request:	5.93ms
	95%	time for request:	6.50ms
	99%	time for request:	8.04ms
	99.9%	time for request:	11.08ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	508399.95 Byte/s (0.51 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	508399.95 Byte/s (0.51 MByte/s)

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
	Total time:	5.78s
	Requests per second:	173035.71
	Fastest time for request:	0.93ms
	Average time per request:	5.91ms
	Slowest time for request:	26.68ms

Time:
	0.1%	time for request:	2.36ms
	1%	time for request:	3.21ms
	5%	time for request:	3.90ms
	10%	time for request:	4.25ms
	25%	time for request:	4.88ms
	50%	time for request:	5.71ms
	75%	time for request:	6.65ms
	90%	time for request:	7.77ms
	95%	time for request:	8.59ms
	99%	time for request:	11.01ms
	99.9%	time for request:	15.62ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	519107.12 Byte/s (0.52 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	519107.12 Byte/s (0.52 MByte/s)

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
	Total time:	55.28s
	Requests per second:	18090.58
	Fastest time for request:	6.52ms
	Average time per request:	28.29ms
	Slowest time for request:	128.99ms

Time:
	0.1%	time for request:	12.41ms
	1%	time for request:	15.38ms
	5%	time for request:	18.75ms
	10%	time for request:	20.60ms
	25%	time for request:	23.62ms
	50%	time for request:	27.11ms
	75%	time for request:	31.27ms
	90%	time for request:	37.27ms
	95%	time for request:	41.73ms
	99%	time for request:	54.66ms
	99.9%	time for request:	78.31ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	651260.87 Byte/s (0.65 MByte/s)

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
	Total time:	34.04s
	Requests per second:	29380.82
	Fastest time for request:	9.48ms
	Average time per request:	34.83ms
	Slowest time for request:	199.21ms

Time:
	0.1%	time for request:	16.39ms
	1%	time for request:	20.19ms
	5%	time for request:	25.16ms
	10%	time for request:	26.91ms
	25%	time for request:	30.03ms
	50%	time for request:	33.22ms
	75%	time for request:	37.71ms
	90%	time for request:	42.92ms
	95%	time for request:	48.34ms
	99%	time for request:	68.72ms
	99.9%	time for request:	118.95ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1057709.70 Byte/s (1.06 MByte/s)

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
	Total time:	26.89s
	Requests per second:	37184.96
	Fastest time for request:	12.87ms
	Average time per request:	41.27ms
	Slowest time for request:	120.37ms

Time:
	0.1%	time for request:	18.56ms
	1%	time for request:	23.88ms
	5%	time for request:	29.80ms
	10%	time for request:	32.12ms
	25%	time for request:	35.79ms
	50%	time for request:	39.80ms
	75%	time for request:	44.72ms
	90%	time for request:	51.05ms
	95%	time for request:	56.43ms
	99%	time for request:	77.69ms
	99.9%	time for request:	110.10ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1338658.47 Byte/s (1.34 MByte/s)

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
	Total time:	23.84s
	Requests per second:	41941.76
	Fastest time for request:	15.69ms
	Average time per request:	48.76ms
	Slowest time for request:	225.43ms

Time:
	0.1%	time for request:	24.36ms
	1%	time for request:	30.62ms
	5%	time for request:	35.97ms
	10%	time for request:	38.26ms
	25%	time for request:	42.39ms
	50%	time for request:	46.60ms
	75%	time for request:	52.41ms
	90%	time for request:	59.14ms
	95%	time for request:	66.52ms
	99%	time for request:	97.41ms
	99.9%	time for request:	139.27ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1509903.33 Byte/s (1.51 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
