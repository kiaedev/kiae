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
	externalEgress: [ for v in parameter.egress if v.external == true {
		v
	}]

	outputs: {
		if len(externalEgress) > 0 {
			kServiceEntry: {
				apiVersion: "networking.istio.io/v1beta1"
				kind:       "ServiceEntry"
				metadata: name: context.name
				spec: {
					resolution: "DNS"
					location:   "MESH_EXTERNAL"
					hosts: [ for v in externalEgress {
						"\(v.host)"
					}]
					ports: [ for v in externalEgress {
						{name: "\(v.protocol)-\(v.port)", number: v.port, protocol: v.protocol}
					}]
				}
			}
		}
	}
	patchOutputs: {
		kSideCar: {
			spec: {
				// +patchStrategy=replace
				egress: [
					{
						hosts: [ for v in parameter.egress {
							"./\(v.host)"
						}]
					},
				]
			}
		}
	}

	parameter: {

		egress: [...{

			host: string

			port: uint32

			protocol: string

			external: bool
		}]
	}
}
