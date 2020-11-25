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

./etcd_read_index -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=128 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	128
	Total calls:	1000000
	Total time:	42.42s
	Requests per second:	23573.20
	Fastest time for request:	0.37ms
	Average time per request:	5.43ms
	Slowest time for request:	124.31ms

Time:
	0.1%	time for request:	0.72ms
	1%	time for request:	1.25ms
	5%	time for request:	1.94ms
	10%	time for request:	2.38ms
	25%	time for request:	3.21ms
	50%	time for request:	4.33ms
	75%	time for request:	6.39ms
	90%	time for request:	9.86ms
	95%	time for request:	12.97ms
	99%	time for request:	19.38ms
	99.9%	time for request:	28.33ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	70719.60 Byte/s (0.07 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	70719.60 Byte/s (0.07 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./etcd_read_index -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=256 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	256
	Total calls:	1000000
	Total time:	35.34s
	Requests per second:	28294.58
	Fastest time for request:	0.38ms
	Average time per request:	9.04ms
	Slowest time for request:	67.62ms

Time:
	0.1%	time for request:	0.84ms
	1%	time for request:	1.57ms
	5%	time for request:	2.65ms
	10%	time for request:	3.42ms
	25%	time for request:	4.96ms
	50%	time for request:	7.21ms
	75%	time for request:	11.35ms
	90%	time for request:	17.41ms
	95%	time for request:	21.66ms
	99%	time for request:	30.16ms
	99.9%	time for request:	40.37ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	84883.75 Byte/s (0.08 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	84883.75 Byte/s (0.08 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./etcd_read_index -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=384 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	384
	Total calls:	1000000
	Total time:	33.26s
	Requests per second:	30065.06
	Fastest time for request:	0.38ms
	Average time per request:	12.77ms
	Slowest time for request:	93.13ms

Time:
	0.1%	time for request:	0.95ms
	1%	time for request:	1.91ms
	5%	time for request:	3.35ms
	10%	time for request:	4.41ms
	25%	time for request:	6.74ms
	50%	time for request:	10.41ms
	75%	time for request:	16.57ms
	90%	time for request:	24.42ms
	95%	time for request:	29.90ms
	99%	time for request:	41.53ms
	99.9%	time for request:	56.96ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	90195.17 Byte/s (0.09 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	90195.17 Byte/s (0.09 MByte/s)

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
	Total time:	31.44s
	Requests per second:	31808.91
	Fastest time for request:	0.38ms
	Average time per request:	16.09ms
	Slowest time for request:	123.04ms

Time:
	0.1%	time for request:	1.15ms
	1%	time for request:	2.30ms
	5%	time for request:	4.03ms
	10%	time for request:	5.40ms
	25%	time for request:	8.50ms
	50%	time for request:	13.26ms
	75%	time for request:	21.09ms
	90%	time for request:	30.72ms
	95%	time for request:	37.76ms
	99%	time for request:	50.97ms
	99.9%	time for request:	63.96ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	95426.73 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	95426.73 Byte/s (0.10 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

## ETCD Write

./etcd_write -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=256 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	256
	Total calls:	1000000
	Total time:	323.01s
	Requests per second:	3095.88
	Fastest time for request:	15.37ms
	Average time per request:	82.68ms
	Slowest time for request:	356.02ms

Time:
	0.1%	time for request:	26.98ms
	1%	time for request:	34.81ms
	5%	time for request:	42.95ms
	10%	time for request:	48.21ms
	25%	time for request:	62.01ms
	50%	time for request:	84.36ms
	75%	time for request:	99.11ms
	90%	time for request:	114.11ms
	95%	time for request:	124.49ms
	99%	time for request:	142.39ms
	99.9%	time for request:	201.24ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	111451.86 Byte/s (0.11 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./etcd_write -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=512 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	512
	Total calls:	1000000
	Total time:	180.06s
	Requests per second:	5553.57
	Fastest time for request:	17.95ms
	Average time per request:	92.17ms
	Slowest time for request:	382.52ms

Time:
	0.1%	time for request:	29.21ms
	1%	time for request:	36.82ms
	5%	time for request:	46.22ms
	10%	time for request:	52.20ms
	25%	time for request:	68.00ms
	50%	time for request:	90.65ms
	75%	time for request:	109.19ms
	90%	time for request:	130.68ms
	95%	time for request:	144.67ms
	99%	time for request:	205.54ms
	99.9%	time for request:	306.49ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	199928.47 Byte/s (0.20 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./etcd_write -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=768 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	768
	Total calls:	1000000
	Total time:	128.18s
	Requests per second:	7801.52
	Fastest time for request:	18.46ms
	Average time per request:	98.40ms
	Slowest time for request:	458.64ms

Time:
	0.1%	time for request:	29.25ms
	1%	time for request:	36.87ms
	5%	time for request:	46.92ms
	10%	time for request:	53.65ms
	25%	time for request:	70.70ms
	50%	time for request:	94.10ms
	75%	time for request:	115.72ms
	90%	time for request:	140.59ms
	95%	time for request:	171.76ms
	99%	time for request:	245.68ms
	99.9%	time for request:	338.08ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	280854.88 Byte/s (0.28 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

./etcd_write -endpoints=127.0.0.1:2379,127.0.0.1:2479,127.0.0.1:2579 -conns=1 -clients=1024 -total=1000000
```
Summary:
	Clients:	1
	Parallel calls per client:	1024
	Total calls:	1000000
	Total time:	103.78s
	Requests per second:	9635.83
	Fastest time for request:	18.91ms
	Average time per request:	106.23ms
	Slowest time for request:	462.46ms

Time:
	0.1%	time for request:	28.06ms
	1%	time for request:	36.00ms
	5%	time for request:	46.84ms
	10%	time for request:	54.03ms
	25%	time for request:	72.16ms
	50%	time for request:	98.02ms
	75%	time for request:	123.80ms
	90%	time for request:	158.95ms
	95%	time for request:	213.05ms
	99%	time for request:	307.00ms
	99.9%	time for request:	404.16ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	346889.83 Byte/s (0.35 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```