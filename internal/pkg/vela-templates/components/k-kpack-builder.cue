"k-kpack-builder": {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "kpack.io/v1alpha2"
		kind:       "Builder"
	}
	description: "kpack builder with store and stack"
	labels: {

	}
	type: "component"
}

template: {
	uniName: "kiae-builder-" + context.name
	output: {
		apiVersion: "kpack.io/v1alpha2"
		kind:       "Builder"
		metadata: name: uniName
		spec: {
			serviceAccountName: uniName
			tag:                parameter.imageTag
			stack: {
				name: uniName
				kind: "ClusterStack"
			}
			store: {
				name: uniName
				kind: "ClusterStore"
			}
			order: [ for idx, pack in parameter.packs {
				{
					group: [{id: pack.id}]
				}
			}]
		}
	}

	outputs: {
		stack: {
			apiVersion: "kpack.io/v1alpha2"
			kind:       "ClusterStore"
			metadata: name: uniName
			spec: {
				sources: [ for idx, pack in parameter.packs {
					{
						image: pack.image
					}
				}]
			}
		}
		store: {
			apiVersion: "kpack.io/v1alpha2"
			kind:       "ClusterStack"
			metadata: name: uniName
			spec: {
				id: parameter.stackId
				buildImage: image: parameter.buildImage
				runImage: image:   parameter.runImage
			}
		}
		serviceAccount: {
			apiVersion: "v1"
			kind:       "ServiceAccount"
			metadata: name: uniName
			secrets: [{name: parameter.imageRegistry}]
		}
	}
	parameter: {

		imageTag: string

		imageRegistry: string

		stackId: string

		buildImage: string

		runImage: string

		packs: [...{
			id:    string
			image: string
		}]
	}
}
