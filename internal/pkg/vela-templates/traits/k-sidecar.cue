"k-sidecar": {
	alias: ""
	annotations: {}
	attributes: {
		appliesToWorkloads: ["*"]
	}
	description: ""
	labels: {}
	type: "trait"
}

template: {
	patchOutputs: {
		kSideCar: {
			spec: {
				// +patchStrategy=replace
				egress: [ for _, v in parameter.egress {
					{
						port:  v.port
						hosts: v.hosts
					}
				}]
			}
		}
	}

	parameter: {

		egress: [...{
			port: {
				number: uint32

				protocol: string
			}

			hosts: [...string]
		}]
	}
}
