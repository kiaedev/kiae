

"k-mw-mysqldb": {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "mysql.sql.crossplane.io/v1alpha1"
		kind:       "Database"
	}
	description: "kiae webservice"
	labels: {

	}
	type: "component"
}

template: {
	output: {
		apiVersion: "mysql.sql.crossplane.io/v1alpha1"
		kind:       "Database"
		metadata:
			name: example
		spec: {
			forProvider: {}
		}
	}
}