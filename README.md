# TestServers
Test golang project with tcp and http servers.

> This test app reads word list by tcp, sends data to non-blocking buffer in concurrent way. 
> Also it reads synchronously from buffer and increment value in storage map for each received word (key).
> Http server makes json for top N words received.
> All data store in memory only.

### tcp server
Config params in INI file:
 - host = localhost //
 - port = 9000 //
 - buffersize = 2048 //bytes read
 - minrunes = 3 //min word length in runes

Run netcat command:
```sh
$ echo "alfa zulu bravo ad whiskey Танго" | nc localhost 9000
```

TCP response:
```sh
tcp:ok:38:5
```
> tcp - constant part, ok/err - result (err - no words accepted), 38 - bytes received, 5 - words accepted count


### http server

Config params in INI file:
 - host = localhost //
 - port = 8000 //

http server has one test api:
[http://localhost:8000/top?N=3](http://localhost:8000/top?N=3)

and returns top received words in json
```json
{"count":3,"top_words":["zulu","alfa","bravo"]}
```

### data setting

Config params in INI file:
 - buffersize = 20 // "slots" in channel
 - debug = true // visualize passing data through channel (console only)