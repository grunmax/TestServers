# TestServers
Test golang project with tcp and http servers.

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

and returns top words in json
```json
{"count":3,"top_words":["zulu","alfa","bravo"]}
```