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
	Total time:	7.76s
	Requests per second:	128826.87
	Fastest time for request:	0.36ms
	Average time per request:	1.98ms
	Slowest time for request:	13.75ms

Time:
	0.1%	time for request:	0.82ms
	1%	time for request:	1.02ms
	5%	time for request:	1.22ms
	10%	time for request:	1.33ms
	25%	time for request:	1.56ms
	50%	time for request:	1.87ms
	75%	time for request:	2.27ms
	90%	time for request:	2.74ms
	95%	time for request:	3.12ms
	99%	time for request:	4.14ms
	99.9%	time for request:	6.66ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	386480.60 Byte/s (0.39 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	386480.60 Byte/s (0.39 MByte/s)

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
	Total time:	6.43s
	Requests per second:	155492.61
	Fastest time for request:	0.62ms
	Average time per request:	3.29ms
	Slowest time for request:	16.26ms

Time:
	0.1%	time for request:	1.37ms
	1%	time for request:	1.74ms
	5%	time for request:	2.10ms
	10%	time for request:	2.31ms
	25%	time for request:	2.69ms
	50%	time for request:	3.17ms
	75%	time for request:	3.73ms
	90%	time for request:	4.38ms
	95%	time for request:	4.86ms
	99%	time for request:	6.11ms
	99.9%	time for request:	8.48ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	466477.82 Byte/s (0.47 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	466477.82 Byte/s (0.47 MByte/s)

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
	Total time:	6.01s
	Requests per second:	166405.88
	Fastest time for request:	0.93ms
	Average time per request:	4.61ms
	Slowest time for request:	16.86ms

Time:
	0.1%	time for request:	1.92ms
	1%	time for request:	2.52ms
	5%	time for request:	3.03ms
	10%	time for request:	3.31ms
	25%	time for request:	3.81ms
	50%	time for request:	4.43ms
	75%	time for request:	5.19ms
	90%	time for request:	6.11ms
	95%	time for request:	6.81ms
	99%	time for request:	8.43ms
	99.9%	time for request:	11.71ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	499217.64 Byte/s (0.50 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	499217.64 Byte/s (0.50 MByte/s)

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
	Total time:	5.82s
	Requests per second:	171728.27
	Fastest time for request:	0.95ms
	Average time per request:	5.95ms
	Slowest time for request:	24.87ms

Time:
	0.1%	time for request:	2.43ms
	1%	time for request:	3.22ms
	5%	time for request:	3.90ms
	10%	time for request:	4.27ms
	25%	time for request:	4.94ms
	50%	time for request:	5.78ms
	75%	time for request:	6.71ms
	90%	time for request:	7.80ms
	95%	time for request:	8.63ms
	99%	time for request:	10.54ms
	99.9%	time for request:	13.88ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	515184.81 Byte/s (0.52 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	515184.81 Byte/s (0.52 MByte/s)

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
	Total time:	55.54s
	Requests per second:	18004.40
	Fastest time for request:	7.35ms
	Average time per request:	28.43ms
	Slowest time for request:	250.55ms

Time:
	0.1%	time for request:	11.91ms
	1%	time for request:	15.64ms
	5%	time for request:	19.10ms
	10%	time for request:	20.74ms
	25%	time for request:	23.82ms
	50%	time for request:	27.15ms
	75%	time for request:	31.37ms
	90%	time for request:	37.11ms
	95%	time for request:	41.57ms
	99%	time for request:	54.20ms
	99.9%	time for request:	70.83ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	648158.49 Byte/s (0.65 MByte/s)

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
	Total time:	34.17s
	Requests per second:	29266.09
	Fastest time for request:	9.24ms
	Average time per request:	34.96ms
	Slowest time for request:	133.14ms

Time:
	0.1%	time for request:	15.72ms
	1%	time for request:	20.45ms
	5%	time for request:	25.13ms
	10%	time for request:	27.36ms
	25%	time for request:	30.45ms
	50%	time for request:	33.78ms
	75%	time for request:	37.91ms
	90%	time for request:	43.00ms
	95%	time for request:	47.75ms
	99%	time for request:	63.60ms
	99.9%	time for request:	101.46ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1053579.25 Byte/s (1.05 MByte/s)

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
	Total time:	27.11s
	Requests per second:	36880.55
	Fastest time for request:	11.90ms
	Average time per request:	41.61ms
	Slowest time for request:	136.57ms

Time:
	0.1%	time for request:	21.09ms
	1%	time for request:	24.81ms
	5%	time for request:	30.46ms
	10%	time for request:	32.69ms
	25%	time for request:	35.87ms
	50%	time for request:	40.25ms
	75%	time for request:	45.02ms
	90%	time for request:	51.27ms
	95%	time for request:	55.89ms
	99%	time for request:	81.90ms
	99.9%	time for request:	122.77ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1327699.71 Byte/s (1.33 MByte/s)

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
	Total time:	24.17s
	Requests per second:	41375.58
	Fastest time for request:	17.48ms
	Average time per request:	49.42ms
	Slowest time for request:	147.35ms

Time:
	0.1%	time for request:	25.30ms
	1%	time for request:	30.96ms
	5%	time for request:	36.77ms
	10%	time for request:	39.24ms
	25%	time for request:	42.94ms
	50%	time for request:	47.42ms
	75%	time for request:	53.50ms
	90%	time for request:	60.29ms
	95%	time for request:	67.08ms
	99%	time for request:	97.55ms
	99.9%	time for request:	139.60ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1489520.73 Byte/s (1.49 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
