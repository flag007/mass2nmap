# mass2nmap
别问 问就是不知道用来干啥的

用法:

```
▶ cat mass.out | grep open | go run main.go >> mass2nmap.out

▶ IFS=$'\n'; for i in $(cat mass2nmap.out); do echo "$i"; done; unset IFS
```


安装:

```
▶ go get -u -v github.com/flag007/mass2nmap
```
