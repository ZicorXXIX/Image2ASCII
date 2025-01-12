package ascii

type Options struct {
	Pixels   []byte
	Reversed bool
	Colored  bool
}

// DefaultOptions that contains the default pixels
var DefaultOptions = Options{
	Pixels:   []byte(" .,:;i1tfLCG08@"),
	Reversed: false,
	Colored:  true,
}
