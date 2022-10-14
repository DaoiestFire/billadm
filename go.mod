module ljw/billadm

go 1.18

replace k8s.io/apimachinery v0.0.0 => github.com/kubernetes/apimachinery v0.25.3

require (
	k8s.io/apimachinery v0.0.0
	k8s.io/klog/v2 v2.80.1
)

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/spf13/cobra v1.6.1
)
