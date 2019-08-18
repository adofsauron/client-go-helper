package kubectl

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type delete interface {
	Delete(opts *v1.DeleteOptions) error
}

type describe interface {
	Describe() (string, error)
}

type get interface {
	Get() error
}

type exec interface {
	Exec(cmd []string) error
}

type cp interface {
	Cp(srcPath string, destPath string) error
}

type logs interface {
	Logs(podLogOpts *corev1.PodLogOptions) (string, error)
}

//used for apply/create
type Kubeapi struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Yaml       string
}
