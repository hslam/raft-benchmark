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
	Total time:	7.30s
	Requests per second:	137000.60
	Fastest time for request:	0.27ms
	Average time per request:	1.87ms
	Slowest time for request:	10.62ms

Time:
	0.1%	time for request:	0.71ms
	1%	time for request:	0.90ms
	5%	time for request:	1.12ms
	10%	time for request:	1.25ms
	25%	time for request:	1.48ms
	50%	time for request:	1.77ms
	75%	time for request:	2.13ms
	90%	time for request:	2.58ms
	95%	time for request:	2.93ms
	99%	time for request:	3.85ms
	99.9%	time for request:	5.80ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	411001.81 Byte/s (0.41 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	411001.81 Byte/s (0.41 MByte/s)

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
	Total time:	6.16s
	Requests per second:	162355.95
	Fastest time for request:	0.47ms
	Average time per request:	3.15ms
	Slowest time for request:	19.42ms

Time:
	0.1%	time for request:	1.11ms
	1%	time for request:	1.48ms
	5%	time for request:	1.86ms
	10%	time for request:	2.09ms
	25%	time for request:	2.51ms
	50%	time for request:	3.00ms
	75%	time for request:	3.59ms
	90%	time for request:	4.33ms
	95%	time for request:	4.93ms
	99%	time for request:	6.50ms
	99.9%	time for request:	10.21ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	487067.86 Byte/s (0.49 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	487067.86 Byte/s (0.49 MByte/s)

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
	Total time:	5.85s
	Requests per second:	170905.29
	Fastest time for request:	0.86ms
	Average time per request:	4.48ms
	Slowest time for request:	21.50ms

Time:
	0.1%	time for request:	1.62ms
	1%	time for request:	2.14ms
	5%	time for request:	2.71ms
	10%	time for request:	3.02ms
	25%	time for request:	3.60ms
	50%	time for request:	4.29ms
	75%	time for request:	5.09ms
	90%	time for request:	6.12ms
	95%	time for request:	6.93ms
	99%	time for request:	8.89ms
	99.9%	time for request:	14.28ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	512715.87 Byte/s (0.51 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	512715.87 Byte/s (0.51 MByte/s)

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
	Total time:	5.84s
	Requests per second:	171266.28
	Fastest time for request:	0.64ms
	Average time per request:	5.97ms
	Slowest time for request:	33.52ms

Time:
	0.1%	time for request:	2.12ms
	1%	time for request:	2.89ms
	5%	time for request:	3.62ms
	10%	time for request:	4.05ms
	25%	time for request:	4.81ms
	50%	time for request:	5.72ms
	75%	time for request:	6.81ms
	90%	time for request:	8.11ms
	95%	time for request:	9.17ms
	99%	time for request:	11.86ms
	99.9%	time for request:	17.08ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	513798.84 Byte/s (0.51 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	513798.84 Byte/s (0.51 MByte/s)

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
	Total time:	53.32s
	Requests per second:	18754.13
	Fastest time for request:	6.49ms
	Average time per request:	27.29ms
	Slowest time for request:	86.19ms

Time:
	0.1%	time for request:	12.03ms
	1%	time for request:	14.24ms
	5%	time for request:	17.09ms
	10%	time for request:	18.75ms
	25%	time for request:	22.35ms
	50%	time for request:	26.48ms
	75%	time for request:	30.86ms
	90%	time for request:	36.07ms
	95%	time for request:	40.22ms
	99%	time for request:	52.98ms
	99.9%	time for request:	71.20ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	675148.70 Byte/s (0.68 MByte/s)

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
	Total time:	31.76s
	Requests per second:	31483.99
	Fastest time for request:	8.15ms
	Average time per request:	32.51ms
	Slowest time for request:	115.29ms

Time:
	0.1%	time for request:	14.98ms
	1%	time for request:	18.71ms
	5%	time for request:	23.00ms
	10%	time for request:	24.88ms
	25%	time for request:	27.73ms
	50%	time for request:	31.20ms
	75%	time for request:	35.28ms
	90%	time for request:	41.09ms
	95%	time for request:	45.89ms
	99%	time for request:	65.08ms
	99.9%	time for request:	92.44ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1133423.60 Byte/s (1.13 MByte/s)

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
	Total time:	24.91s
	Requests per second:	40149.37
	Fastest time for request:	10.55ms
	Average time per request:	38.23ms
	Slowest time for request:	134.12ms

Time:
	0.1%	time for request:	18.36ms
	1%	time for request:	21.94ms
	5%	time for request:	27.08ms
	10%	time for request:	29.82ms
	25%	time for request:	33.25ms
	50%	time for request:	37.33ms
	75%	time for request:	41.22ms
	90%	time for request:	46.31ms
	95%	time for request:	51.54ms
	99%	time for request:	74.27ms
	99.9%	time for request:	110.63ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1445377.32 Byte/s (1.45 MByte/s)

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
	Total time:	22.37s
	Requests per second:	44695.31
	Fastest time for request:	9.96ms
	Average time per request:	45.71ms
	Slowest time for request:	183.20ms

Time:
	0.1%	time for request:	21.64ms
	1%	time for request:	26.81ms
	5%	time for request:	32.36ms
	10%	time for request:	35.00ms
	25%	time for request:	39.80ms
	50%	time for request:	44.67ms
	75%	time for request:	50.14ms
	90%	time for request:	54.82ms
	95%	time for request:	60.73ms
	99%	time for request:	82.13ms
	99.9%	time for request:	136.47ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1609031.17 Byte/s (1.61 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
