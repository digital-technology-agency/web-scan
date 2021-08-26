package generators

// Generator base generator interface
type Generator interface {
	Gen() <-chan string
}
