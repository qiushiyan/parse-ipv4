```
❯ go build -o main .
❯ ./main 172.168.5.1
172.168.5.1 -> 2896692481
❯ ./main "1 72.168.5.1"
invalid IPv4 address
```

```
❯ go test
PASS
ok      github.com/qiushiyan/parse-ipv4 0.004s
```
