package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// s := "postgres://user:pass@host.com:5432/path?k=v#f"
	str := "http://localhost:8080/api/orders?page=1&limit=10"

	ustr, err := url.Parse(str)

	if err != nil {
		panic(err)
	}

	host, port, _ := net.SplitHostPort(ustr.Host)
	fmt.Println(host)
	fmt.Println(port)

	return
	/*
		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}

		fmt.Println(u.Scheme)

		fmt.Println(u.User)
		fmt.Println(u.User.Username())
		p, _ := u.User.Password()
		fmt.Println(p)

		fmt.Println(u.Host)
		host, port, _ := net.SplitHostPort(u.Host)
		fmt.Println(host)
		fmt.Println(port)

		fmt.Println(u.Path)
		fmt.Println(u.Fragment)

		fmt.Println(u.RawQuery)
		m, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(m)
		fmt.Println(m["k"][0]) */
}
