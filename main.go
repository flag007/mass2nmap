package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	portsMap := make(map[string][]string)

	for sc.Scan() {
		line := strings.ToLower(sc.Text())
		s := strings.Fields(line)
		ip := s[1]
		portNumber := strings.Split(s[4], "/")[0]

		portsMap[ip] = append(portsMap[ip], portNumber)


//		fmt.Println(ip)
//		fmt.Println(portNumber)
//		fmt.Println(line)
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}


	for ip, ports := range portsMap {
		out := "nmap -sV " + ip + " -p" 
		for _, port := range ports {
			out = out + port + ","
		}
		out = strings.TrimRight(out, ",")
		out = out + " -oA " + ip
		fmt.Println(out)
	}

}
