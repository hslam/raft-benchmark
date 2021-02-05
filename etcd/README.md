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
	Total time:	37.54s
	Requests per second:	26635.33
	Fastest time for request:	0.40ms
	Average time per request:	9.61ms
	Slowest time for request:	78.45ms

Time:
	0.1%	time for request:	0.89ms
	1%	time for request:	1.63ms
	5%	time for request:	2.69ms
	10%	time for request:	3.42ms
	25%	time for request:	4.99ms
	50%	time for request:	7.44ms
	75%	time for request:	12.15ms
	90%	time for request:	19.08ms
	95%	time for request:	23.89ms
	99%	time for request:	33.54ms
	99.9%	time for request:	44.26ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	79905.99 Byte/s (0.08 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	79905.99 Byte/s (0.08 MByte/s)

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
	Total time:	33.12s
	Requests per second:	30191.85
	Fastest time for request:	0.42ms
	Average time per request:	16.95ms
	Slowest time for request:	106.38ms

Time:
	0.1%	time for request:	1.10ms
	1%	time for request:	2.23ms
	5%	time for request:	4.06ms
	10%	time for request:	5.45ms
	25%	time for request:	8.64ms
	50%	time for request:	13.77ms
	75%	time for request:	22.15ms
	90%	time for request:	32.91ms
	95%	time for request:	40.91ms
	99%	time for request:	56.35ms
	99.9%	time for request:	71.53ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	90575.54 Byte/s (0.09 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	90575.54 Byte/s (0.09 MByte/s)

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
	Total time:	31.27s
	Requests per second:	31983.85
	Fastest time for request:	0.44ms
	Average time per request:	24.00ms
	Slowest time for request:	128.02ms

Time:
	0.1%	time for request:	1.63ms
	1%	time for request:	3.37ms
	5%	time for request:	6.02ms
	10%	time for request:	8.14ms
	25%	time for request:	13.00ms
	50%	time for request:	20.62ms
	75%	time for request:	31.57ms
	90%	time for request:	44.63ms
	95%	time for request:	53.37ms
	99%	time for request:	71.80ms
	99.9%	time for request:	90.86ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	95951.55 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	95951.55 Byte/s (0.10 MByte/s)

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
	Total time:	30.70s
	Requests per second:	32570.79
	Fastest time for request:	0.48ms
	Average time per request:	31.42ms
	Slowest time for request:	160.86ms

Time:
	0.1%	time for request:	1.90ms
	1%	time for request:	4.26ms
	5%	time for request:	8.07ms
	10%	time for request:	11.01ms
	25%	time for request:	17.34ms
	50%	time for request:	27.41ms
	75%	time for request:	40.96ms
	90%	time for request:	57.42ms
	95%	time for request:	69.39ms
	99%	time for request:	92.74ms
	99.9%	time for request:	118.53ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	97712.38 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	97712.38 Byte/s (0.10 MByte/s)

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
	Total time:	167.77s
	Requests per second:	5960.41
	Fastest time for request:	19.94ms
	Average time per request:	85.87ms
	Slowest time for request:	316.50ms

Time:
	0.1%	time for request:	27.88ms
	1%	time for request:	35.09ms
	5%	time for request:	43.37ms
	10%	time for request:	49.18ms
	25%	time for request:	64.00ms
	50%	time for request:	85.88ms
	75%	time for request:	103.31ms
	90%	time for request:	120.41ms
	95%	time for request:	131.63ms
	99%	time for request:	168.50ms
	99.9%	time for request:	241.22ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	214574.88 Byte/s (0.21 MByte/s)

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
	Total time:	101.76s
	Requests per second:	9827.13
	Fastest time for request:	19.25ms
	Average time per request:	104.17ms
	Slowest time for request:	534.85ms

Time:
	0.1%	time for request:	27.54ms
	1%	time for request:	35.23ms
	5%	time for request:	44.80ms
	10%	time for request:	52.01ms
	25%	time for request:	69.57ms
	50%	time for request:	94.38ms
	75%	time for request:	122.31ms
	90%	time for request:	163.25ms
	95%	time for request:	209.60ms
	99%	time for request:	307.84ms
	99.9%	time for request:	401.95ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	353776.60 Byte/s (0.35 MByte/s)

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
	Total time:	84.28s
	Requests per second:	11865.04
	Fastest time for request:	18.09ms
	Average time per request:	129.39ms
	Slowest time for request:	677.91ms

Time:
	0.1%	time for request:	28.11ms
	1%	time for request:	36.24ms
	5%	time for request:	47.89ms
	10%	time for request:	57.04ms
	25%	time for request:	77.51ms
	50%	time for request:	107.56ms
	75%	time for request:	149.30ms
	90%	time for request:	245.07ms
	95%	time for request:	312.55ms
	99%	time for request:	402.38ms
	99.9%	time for request:	554.76ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	427141.30 Byte/s (0.43 MByte/s)

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
	Total time:	74.88s
	Requests per second:	13354.91
	Fastest time for request:	20.48ms
	Average time per request:	153.10ms
	Slowest time for request:	744.96ms

Time:
	0.1%	time for request:	28.81ms
	1%	time for request:	37.03ms
	5%	time for request:	51.87ms
	10%	time for request:	62.60ms
	25%	time for request:	86.58ms
	50%	time for request:	122.11ms
	75%	time for request:	182.07ms
	90%	time for request:	307.62ms
	95%	time for request:	378.25ms
	99%	time for request:	484.61ms
	99.9%	time for request:	611.24ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	480776.65 Byte/s (0.48 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```