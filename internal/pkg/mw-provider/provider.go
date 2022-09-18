package mw_provider

import "sigs.k8s.io/controller-runtime/pkg/client"

type Constructor func(instanceName, name string) client.Object

var supportProvider = map[string]Constructor{
	"mysql": MySQLConfig,
	"mssql": MsSQLConfig,
	"pgsql": PgSQLConfig,
}

func BuildConfig(pType string, name, secretName string) client.Object {
	return supportProvider[pType](name, secretName)
}
