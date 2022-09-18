"k-secret2file": {
	alias: ""
	annotations: {}
	attributes: {
		podDisruptive: true
		appliesToWorkloads: ["*"]
	}
	description: ""
	labels: {}
	type: "trait"
}

template: {
	patch: {
		spec: template: spec: {
			// +patchKey=name
			containers: [
				{
					name: "kapp"

					// +patchKey=name
					volumeMounts: [
						{
							name:      parameter.secretName
							mountPath: "/kiae/mws/" + parameter.secretName
							readOnly:  true
						},
					]
				},
			]
			// +patchKey=name
			volumes: [
				{
					name: parameter.secretName
					secret: {
						secretName: parameter.secretName
					}
				},
			]
		}
	}

	parameter: {
		secretName: string
	}
}
