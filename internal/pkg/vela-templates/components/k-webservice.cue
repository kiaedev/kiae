import (
	"strconv"
)

"k-webservice": {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "apps/v1"
		kind:       "Deployment"
	}
	description: "kiae webservice"
	labels: {

	}
	type: "component"
}

template: {
	output: {
		apiVersion: "apps/v1"
		kind:       "Deployment"
		spec: {
			replicas: parameter.replicas
			selector: {
				matchLabels: {
					"kiae.dev/component": context.name
				}
			}
			//   restartAt: parameter.restartAt
			template: {
				metadata: {
					labels: {
						"kiae.dev/component":   context.name
						"kiae.dev/revision":    context.revision
						"kiae.dev/appRevision": context.appRevision
					}
					if parameter.annotations != _|_ {
						annotations: parameter.annotations
					}
				}
				spec: {
					serviceAccountName: context.name
					containers: [
						{
							name:      "kapp"
							image:     parameter.image
							resources: parameter.resources
							if parameter["ports"] != _|_ {
								ports: [ for v in parameter.ports {
									{
										name:          v.appProtocol + "-" + strconv.FormatInt(v.port, 10)
										containerPort: v.port
										protocol:      v.protocol
										appProtocol:   v.appProtocol
									}}]
							}

							if parameter["livenessProbe"] != _|_ {
								livenessProbe: parameter.livenessProbe
							}

							if parameter["readinessProbe"] != _|_ {
								readinessProbe: parameter.readinessProbe
							}
						},
					]
				}
			}
		}
	}
	outputs: {
		kAppAccount: {
			apiVersion: "v1"
			kind:       "ServiceAccount"
			metadata: name: context.name
		}
		kAppService: {
			apiVersion: "v1"
			kind:       "Service"
			metadata: name: context.name
			spec: {
				selector: "kiae.dev/component": context.name
				ports: parameter.ports
			}
		}
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

		// +usage=Specify the name in the workload
		name: string

		// +usage=Specify the labels in the workload
		labels?: [string]: string

		// +usage=Specify the annotations in the workload
		annotations?: [string]: string

		// +usage=Which image would you like to use for your service
		// +short=i
		image: string

		// +usage=Specify image pull policy for your service
		imagePullPolicy?: "Always" | "Never" | "IfNotPresent"

		// +usage=Specify image pull secrets for your service
		imagePullSecrets?: [...string]

		// +usage=Which ports do you want customer traffic sent to, defaults to 80
		ports: [...{
			// +usage=Number of port to expose on the pod's IP address
			port: int
			// +usage=Protocol for port. Must be UDP, TCP, or SCTP
			protocol: *"TCP" | "UDP" | "SCTP"
			// +usage=Application Protocol of the port
			appProtocol: string
		}]

		replicas: *1 | int

		// +usage=Specify the resources in requests
		requests?: {
			// +usage=Specify the amount of cpu for requests
			cpu: *1 | number
			// +usage=Specify the amount of memory for requests
			memory: *"2048Mi" | =~"^([1-9][0-9]{0,63})(E|P|T|G|M|K|Ei|Pi|Ti|Gi|Mi|Ki)$"
		}
		// +usage=Specify the resources in limits
		limits?: {
			// +usage=Specify the amount of cpu for limits
			cpu: *1 | number
			// +usage=Specify the amount of memory for limits
			memory: *"2048Mi" | =~"^([1-9][0-9]{0,63})(E|P|T|G|M|K|Ei|Pi|Ti|Gi|Mi|Ki)$"
		}

		// +usage=Define arguments by using environment variables
		envs?: [string]: string

		configs?: [...{
			filename: string

			content: string

			mountPath: string
		}]

		// +usage=Instructions for assessing whether the container is alive.
		livenessProbe?: #HealthProbe

		// +usage=Instructions for assessing whether the container is in a suitable state to serve traffic.
		readinessProbe?: #HealthProbe
	}

	#HealthProbe: {

		// +usage=Instructions for assessing container health by executing a command. Either this attribute or the httpGet attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the httpGet attribute and the tcpSocket attribute.
		exec?: {
			// +usage=A command to be executed inside the container to assess its health. Each space delimited token of the command is a separate array element. Commands exiting 0 are considered to be successful probes, whilst all other exit codes are considered failures.
			command: [...string]
		}

		// +usage=Instructions for assessing container health by executing an HTTP GET request. Either this attribute or the exec attribute or the tcpSocket attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the tcpSocket attribute.
		httpGet?: {
			// +usage=The endpoint, relative to the port, to which the HTTP GET request should be directed.
			path: string
			// +usage=The TCP socket within the container to which the HTTP GET request should be directed.
			port: int
			httpHeaders?: [...{
				name:  string
				value: string
			}]
		}

		// +usage=Instructions for assessing container health by probing a TCP socket. Either this attribute or the exec attribute or the httpGet attribute MUST be specified. This attribute is mutually exclusive with both the exec attribute and the httpGet attribute.
		tcpSocket?: {
			// +usage=The TCP socket within the container that should be probed to assess container health.
			port: int
		}

		// +usage=Number of seconds after the container is started before the first probe is initiated.
		initialDelaySeconds: *0 | int

		// +usage=How often, in seconds, to execute the probe.
		periodSeconds: *10 | int

		// +usage=Number of seconds after which the probe times out.
		timeoutSeconds: *1 | int

		// +usage=Minimum consecutive successes for the probe to be considered successful after having failed.
		successThreshold: *1 | int

		// +usage=Number of consecutive failures required to determine the container is not alive (liveness probe) or not ready (readiness probe).
		failureThreshold: *3 | int

		// +usage=Specify the hostAliases to add
		hostAliases: [...{
			ip: string
			hostnames: [...string]
		}]
	}
}
