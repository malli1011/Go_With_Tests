package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

var websites = []string{
	"http://google.com",
	"http://blog.gypsydave5.com",
	"waat://furhurterwe.geds",
}

var want = map[string]bool{
	"http://google.com":          true,
	"http://blog.gypsydave5.com": true,
	"waat://furhurterwe.geds":    false,
}

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"

}

func TestCheckWebsites_v1(t *testing.T) {

	got := CheckWebsites_v1(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

//go test -race // to check the race condition
func TestCheckWebsites_v2(t *testing.T) {

	got := CheckWebsites_v2(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wnated %v, got %v", want, got)
	}
}

func TestCheckWebsites_v3(t *testing.T) {

	got := CheckWebsites_v3(slowStubWebsiteChecker, websites)
	for key, val := range got {
		fmt.Printf("key : %s, Value: %v\n", key, val)
	}

}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites_v1(slowStubWebsiteChecker, urls)
	}
}
