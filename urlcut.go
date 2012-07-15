package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		link, err := url.Parse(strings.TrimSpace(line))

		if err != nil {
			break
		}

		fmt.Print(link.Scheme)
		fmt.Print("\t")
		fmt.Print(link.Host)
		fmt.Print("\t")
		fmt.Print(link.Path)
		fmt.Print("\t")
		fmt.Print(link.RawQuery)
		fmt.Print("\t")
		fmt.Print(link.Fragment)
		fmt.Println()
	}
}
