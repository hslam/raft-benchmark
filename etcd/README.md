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
	Total time:	37.65s
	Requests per second:	26558.23
	Fastest time for request:	0.40ms
	Average time per request:	9.64ms
	Slowest time for request:	76.93ms

Time:
	0.1%	time for request:	0.87ms
	1%	time for request:	1.62ms
	5%	time for request:	2.70ms
	10%	time for request:	3.45ms
	25%	time for request:	5.03ms
	50%	time for request:	7.46ms
	75%	time for request:	12.15ms
	90%	time for request:	19.05ms
	95%	time for request:	23.84ms
	99%	time for request:	33.87ms
	99.9%	time for request:	43.89ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	79674.69 Byte/s (0.08 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	79674.69 Byte/s (0.08 MByte/s)

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
	Total time:	32.64s
	Requests per second:	30637.96
	Fastest time for request:	0.43ms
	Average time per request:	16.70ms
	Slowest time for request:	102.99ms

Time:
	0.1%	time for request:	1.08ms
	1%	time for request:	2.21ms
	5%	time for request:	3.89ms
	10%	time for request:	5.17ms
	25%	time for request:	8.28ms
	50%	time for request:	13.37ms
	75%	time for request:	22.08ms
	90%	time for request:	33.10ms
	95%	time for request:	40.43ms
	99%	time for request:	55.49ms
	99.9%	time for request:	72.31ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	91913.88 Byte/s (0.09 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	91913.88 Byte/s (0.09 MByte/s)

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
	Total time:	30.82s
	Requests per second:	32451.43
	Fastest time for request:	0.42ms
	Average time per request:	23.65ms
	Slowest time for request:	163.60ms

Time:
	0.1%	time for request:	1.23ms
	1%	time for request:	2.78ms
	5%	time for request:	5.42ms
	10%	time for request:	7.47ms
	25%	time for request:	12.19ms
	50%	time for request:	19.67ms
	75%	time for request:	31.11ms
	90%	time for request:	45.19ms
	95%	time for request:	55.83ms
	99%	time for request:	76.86ms
	99.9%	time for request:	99.42ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	97354.30 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	97354.30 Byte/s (0.10 MByte/s)

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
	Total time:	29.69s
	Requests per second:	33686.77
	Fastest time for request:	0.44ms
	Average time per request:	30.37ms
	Slowest time for request:	135.09ms

Time:
	0.1%	time for request:	2.08ms
	1%	time for request:	4.53ms
	5%	time for request:	8.22ms
	10%	time for request:	11.10ms
	25%	time for request:	17.13ms
	50%	time for request:	26.80ms
	75%	time for request:	39.62ms
	90%	time for request:	54.61ms
	95%	time for request:	65.42ms
	99%	time for request:	86.05ms
	99.9%	time for request:	110.39ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	101060.31 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	101060.31 Byte/s (0.10 MByte/s)

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
	Total time:	166.85s
	Requests per second:	5993.35
	Fastest time for request:	19.64ms
	Average time per request:	85.41ms
	Slowest time for request:	337.90ms

Time:
	0.1%	time for request:	28.42ms
	1%	time for request:	34.99ms
	5%	time for request:	43.06ms
	10%	time for request:	48.81ms
	25%	time for request:	63.88ms
	50%	time for request:	85.86ms
	75%	time for request:	102.62ms
	90%	time for request:	119.45ms
	95%	time for request:	130.97ms
	99%	time for request:	161.70ms
	99.9%	time for request:	223.13ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	215760.57 Byte/s (0.22 MByte/s)

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
	Total time:	102.67s
	Requests per second:	9739.60
	Fastest time for request:	18.59ms
	Average time per request:	105.10ms
	Slowest time for request:	493.66ms

Time:
	0.1%	time for request:	27.89ms
	1%	time for request:	35.27ms
	5%	time for request:	45.43ms
	10%	time for request:	52.64ms
	25%	time for request:	70.70ms
	50%	time for request:	96.20ms
	75%	time for request:	123.05ms
	90%	time for request:	163.45ms
	95%	time for request:	213.12ms
	99%	time for request:	295.47ms
	99.9%	time for request:	412.51ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	350625.51 Byte/s (0.35 MByte/s)

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
	Total time:	83.94s
	Requests per second:	11913.14
	Fastest time for request:	20.48ms
	Average time per request:	128.84ms
	Slowest time for request:	671.13ms

Time:
	0.1%	time for request:	28.41ms
	1%	time for request:	36.74ms
	5%	time for request:	48.64ms
	10%	time for request:	57.93ms
	25%	time for request:	78.32ms
	50%	time for request:	107.71ms
	75%	time for request:	145.30ms
	90%	time for request:	238.24ms
	95%	time for request:	315.56ms
	99%	time for request:	416.15ms
	99.9%	time for request:	534.13ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	428872.94 Byte/s (0.43 MByte/s)

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
	Total time:	74.80s
	Requests per second:	13369.63
	Fastest time for request:	19.26ms
	Average time per request:	153.06ms
	Slowest time for request:	702.36ms

Time:
	0.1%	time for request:	30.94ms
	1%	time for request:	41.64ms
	5%	time for request:	57.74ms
	10%	time for request:	70.03ms
	25%	time for request:	95.56ms
	50%	time for request:	127.42ms
	75%	time for request:	172.06ms
	90%	time for request:	272.95ms
	95%	time for request:	377.27ms
	99%	time for request:	502.02ms
	99.9%	time for request:	609.90ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	481306.63 Byte/s (0.48 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```