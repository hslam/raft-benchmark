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
	Total time:	7.38s
	Requests per second:	135590.15
	Fastest time for request:	0.41ms
	Average time per request:	1.89ms
	Slowest time for request:	11.43ms

Time:
	0.1%	time for request:	0.72ms
	1%	time for request:	0.92ms
	5%	time for request:	1.13ms
	10%	time for request:	1.26ms
	25%	time for request:	1.50ms
	50%	time for request:	1.78ms
	75%	time for request:	2.15ms
	90%	time for request:	2.60ms
	95%	time for request:	2.96ms
	99%	time for request:	3.97ms
	99.9%	time for request:	5.92ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	406770.45 Byte/s (0.41 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	406770.45 Byte/s (0.41 MByte/s)

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
	Total time:	6.21s
	Requests per second:	161159.68
	Fastest time for request:	0.62ms
	Average time per request:	3.17ms
	Slowest time for request:	12.42ms

Time:
	0.1%	time for request:	1.14ms
	1%	time for request:	1.50ms
	5%	time for request:	1.86ms
	10%	time for request:	2.10ms
	25%	time for request:	2.54ms
	50%	time for request:	3.05ms
	75%	time for request:	3.64ms
	90%	time for request:	4.34ms
	95%	time for request:	4.90ms
	99%	time for request:	6.27ms
	99.9%	time for request:	7.97ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	483479.04 Byte/s (0.48 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	483479.04 Byte/s (0.48 MByte/s)

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
	Total time:	5.83s
	Requests per second:	171529.73
	Fastest time for request:	0.42ms
	Average time per request:	4.47ms
	Slowest time for request:	24.57ms

Time:
	0.1%	time for request:	1.60ms
	1%	time for request:	2.17ms
	5%	time for request:	2.69ms
	10%	time for request:	3.00ms
	25%	time for request:	3.56ms
	50%	time for request:	4.26ms
	75%	time for request:	5.09ms
	90%	time for request:	6.09ms
	95%	time for request:	6.95ms
	99%	time for request:	9.21ms
	99.9%	time for request:	14.71ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	514589.20 Byte/s (0.51 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	514589.20 Byte/s (0.51 MByte/s)

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
	Total time:	5.56s
	Requests per second:	179958.65
	Fastest time for request:	0.36ms
	Average time per request:	5.68ms
	Slowest time for request:	19.46ms

Time:
	0.1%	time for request:	1.98ms
	1%	time for request:	2.77ms
	5%	time for request:	3.47ms
	10%	time for request:	3.85ms
	25%	time for request:	4.58ms
	50%	time for request:	5.46ms
	75%	time for request:	6.50ms
	90%	time for request:	7.74ms
	95%	time for request:	8.66ms
	99%	time for request:	10.74ms
	99.9%	time for request:	14.95ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	539875.96 Byte/s (0.54 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	539875.96 Byte/s (0.54 MByte/s)

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
	Total time:	52.68s
	Requests per second:	18982.24
	Fastest time for request:	6.65ms
	Average time per request:	26.97ms
	Slowest time for request:	208.06ms

Time:
	0.1%	time for request:	11.66ms
	1%	time for request:	14.76ms
	5%	time for request:	17.57ms
	10%	time for request:	19.22ms
	25%	time for request:	22.38ms
	50%	time for request:	25.66ms
	75%	time for request:	30.52ms
	90%	time for request:	35.52ms
	95%	time for request:	38.95ms
	99%	time for request:	50.01ms
	99.9%	time for request:	100.70ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	683360.51 Byte/s (0.68 MByte/s)

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
	Total time:	31.50s
	Requests per second:	31745.83
	Fastest time for request:	9.70ms
	Average time per request:	32.23ms
	Slowest time for request:	96.69ms

Time:
	0.1%	time for request:	15.72ms
	1%	time for request:	19.03ms
	5%	time for request:	23.05ms
	10%	time for request:	24.88ms
	25%	time for request:	27.85ms
	50%	time for request:	31.29ms
	75%	time for request:	35.17ms
	90%	time for request:	40.16ms
	95%	time for request:	43.89ms
	99%	time for request:	60.42ms
	99.9%	time for request:	86.68ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1142849.96 Byte/s (1.14 MByte/s)

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
	Total time:	25.43s
	Requests per second:	39319.77
	Fastest time for request:	12.10ms
	Average time per request:	39.01ms
	Slowest time for request:	212.75ms

Time:
	0.1%	time for request:	19.25ms
	1%	time for request:	24.67ms
	5%	time for request:	28.83ms
	10%	time for request:	30.75ms
	25%	time for request:	34.05ms
	50%	time for request:	37.85ms
	75%	time for request:	41.61ms
	90%	time for request:	47.34ms
	95%	time for request:	53.46ms
	99%	time for request:	71.25ms
	99.9%	time for request:	147.26ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1415511.76 Byte/s (1.42 MByte/s)

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
	Total time:	21.72s
	Requests per second:	46044.32
	Fastest time for request:	11.01ms
	Average time per request:	44.41ms
	Slowest time for request:	233.80ms

Time:
	0.1%	time for request:	20.70ms
	1%	time for request:	27.39ms
	5%	time for request:	31.56ms
	10%	time for request:	33.96ms
	25%	time for request:	38.03ms
	50%	time for request:	42.66ms
	75%	time for request:	47.52ms
	90%	time for request:	54.57ms
	95%	time for request:	63.97ms
	99%	time for request:	87.18ms
	99.9%	time for request:	217.03ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1657595.63 Byte/s (1.66 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
