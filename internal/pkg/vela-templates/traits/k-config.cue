import (
	"strconv"
)

"k-config": {
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
					volumeMounts: [ for idx, v in parameter.configs {
						{
							name:      "kcfg-" + context.name + "-" + strconv.FormatInt(idx+1, 10)
							mountPath: v.mountPath
						}
					}]
				},
			]
			// +patchKey=name
			volumes: [ for idx, v in parameter.configs {
				{
					name: "kcfg-" + context.name + "-" + strconv.FormatInt(idx+1, 10)
					configMap: {
						name: context.name + "-" + v.filename
					}
				}
			}]
		}
	}
	outputs: {
		for v in parameter.configs {
			"\(v.filename)": {
				apiVersion: "v1"
				kind:       "ConfigMap"
				metadata: name: context.name + "-" + v.filename
				data: {
					"\(v.filename)": v.content
				}
			}
		}
	}
	parameter: {

		configs?: [...{
			filename: string

			content: string

			mountPath: string
		}]

	}
}
