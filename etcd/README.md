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
	Total time:	36.19s
	Requests per second:	27634.33
	Fastest time for request:	0.38ms
	Average time per request:	9.26ms
	Slowest time for request:	111.51ms

Time:
	0.1%	time for request:	0.87ms
	1%	time for request:	1.62ms
	5%	time for request:	2.71ms
	10%	time for request:	3.47ms
	25%	time for request:	5.03ms
	50%	time for request:	7.33ms
	75%	time for request:	11.72ms
	90%	time for request:	17.93ms
	95%	time for request:	22.25ms
	99%	time for request:	30.85ms
	99.9%	time for request:	41.42ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	82903.00 Byte/s (0.08 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	82903.00 Byte/s (0.08 MByte/s)

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
	Total time:	32.72s
	Requests per second:	30562.49
	Fastest time for request:	0.38ms
	Average time per request:	16.74ms
	Slowest time for request:	108.62ms

Time:
	0.1%	time for request:	1.03ms
	1%	time for request:	2.25ms
	5%	time for request:	4.06ms
	10%	time for request:	5.44ms
	25%	time for request:	8.68ms
	50%	time for request:	13.77ms
	75%	time for request:	21.82ms
	90%	time for request:	32.29ms
	95%	time for request:	39.33ms
	99%	time for request:	54.66ms
	99.9%	time for request:	74.94ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	91687.48 Byte/s (0.09 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	91687.48 Byte/s (0.09 MByte/s)

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
	Requests per second:	32596.63
	Fastest time for request:	0.38ms
	Average time per request:	23.55ms
	Slowest time for request:	123.34ms

Time:
	0.1%	time for request:	1.46ms
	1%	time for request:	3.23ms
	5%	time for request:	6.03ms
	10%	time for request:	8.16ms
	25%	time for request:	13.18ms
	50%	time for request:	20.45ms
	75%	time for request:	30.71ms
	90%	time for request:	43.09ms
	95%	time for request:	51.99ms
	99%	time for request:	69.56ms
	99.9%	time for request:	89.51ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	97789.90 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	97789.90 Byte/s (0.10 MByte/s)

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
	Total time:	30.17s
	Requests per second:	33140.83
	Fastest time for request:	0.39ms
	Average time per request:	30.88ms
	Slowest time for request:	163.34ms

Time:
	0.1%	time for request:	1.68ms
	1%	time for request:	4.11ms
	5%	time for request:	8.03ms
	10%	time for request:	11.04ms
	25%	time for request:	17.39ms
	50%	time for request:	27.29ms
	75%	time for request:	40.32ms
	90%	time for request:	55.30ms
	95%	time for request:	66.36ms
	99%	time for request:	91.03ms
	99.9%	time for request:	118.86ms

Request:
	Total request body sizes:	3000000
	Average body size per request:	3.00 Byte
	Request rate per second:	99422.48 Byte/s (0.10 MByte/s)

Response:
	Total response body sizes:	3000000
	Average body size per response:	3.00 Byte
	Response rate per second:	99422.48 Byte/s (0.10 MByte/s)

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
	Total time:	173.66s
	Requests per second:	5758.29
	Fastest time for request:	17.72ms
	Average time per request:	88.89ms
	Slowest time for request:	694.09ms

Time:
	0.1%	time for request:	28.71ms
	1%	time for request:	36.56ms
	5%	time for request:	45.53ms
	10%	time for request:	51.15ms
	25%	time for request:	65.82ms
	50%	time for request:	88.01ms
	75%	time for request:	104.39ms
	90%	time for request:	122.77ms
	95%	time for request:	135.21ms
	99%	time for request:	190.28ms
	99.9%	time for request:	463.04ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	207298.56 Byte/s (0.21 MByte/s)

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
	Total time:	103.26s
	Requests per second:	9683.86
	Fastest time for request:	16.62ms
	Average time per request:	105.70ms
	Slowest time for request:	494.16ms

Time:
	0.1%	time for request:	28.95ms
	1%	time for request:	36.54ms
	5%	time for request:	46.86ms
	10%	time for request:	54.51ms
	25%	time for request:	73.86ms
	50%	time for request:	97.99ms
	75%	time for request:	122.43ms
	90%	time for request:	157.65ms
	95%	time for request:	208.15ms
	99%	time for request:	299.68ms
	99.9%	time for request:	401.25ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	348618.90 Byte/s (0.35 MByte/s)

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
	Total time:	82.27s
	Requests per second:	12155.49
	Fastest time for request:	18.72ms
	Average time per request:	126.27ms
	Slowest time for request:	595.09ms

Time:
	0.1%	time for request:	28.65ms
	1%	time for request:	37.25ms
	5%	time for request:	49.77ms
	10%	time for request:	59.08ms
	25%	time for request:	79.39ms
	50%	time for request:	108.59ms
	75%	time for request:	145.39ms
	90%	time for request:	227.81ms
	95%	time for request:	290.93ms
	99%	time for request:	372.25ms
	99.9%	time for request:	440.46ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	437597.58 Byte/s (0.44 MByte/s)

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
	Total time:	72.52s
	Requests per second:	13789.60
	Fastest time for request:	22.29ms
	Average time per request:	148.32ms
	Slowest time for request:	609.40ms

Time:
	0.1%	time for request:	30.24ms
	1%	time for request:	38.51ms
	5%	time for request:	52.16ms
	10%	time for request:	63.35ms
	25%	time for request:	86.64ms
	50%	time for request:	121.75ms
	75%	time for request:	171.56ms
	90%	time for request:	295.38ms
	95%	time for request:	366.04ms
	99%	time for request:	444.25ms
	99.9%	time for request:	535.01ms

Request:
	Total request body sizes:	36000000
	Average body size per request:	36.00 Byte
	Request rate per second:	496425.60 Byte/s (0.50 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```