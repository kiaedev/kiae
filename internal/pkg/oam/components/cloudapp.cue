cloudapp: {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "apps/v1"
		kind:       "Deployment"
	}
	description: ""
	labels: {

	}
	type: "component"
}

template: {
	output: {
		apiVersion: "apps/v1"
		kind:       "Deployment"
		spec: {
			template: {
				spec: {
					containers: [
						{
							image: parameter.image
						}
					]
				}
			}
		}
	}
	parameter: {
		image: string
		port: [...int]
		config: string
	}
}
