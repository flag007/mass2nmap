# mass2nmap
别问 问就是不知道用来干啥的

用法:

```
▶ cat mass.out | grep open | mass2nmap > mass2nmap.out

▶ IFS=$'\n'; for i in $(cat mass2nmap.out); do cmd="$i"; eval $cmd;done; unset IFS

▶ cat *.gnmap | grep Port
```


安装:

```
▶ go get -u -v github.com/flag007/mass2nmap
```
