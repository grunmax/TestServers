# TestServers
Test golang project with tcp and http servers.

> This app reads word list by tcp, sends data to non-blocking buffer in concurrent way. 
> Also reads synchronously from buffer and increment value in storage map for each received word (key).
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

http server has only one test api:
[http://localhost:8000/top?N=5](http://localhost:8000/top?N=5)

and returns top received words in json
```json
{"count":3,"top_words":["zulu","alfa","bravo"]}
```
> count = 3 is a real count value for top word list, 
> in case with equal count values sorts by name

### data setting

Config params in INI file:
 - buffersize = 20 // "slots" in buffered channel (received lists)
 - debug = true // visualize passing data in console