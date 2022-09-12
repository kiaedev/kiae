"k-route": {
	alias: ""
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "networking.istio.io/v1beta1"
		kind:       "VirtualService"
	}
	description: ""
	labels: {}
	type: "component"
}

template: {
	output: {
		apiVersion: "networking.istio.io/v1beta1"
		kind:       "VirtualService"
		spec: {
			hosts: [ for gateway in parameter.gateways {
				gateway.host
			}]
			gateways: [ for gateway in parameter.gateways {
				gateway.name
			}]

			http: [ for r in parameter.routes {
				{
					name: r.name
					if r.methods != _|_ {
						match: [ for m in r.methods {
							{
								method: exact: m
								if r["exactMatch"] == _|_ {
									uri: prefix: r.uri
								}
								if r["exactMatch"] != _|_ {
									if r["exactMatch"] == false {
										uri: prefix: r.uri
									}
									if r["exactMatch"] == true {
										uri: exact: r.uri
									}
								}
							}
						}]
					}

					if r["directResponse"] != _|_ {
						directResponse: {
							body: string: r.directResponse.body
							status: r.directResponse.code
						}
					}

					if r["redirect"] != _|_ {
						redirect: {
							redirectCode: r.redirect.code
						}
						if r.redirect.url.schema != _|_ {
							redirect: schema: r.redirect.url.schema
						}
						if r.redirect.url.host != _|_ {
							redirect: authority: r.redirect.url.host
						}
						if r.redirect.url.path != _|_ {
							redirect: uri: r.redirect.url.path
						}
					}

					if r.forward.rewrite.authority != _|_ {
						rewrite: authority: r.forward.rewrite.authority
					}
					if r.forward.rewrite.uri != _|_ {
						rewrite: uri: r.forward.rewrite.uri
					}

					if r.forward.cors.allowOrigins != _|_ {
						corsPolicy: allowOrigins: [ for origin in r.forward.cors.allowOrigins {
							{
								// todo how to handle *
								exact: origin
							}
						}]
					}
					if r.forward.cors.allowMethods != _|_ {
						corsPolicy: allowMethods: r.forward.cors.allowMethods
					}
					if r.forward.cors.allowHeaders != _|_ {
						corsPolicy: allowHeaders: r.forward.cors.allowHeaders
					}
					if r.forward.cors.allowCredentials != _|_ {
						corsPolicy: allowCredentials: r.forward.cors.allowCredentials
					}

					if r.forward.cors.maxAge != _|_ {
						corsPolicy: maxAge: r.forward.cors.maxAge
					}

					if r["directResponse"] == _|_ && r["redirect"] == _|_ {
						route: [{
							destination: {
								host: context.appName
							}
						}]
					}
				}
			}]
		}
	}
	outputs: {
		kLimiter: {
			apiVersion: "networking.istio.io/v1alpha3"
			kind:       "EnvoyFilter"
			metadata: name: context.name
			spec: {
				workloadSelector: {
					labels: "kiae.dev/component": context.appName
				}
				configPatches: [
					{
						applyTo: "HTTP_FILTER"
						match: {
							context: "SIDECAR_INBOUND"
							listener: filterChain: filter: name: "envoy.filters.network.http_connection_manager"
						}
						patch: {
							operation: "INSERT_BEFORE"
							value: {
								name: "envoy.filters.http.local_ratelimit"
								typed_config: {
									"@type":  "type.googleapis.com/udpa.type.v1.TypedStruct"
									type_url: "type.googleapis.com/envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit"
									value: stat_prefix: "http_local_rate_limiter"
								}
							}
						}
					},
//					{
//						applyTo: "HTTP_ROUTE"
//						match: {
//							context: "SIDECAR_INBOUND"
//							routeConfiguration: vhost: name: "inbound|http|{{.ServicePort}}"
//						}
//						patch: {
//							operation: "INSERT_FIRST"
//							value: {
//								name: context.name
//							}
//						}
//					},
				]
			}
		}
	}
	parameter: {

		gateways?: [...{
			name: string
			host: string
		}]

		routes: [...{
			name: string

			uri: string

			methods: [...string]

			exactMatch?: bool

			forward?: #Forward

			redirect?: #Redirect

			directResponse?: #DirectResponse
		}]
	}

	#Forward: {
		cors?: #CORS
		rewrite?: {
			uri?:       string
			authority?: string
		}
		limiter?: {
			qps: uint32
			fallback: {
				url: string, body: string
			}
		}
	}

	#CORS: {
		enabled: bool
		allowOrigins?: [...string]
		allowMethods?: [...string]
		allowHeaders?: [...string]
		exposeHeaders?: [...string]
		maxAge?:           uint32
		allowCredentials?: bool
	}

	#Redirect: {
		url:  #URL
		code: uint32
	}

	#DirectResponse: {
		body: string
		code: uint32
	}

	#URL: {
		schema?: string
		host?:   string
		path?:   string
	}
}
