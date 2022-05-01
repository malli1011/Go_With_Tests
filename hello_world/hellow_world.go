package main

func Hello(args string) string {
	return "Hello, " + args
}

func Hello_v1() string {
	return "Hello, world"
}

func Hello_v2(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}

func Hello_v3(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return geetings(language) + ", " + name
}

func geetings(lang string) (prefix string) {
	switch lang {
	case "hungary":
		prefix = "Szia"
	case "spanish":
		prefix = "Hola"
	default:
		prefix = "Hello"
	}
	return
}
