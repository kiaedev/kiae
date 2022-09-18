
"k-secret2envs": {
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

					// +patchKey=secretRef.name
					// +patchStrategy=retainKeys
					envFrom: [
						{secretRef: name: parameter.secretName},
					]
				},
			]
		}
	}

	parameter: {
		secretName: string
	}
}
