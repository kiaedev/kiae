# Kiae - An AppEngine base on kubernetes and istio

[![](https://github.com/kiaedev/kiae/workflows/build/badge.svg)](https://github.com/kiaedev/kiae/actions?query=workflow%3Abuild)
[![](https://codecov.io/gh/kiaedev/kiae/branch/master/graph/badge.svg)](https://codecov.io/gh/kiaedev/kiae)
[![](https://img.shields.io/github/v/release/kiaedev/kiae.svg)](https://github.com/kiaedev/kiae/releases)
[![](https://img.shields.io/github/license/kiaedev/kiae.svg)](https://github.com/kiaedev/kiae/blob/master/LICENSE)

## What is Kiae?

Kiae is a cloud native application develop platform base on the Kubernetes and Istio.

## Why Kiae?

Kubernetes and Istio are declarative softwares, and they are professional. So they are difficult to use for the application developer. For the company team, we usually build an internal cloud platform base on the Kubernetes. But it's always deeply integrated with the internal micro-services, and it's not integrated with Istio.

Kiae built a open-source cloud platform completely base on the Kubernetes and Istio.

## Features

- Git integration (GitHub, BitBucket, GitLab)
- Build image from the Git source repository
- Automatically push image to the image registry
- Deploy any image to multiple environments and multiple clusters
- Application level Observability with the Open-Telemetry
- Dependents management for the applications and the middlewares
- ConfigFiles management base on the ConfigMap and Secret
- Environments management for the multiple environments and multiple clusters
- Routes management base on the Istio VisualServices
- Entrypoint management base on the Istio IngressGateway
- Access controls management base on the Istio AuthzPolicy
- Web UI which provides real-time view of application activity
- SSO Integration (OIDC, OAuth2, LDAP, SAML 2.0, GitHub, GitLab, Microsoft, LinkedIn)

## Documentation
To learn more about Kiae go to the [complete documentation](https://kiae.dev/).

## Community

You can reach the Kiae community and developers via the following channels:

* Q & A: [Github Discussions](https://github.com/kiaedev/kiae/discussions)

## Special thanks

[![JetBrains](https://raw.githubusercontent.com/kainonly/ngx-bit/main/resource/jetbrains.svg)](https://www.jetbrains.com/?from=saltbo)

Thanks for non-commercial open source development authorization by JetBrains

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.

## License

Kiae is under the Apache-2.0 license. See the [LICENSE](/LICENSE) file for details.
