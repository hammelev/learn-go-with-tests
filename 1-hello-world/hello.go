package helloWorld

const (
	spanish = "Spanish"
	french  = "French"
	danish  = "Danish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	danishHelloPrefix  = "Hej, "
)

func Hello(language string, name string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case danish:
		prefix = danishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
