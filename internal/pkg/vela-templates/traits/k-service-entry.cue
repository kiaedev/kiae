"k-service-entry": {
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
		kServiceEntry: {
			apiVersion: "networking.istio.io/v1beta1"
			kind:       "ServiceEntry"
			metadata: name: context.name
			spec: {
				resolution: "DNS"
				location:   "MESH_EXTERNAL"
				hosts:      parameter.hosts
				ports:      parameter.ports
			}
		}
	}

	parameter: {

		hosts: [...string]

		ports: [...{
			number: uint32

			protocol: string
		}]

	}
}
