// This is a generated file. Do not edit directly.

module k8s.io/kubelet

go 1.16

require (
	github.com/gogo/protobuf v1.3.2
	golang.org/x/net v0.3.1-0.20221206200815-1e63c2f08a10
	google.golang.org/grpc v1.51.0
	k8s.io/api v0.24.6
	k8s.io/apimachinery v0.24.6
	k8s.io/component-base v0.24.6
)

replace (
	k8s.io/api => k8s.io/api v0.24.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.24.6
	k8s.io/client-go => k8s.io/client-go v0.24.6
	k8s.io/component-base => k8s.io/component-base v0.24.6
	k8s.io/kubelet => ./
)
