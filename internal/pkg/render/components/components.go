package components

import "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"

type Component interface {
	GetName() string
	GetType() string
	GetTraits() []common.ApplicationTrait
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
