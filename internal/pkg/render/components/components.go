package components

type Component interface {
	GetName() string
	GetType() string
}

var _ = KWebservice{}
var _ = KCronjob{}

type ComponentConstructor func(instance, name string) Component

var supportMiddlewares = map[string]ComponentConstructor{
	"mysql": NewMwMySQLDb,
	// "mssql": NewMwMySQLDb,
	// "pgsql": NewMwMySQLDb,
}

func MwConstructor(mType string, instance, name string) Component {
	return supportMiddlewares[mType](instance, name)
}
