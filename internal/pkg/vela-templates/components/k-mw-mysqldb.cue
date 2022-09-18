"k-mw-mysqldb": {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "mysql.sql.crossplane.io/v1alpha1"
		kind:       "Database"
	}
	description: "claim a mysql database and get connection info into the secret"
	labels: {

	}
	type: "component"
}

template: {
	output: {
		apiVersion: "mysql.sql.crossplane.io/v1alpha1"
		kind:       "Database"
		metadata:
			name: parameter.dbname
		spec: {
			providerConfigRef: {
				name: parameter.instance
			}
		}
	}

	outputs: {
		user: {
			apiVersion: "mysql.sql.crossplane.io/v1alpha1"
			kind:       "User"
			metadata: {
				name: parameter.dbname + "-rw"
			}
			spec: {
				providerConfigRef: {
					name: parameter.instance
				}
				forProvider: {
					resourceOptions: {
						maxQueriesPerHour:     1000
						maxUpdatesPerHour:     1000
						maxConnectionsPerHour: 100
						maxUserConnections:    10
					}
				}
				writeConnectionSecretToRef: {
					name:      parameter.dbname
					namespace: context.namespace
				}
			}
		}
		grant: {
			apiVersion: "mysql.sql.crossplane.io/v1alpha1"
			kind:       "Grant"
			metadata: {
				name: parameter.dbname + "-rw"
			}
			spec: {
				providerConfigRef: {
					name: parameter.instance
				}
				forProvider: {
					privileges: ["SELECT", "INSERT", "UPDATE", "DELETE"]
					userRef: name:     parameter.dbname + "-rw"
					databaseRef: name: parameter.dbname
				}
			}
		}
	}

	parameter: {
		instance: string

		dbname: string
	}
}
