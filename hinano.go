package hinano

var (
	X    = 22
	name = "hinano"
)

type Yangcainai struct {
	Gender bool
	Age    uint
	Hobby  string
}

func NewYangcainai() Yangcainai {
	return Yangcainai{false, uint(X), name}
}
