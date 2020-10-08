# mass2nmap
别问 问就是不知道用来干啥的

用法:

```
cat mass.out | grep open | go run main.go >> 
```

```
▶ IFS=$'\n'; for i in $(cat masscan2nmap.out); do echo "$i"; done; unset IFS
```




安装:

```
▶ go get -u -v github.com/flag007/masscan-parser
```
