package main

import (
	"bytes"
	"strings"
)

// Travel - returns a list of addresses in the given zipcode
func Travel(r, zipcode string) string {
	var result bytes.Buffer
	addrs := getAddressesWithZip(r, zipcode)

	result.WriteString(zipcode)
	result.WriteString(":")

	if len(addrs["houseNums"]) > 0 && len(addrs["streets"]) > 0 {
		result.WriteString(strings.Join(addrs["streets"], ","))
		result.WriteString("/")
		result.WriteString(strings.Join(addrs["houseNums"], ","))
	} else {
		result.WriteString("/")
	}

	return result.String()
}

func getAddressesWithZip(r, zipcode string) map[string][]string {
	addrs := make(map[string][]string)
	normalizedAddrs := strings.ReplaceAll(r, "\n", "")
	addrSplit := strings.Split(normalizedAddrs, ",")

	for _, str := range addrSplit {
		normalizedAddr := strings.Trim(str, " ")
		addr := strings.Split(normalizedAddr, " ")
		houseNum := addr[0]
		street := strings.Join(addr[1:len(addr)-2], " ")
		addrZip := strings.Join(addr[len(addr)-2:], " ")

		if strings.Compare(zipcode, addrZip) == 0 {
			if _, ok := addrs["houseNums"]; ok {
				addrs["houseNums"] = append(addrs["houseNums"], houseNum)
			} else {
				addrs["houseNums"] = []string{houseNum}
			}

			if _, ok := addrs["streets"]; ok {
				addrs["streets"] = append(addrs["streets"], street)
			} else {
				addrs["streets"] = []string{street}
			}
		}
	}

	return addrs
}

// TravelSimple - cleaner solution for the problem...
func TravelSimple(r, zipcode string) string {
	var strInfo []string
	var strNumbers []string

	for _, str := range strings.Split(r, ",") {
		tmp := strings.Fields(str) // need to remember this method
		if zipcode == strings.Join(tmp[(len(tmp)-2):], " ") {
			strNumbers = append(strNumbers, tmp[0])
			strInfo = append(strInfo, strings.Join(tmp[1:len(tmp)-2], " "))
		}
	}

	return zipcode + ":" + strings.Join(strInfo, ",") + "/" + strings.Join(strNumbers, ",")
}
