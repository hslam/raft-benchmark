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
	Total time:	36.17s
	Requests per second:	27649.51
	Fastest time for request:	0.38ms
	Average time per request:	9.25ms
	Slowest time for request:	109.10ms

Time:
	0.1%	time for request:	0.86ms
	1%	time for request:	1.61ms
	5%	time for request:	2.72ms
	10%	time for request:	3.48ms
	25%	time for request:	5.07ms
	50%	time for request:	7.37ms
	75%	time for request:	11.52ms
	90%	time for request:	17.85ms
	95%	time for request:	22.30ms
	99%	time for request:	31.41ms
	99.9%	time for request:	43.31ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	82948.54 Byte/s (0.08 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	82948.54 Byte/s (0.08 MByte/s)

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
	Total time:	32.06s
	Requests per second:	31190.33
	Fastest time for request:	0.39ms
	Average time per request:	16.41ms
	Slowest time for request:	111.89ms

Time:
	0.1%	time for request:	1.08ms
	1%	time for request:	2.31ms
	5%	time for request:	4.20ms
	10%	time for request:	5.58ms
	25%	time for request:	8.70ms
	50%	time for request:	13.48ms
	75%	time for request:	21.63ms
	90%	time for request:	31.32ms
	95%	time for request:	38.02ms
	99%	time for request:	51.25ms
	99.9%	time for request:	65.97ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	93571.00 Byte/s (0.09 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	93571.00 Byte/s (0.09 MByte/s)

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
	Total time:	30.68s
	Requests per second:	32597.25
	Fastest time for request:	0.39ms
	Average time per request:	23.54ms
	Slowest time for request:	118.23ms

Time:
	0.1%	time for request:	1.45ms
	1%	time for request:	3.39ms
	5%	time for request:	6.05ms
	10%	time for request:	8.12ms
	25%	time for request:	13.19ms
	50%	time for request:	20.48ms
	75%	time for request:	30.95ms
	90%	time for request:	42.84ms
	95%	time for request:	51.05ms
	99%	time for request:	68.81ms
	99.9%	time for request:	90.95ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	97791.76 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	97791.76 Byte/s (0.10 MByte/s)

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
	Total time:	30.27s
	Requests per second:	33038.68
	Fastest time for request:	0.38ms
	Average time per request:	30.96ms
	Slowest time for request:	191.42ms

Time:
	0.1%	time for request:	1.73ms
	1%	time for request:	4.35ms
	5%	time for request:	8.22ms
	10%	time for request:	11.18ms
	25%	time for request:	17.40ms
	50%	time for request:	27.32ms
	75%	time for request:	40.51ms
	90%	time for request:	55.68ms
	95%	time for request:	65.92ms
	99%	time for request:	87.83ms
	99.9%	time for request:	120.87ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	99116.03 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	99116.03 Byte/s (0.10 MByte/s)

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
	Total time:	172.27s
	Requests per second:	5804.73
	Fastest time for request:	20.95ms
	Average time per request:	88.19ms
	Slowest time for request:	324.81ms

Time:
	0.1%	time for request:	28.66ms
	1%	time for request:	36.71ms
	5%	time for request:	45.66ms
	10%	time for request:	51.35ms
	25%	time for request:	66.07ms
	50%	time for request:	88.71ms
	75%	time for request:	105.34ms
	90%	time for request:	123.53ms
	95%	time for request:	133.91ms
	99%	time for request:	162.85ms
	99.9%	time for request:	253.16ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	208970.23 Byte/s (0.21 MByte/s)

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
	Total time:	104.44s
	Requests per second:	9574.90
	Fastest time for request:	19.31ms
	Average time per request:	106.90ms
	Slowest time for request:	460.86ms

Time:
	0.1%	time for request:	27.93ms
	1%	time for request:	36.35ms
	5%	time for request:	47.00ms
	10%	time for request:	54.82ms
	25%	time for request:	73.03ms
	50%	time for request:	98.50ms
	75%	time for request:	124.64ms
	90%	time for request:	164.17ms
	95%	time for request:	211.72ms
	99%	time for request:	299.14ms
	99.9%	time for request:	389.40ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	344696.31 Byte/s (0.34 MByte/s)

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
	Total time:	83.72s
	Requests per second:	11944.59
	Fastest time for request:	18.69ms
	Average time per request:	128.46ms
	Slowest time for request:	619.84ms

Time:
	0.1%	time for request:	30.08ms
	1%	time for request:	37.96ms
	5%	time for request:	49.94ms
	10%	time for request:	58.64ms
	25%	time for request:	80.22ms
	50%	time for request:	110.36ms
	75%	time for request:	147.55ms
	90%	time for request:	233.27ms
	95%	time for request:	297.40ms
	99%	time for request:	380.96ms
	99.9%	time for request:	459.92ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	430005.10 Byte/s (0.43 MByte/s)

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
	Total time:	74.67s
	Requests per second:	13392.03
	Fastest time for request:	18.44ms
	Average time per request:	152.68ms
	Slowest time for request:	752.31ms

Time:
	0.1%	time for request:	29.16ms
	1%	time for request:	39.56ms
	5%	time for request:	54.40ms
	10%	time for request:	65.36ms
	25%	time for request:	91.35ms
	50%	time for request:	126.27ms
	75%	time for request:	178.34ms
	90%	time for request:	290.48ms
	95%	time for request:	367.90ms
	99%	time for request:	476.10ms
	99.9%	time for request:	577.18ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	482113.21 Byte/s (0.48 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```