package concurrency

import "time"

type WebsiteChecker func(string) bool

/*
Alongside the results map we now have a resultChannel, which we make in the same way.
chan result is the type of the channel - a channel of result.
The new type, result has been made to associate the return value of the WebsiteChecker
with the url being checked - it's a struct of string and bool.
As we don't need either value to be named, each of them is anonymous within the struct;
this can be useful in when it's hard to know what to name a value.
*/
type result struct {
	string
	bool
}

func CheckWebsites_v1(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}
	return results

}

//below function causes race conditon as multiple
//goroutines tries to update results map concurrently
func CheckWebsites_v2(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)

	}
	time.Sleep(2 * time.Second)
	return results

}

/*
Using channels to coordinate between various coroutines
*/
func CheckWebsites_v3(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
