package iteration

const repeatCount int = 5

func Repeat(chars string) (repeated string) {
	//var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + chars
	}
	return
}
