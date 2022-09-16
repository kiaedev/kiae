"k-mw-mssqldb": {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "mssql.sql.crossplane.io/v1alpha1"
		kind:       "Database"
	}
	description: "kiae webservice"
	labels: {

	}
	type: "component"
}

template: {
	output: {
		apiVersion: "mssql.sql.crossplane.io/v1alpha1"
		kind:       "Database"
		metadata:
			name: example
		spec: {
			forProvider: {}
		}
	}
}
