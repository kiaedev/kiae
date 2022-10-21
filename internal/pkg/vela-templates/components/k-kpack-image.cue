"k-kpack-image": {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "kpack.io/v1alpha2"
		kind:       "Image"
	}
	description: "kpack image"
	labels: {

	}
	type: "component"
}

template: {
	uniName: "kiae-image-" + context.name
	output: {
		apiVersion: "kpack.io/v1alpha2"
		kind:       "Image"
		metadata: name: uniName
		spec: {
			serviceAccountName: uniName
			tag:                parameter.imageTag
			builder: {
				name: "kiae-builder-" + parameter.builderName
				kind: "Builder"
			}
			source: git: {
				url:      parameter.gitUrl
				revision: parameter.gitCommit
			}
		}
	}

	outputs: {
		serviceAccount: {
			apiVersion: "v1"
			kind:       "ServiceAccount"
			metadata: name: uniName
			secrets: [
				{name: parameter.gitRepoSecret},
				{name: parameter.imgRegSecret},
			]
		}

	}
	parameter: {

		builderName: string

		imageTag: string

		gitUrl: string

		gitCommit: string

		gitRepoSecret: string

		imgRegSecret: string
	}
}
