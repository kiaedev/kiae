package traits

type Trait interface {
	GetName() string
	GetType() string
}

var (
	_ = &RouteTrait{}
)
