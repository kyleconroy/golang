package main

import (
	"flag"
	"github.com/wsxiaoys/colors"
	"net/http"
)

func checkHandle(url string) bool {
	resp, err := http.Get(url)

	if err != nil {
		return false
	}

	return resp.StatusCode >= 400
}

func main() {
	services := []string{
		"http://www.twitter.com/",
		"http://www.youtube.com/user/",
		"http://soundcloud.com/",
		"https://www.facebook.com/",
		"http://www.reddit.com/user/",
	}

	flag.Parse()

	handle := flag.Arg(0)

	for _, value := range services {

		handleUrl := value + handle
		available := checkHandle(handleUrl)

		if available {
			colors.Printf("@g[AVAILABLE]\t@w%s\n", handleUrl)
		} else {
			colors.Printf("@r[RESERVED]\t@w%s\n", handleUrl)
		}
	}
}
