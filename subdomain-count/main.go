package main

import (
	"fmt"
	"strings"
	"strconv"
	"bytes"
)

func subdomainVisits(cpdomains []string) []string {
	var visits []string
	subDomains := make(map[string]int, 0)

	for _, domain := range cpdomains {
		domainSplit := strings.Split(domain, " ")
		countStr, sdlist := domainSplit[0], strings.Split(domainSplit[1], ".")
		count, _ := strconv.Atoi(countStr)

		insertDomain(subDomains, count, 0, sdlist)
	}

	for k, v := range subDomains {
		strV := strconv.Itoa(v)
		visits = append(visits, strV+" "+k)
	}

	return visits
}

func insertDomain(subDomains map[string]int, count, start int, list []string) {
	if start == len(list) {
		return
	}

	var buf bytes.Buffer

	for i := start; i <= len(list)-1; i++ {
		buf.WriteString(list[i])

		if i+1 <= len(list)-1 {
			buf.WriteString(".")
		}
	}

	domainName := buf.String()
	subDomains[domainName] += count
	
	insertDomain(subDomains, count, start+1, list)
}

func main() {
	foo := subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"})

	for _, s := range foo {
		fmt.Println(s)
	}
}

