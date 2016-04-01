# TestServers
Test golang project with tcp and http servers.

### tcp server
Config params in INI file:
 - host = localhost //
 - port = 9000 //
 - buffersize = 2048 //bytes read
 - minrunes = 3 //word consist of 3 chars at least

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