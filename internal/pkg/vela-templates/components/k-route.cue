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
								method: m
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

					if parameter["directResponse"] != _|_ {
						directResponse: {
							body: string: parameter.directResponse.body
							status: parameter.directResponse.code
						}
					}

					if parameter["redirect"] != _|_ {
						redirect: {
							schema:       parameter.redirect.url.schema
							authority:    parameter.redirect.url.host
							uri:          parameter.redirect.url.path
							redirectCode: parameter.redirect.code
						}
					}

					if parameter["rewrite"] != _|_ {
						rewrite: {
							authority: parameter.rewrite.authority
							uri:       parameter.rewrite.uri
						}
					}

					if parameter["corsPolicy"] != _|_ {
						corsPolicy: {
							allowCredentials: parameter.corsPolicy.allowCredentials
							allowHeaders:     parameter.corsPolicy.allowHeaders
							allowMethods:     parameter.corsPolicy.allowMethods
							allowOrigins: [ for origin in parameter.corsPolicy.allowOrigins {
								{
									// todo how to handle *
									exact: origin
								}
							}]
							maxAge: parameter.corsPolicy.maxAge
						}
					}

					route: [{
						destination: {
							host: context.appName
						}
					}]
				}
			}]
		}
	}
	outputs: {
		//  kLimiter: {
		//   apiVersion: "networking.istio.io/v1alpha3"
		//   kind:       "EnovyFilter"
		//   metadata: name: context.name
		//   spec: {
		//   }
		//  }
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

			forward: #Forward

			redirect: #Redirect

			directResponse: #DirectResponse
		}]
	}

	#Forward: {
		corsPolicy: #CORS
		rewrite: {
			uri:       string
			authority: string
		}
		limitPolicy: {
			qps: uint32
			fallback: {
				url: string, body: string
			}
		}
	}

	#CORS: {
		allow_origins: [...string]
		allow_methods: [...string]
		allow_headers: [...string]
		expose_headers: [...string]
		max_age:           uint32
		allow_credentials: bool
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
		schema: string
		host:   string
		path:   string
	}
}
