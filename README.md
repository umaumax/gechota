# gechota

This command is named after 'echo' and '[ゲコ太]( http://dic.nicovideo.jp/a/%E3%82%B2%E3%82%B3%E5%A4%AA )'

tcp echo server

## how to install
```
go get -u github.com/umaumax/gechota
```

## how to run
```
$ gechota -p=3939
$ curl -x localhost:3939 -L https://www.google.com
CONNECT www.google.com:443 HTTP/1.1
Host: www.google.com:443
User-Agent: curl/7.54.0
Proxy-Connection: Keep-Alive
```

FYI
```
curl http://localhost:3939 --max-time 5
curl -F "name1=misaka1" -F "name2=mikoto2" http://localhost:3939/ --max-time 3
```
