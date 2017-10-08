This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)


Server Software:        nginx/1.12.1
Server Hostname:        127.0.0.1
Server Port:            4000

Document Path:          /httptest/160313.jpg
Document Length:        267037 bytes

Concurrency Level:      10
Time taken for tests:   0.900 seconds
Complete requests:      6000
Failed requests:        0
Total transferred:      1603656000 bytes
HTML transferred:       1602222000 bytes
Requests per second:    6667.59 [#/sec] (mean)
Time per request:       1.500 [ms] (mean)
Time per request:       0.150 [ms] (mean, across all concurrent requests)
Transfer rate:          1740319.84 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       3
Processing:     0    1   0.4      1       7
Waiting:        0    0   0.3      0       3
Total:          1    1   0.4      1       7

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      2
  80%      2
  90%      2
  95%      2
  98%      3
  99%      3
 100%      7 (longest request)
