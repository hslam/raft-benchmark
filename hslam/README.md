# HSLAM

curl -XPOST http://localhost:7001/db/foo -d 'bar'

curl http://localhost:7001/db/foo
```
bar
```

## HSLAM Read Index

./rpc_read_index -network=tcp -codec=pb -addr=:8001 -parallel=128 -total=1000000 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	128
	Total calls:	1000000
	Total time:	9.95s
	Requests per second:	100534.83
	Fastest time for request:	0.35ms
	Average time per request:	1.27ms
	Slowest time for request:	6.40ms

Time:
	0.1%	time for request:	0.55ms
	1%	time for request:	0.69ms
	5%	time for request:	0.83ms
	10%	time for request:	0.90ms
	25%	time for request:	1.03ms
	50%	time for request:	1.21ms
	75%	time for request:	1.42ms
	90%	time for request:	1.69ms
	95%	time for request:	1.92ms
	99%	time for request:	2.69ms
	99.9%	time for request:	3.75ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	301604.48 Byte/s (0.30 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	301604.48 Byte/s (0.30 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_read_index -network=tcp -codec=pb -addr=:8001 -parallel=256 -total=1000000 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	256
	Total calls:	1000000
	Total time:	7.96s
	Requests per second:	125659.82
	Fastest time for request:	0.36ms
	Average time per request:	2.04ms
	Slowest time for request:	8.07ms

Time:
	0.1%	time for request:	0.79ms
	1%	time for request:	1.01ms
	5%	time for request:	1.25ms
	10%	time for request:	1.38ms
	25%	time for request:	1.62ms
	50%	time for request:	1.92ms
	75%	time for request:	2.32ms
	90%	time for request:	2.81ms
	95%	time for request:	3.22ms
	99%	time for request:	4.17ms
	99.9%	time for request:	5.50ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	376979.47 Byte/s (0.38 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	376979.47 Byte/s (0.38 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_read_index -network=tcp -codec=pb -addr=:8001 -parallel=384 -total=1000000 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	384
	Total calls:	1000000
	Total time:	7.69s
	Requests per second:	129966.79
	Fastest time for request:	0.57ms
	Average time per request:	2.95ms
	Slowest time for request:	29.50ms

Time:
	0.1%	time for request:	1.06ms
	1%	time for request:	1.38ms
	5%	time for request:	1.71ms
	10%	time for request:	1.91ms
	25%	time for request:	2.27ms
	50%	time for request:	2.76ms
	75%	time for request:	3.39ms
	90%	time for request:	4.20ms
	95%	time for request:	4.81ms
	99%	time for request:	6.45ms
	99.9%	time for request:	9.93ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	389900.38 Byte/s (0.39 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	389900.38 Byte/s (0.39 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_read_index -network=tcp -codec=pb -addr=:8001 -parallel=512 -total=1000000 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	512
	Total calls:	1000000
	Total time:	6.96s
	Requests per second:	143726.61
	Fastest time for request:	0.65ms
	Average time per request:	3.56ms
	Slowest time for request:	13.74ms

Time:
	0.1%	time for request:	1.28ms
	1%	time for request:	1.71ms
	5%	time for request:	2.14ms
	10%	time for request:	2.38ms
	25%	time for request:	2.82ms
	50%	time for request:	3.36ms
	75%	time for request:	4.08ms
	90%	time for request:	4.98ms
	95%	time for request:	5.64ms
	99%	time for request:	7.21ms
	99.9%	time for request:	9.46ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	431179.82 Byte/s (0.43 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	431179.82 Byte/s (0.43 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

## HSLAM Write

./rpc_write -network=tcp -codec=pb -addr=:8001 -parallel=256 -total=1000000 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	256
	Total calls:	1000000
	Total time:	101.91s
	Requests per second:	9812.32
	Fastest time for request:	7.96ms
	Average time per request:	26.08ms
	Slowest time for request:	64.19ms

Time:
	0.1%	time for request:	9.39ms
	1%	time for request:	11.51ms
	5%	time for request:	15.32ms
	10%	time for request:	17.19ms
	25%	time for request:	22.18ms
	50%	time for request:	26.64ms
	75%	time for request:	30.45ms
	90%	time for request:	33.45ms
	95%	time for request:	35.19ms
	99%	time for request:	39.22ms
	99.9%	time for request:	47.52ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	353243.60 Byte/s (0.35 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_write -network=tcp -codec=pb -addr=:8001 -parallel=512 -total=1000000 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	512
	Total calls:	1000000
	Total time:	55.55s
	Requests per second:	18001.91
	Fastest time for request:	10.21ms
	Average time per request:	28.43ms
	Slowest time for request:	235.59ms

Time:
	0.1%	time for request:	12.51ms
	1%	time for request:	14.08ms
	5%	time for request:	17.32ms
	10%	time for request:	19.22ms
	25%	time for request:	24.67ms
	50%	time for request:	28.52ms
	75%	time for request:	32.32ms
	90%	time for request:	35.11ms
	95%	time for request:	36.85ms
	99%	time for request:	47.06ms
	99.9%	time for request:	74.39ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	648068.75 Byte/s (0.65 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_write -network=tcp -codec=pb -addr=:8001 -parallel=768 -total=1000000 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	768
	Total calls:	1000000
	Total time:	36.32s
	Requests per second:	27534.67
	Fastest time for request:	10.50ms
	Average time per request:	27.88ms
	Slowest time for request:	93.19ms

Time:
	0.1%	time for request:	14.75ms
	1%	time for request:	16.97ms
	5%	time for request:	19.50ms
	10%	time for request:	21.44ms
	25%	time for request:	24.55ms
	50%	time for request:	26.97ms
	75%	time for request:	30.14ms
	90%	time for request:	34.52ms
	95%	time for request:	37.68ms
	99%	time for request:	53.71ms
	99.9%	time for request:	77.31ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	991248.27 Byte/s (0.99 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./rpc_write -network=tcp -codec=pb -addr=:8001 -parallel=1024 -total=1000000 -clients=1
```
Summary:
	Clients:	1
	Parallel calls per client:	1024
	Total calls:	1000000
	Total time:	29.52s
	Requests per second:	33874.20
	Fastest time for request:	14.32ms
	Average time per request:	30.21ms
	Slowest time for request:	89.80ms

Time:
	0.1%	time for request:	17.75ms
	1%	time for request:	20.49ms
	5%	time for request:	22.54ms
	10%	time for request:	23.84ms
	25%	time for request:	26.25ms
	50%	time for request:	28.78ms
	75%	time for request:	31.97ms
	90%	time for request:	37.13ms
	95%	time for request:	42.56ms
	99%	time for request:	62.46ms
	99.9%	time for request:	77.07ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	1219471.22 Byte/s (1.22 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
