package steps

import (
	"time"

	"github.com/supergiant/supergiant/pkg/runner"
)

type CertificatesConfig struct {
	KubernetesConfigDir   string `json:"kubernetesConfigDir"`
	CACert                string `json:"caCert"`
	CACertName            string `json:"caCertName"`
	CAKeyCert             string `json:"caKeyCert"`
	CAKeyName             string `json:"caKeyName"`
	APIServerCert         string `json:"apiServerCert"`
	APIServerCertName     string `json:"apiServerCertName"`
	APIServerKey          string `json:"apiServerKey"`
	APIServerKeyName      string `json:"apiServerKeyName"`
	KubeletClientCert     string `json:"kubeletClientCert"`
	KubeletClientCertName string `json:"kubeletClientCertName"`
	KubeletClientKey      string `json:"kubeletClientKey"`
	KubeletClientKeyName  string `json:"kubeletClientKeyName"`
}

type DOConfig struct {
	Name         string   `json:"name" valid:"required"`
	K8sVersion   string   `json:"k8sVersion"`
	Region       string   `json:"region" valid:"required"`
	Size         string   `json:"size" valid:"required"`
	Role         string   `json:"role" valid:"in(master|node)"` // master/node
	Image        string   `json:"image" valid:"required"`
	Fingerprints []string `json:"fingerprints" valid:"required"`
	AccessToken  string   `json:"accessToken" valid:"required"`
}

type FlannelConfig struct {
	Arch           string `json:"arch"`
	Version string 	`json:"version"`
	Network        string `json:"network"`
	NetworkType    string `json:"networkType"`
}

type KubeletConfig struct {
	MasterPrivateIP string `json:"masterPrivateIp"`
	ProxyPort       string `json:"proxyPort"`
	EtcdClientPort  string `json:"etcdClientPort"`
	K8SVersion      string `json:"k8sVersion"`
}

type KubeletConfConfig struct {
	Host string `json:"host" valid:"required"`
	Port string `json:"port" valid:"required"`
}

type KubeProxyConfig struct {
	MasterPrivateIP string `json:"masterPrivateIp"`
	ProxyPort       string `json:"proxyPort"`
	EtcdClientPort  string `json:"etcdClientPort"`
	K8SVersion      string `json:"k8sVersion"`
}

type ManifestConfig struct {
	K8SVersion          string `json:"k8sVersion"`
	KubernetesConfigDir string `json:"kubernetesConfigDir"`
	RBACEnabled         bool   `json:"rbacEnabled"`
	EtcdHost            string `json:"etcdHost"`
	EtcdPort            string `json:"etcdPort"`
	PrivateIpv4         string `json:"privateIpv4"`
	ProviderString      string `json:"providerString"`
	MasterHost          string `json:"masterHost"`
	MasterPort          string `json:"masterPort"`
}

type PostStartConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Username    string `json:"username"`
	RBACEnabled bool   `json:"rbacEnabled"`
}

type KubeletSystemdServiceConfig struct {
	K8SVersion         string `json:"k8sVersion"`
	KubeletService     string `json:"kubeletService"`
	KubernetesProvider string `json:"kubernetesProvider"`
}

type TillerConfig struct {
	HelmVersion     string `json:"helmVersion"`
	OperatingSystem string `json:"operatingSystem"`
	Arch            string `json:"arch"`
}

type DockerConfig struct {
	DockerVersion  string `json:"dockerVersion"`
	ReleaseVersion string `json:"releaseVersion"`
	Arch           string `json:"arch"`
}

type DownloadK8sBinary struct {
	K8SVersion      string `json:"k8sVersion"`
	Arch            string `json:"arch"`
	OperatingSystem string `json:"operatingSystem"`
}

type Config struct {
	DigitalOceanConfig DOConfig `json:"digitalOceanConfig"`

	DockerConfig                DockerConfig                `json:"dockerConfig"`
	DownloadK8sBinary           DownloadK8sBinary           `json:"downloadK8sBinary"`
	CertificatesConfig          CertificatesConfig          `json:"certificatesConfig"`
	FlannelConfig               FlannelConfig               `json:"flannelConfig"`
	KubeletConfig               KubeletConfig               `json:"kubeletConfig"`
	KubeletConfConfig           KubeletConfConfig           `json:"kubeletConfConfig"`
	KubeProxyConfig             KubeProxyConfig             `json:"kubeProxyConfig"`
	ManifestConfig              ManifestConfig              `json:"manifestConfig"`
	PostStartConfig             PostStartConfig             `json:"postStartConfig"`
	KubeletSystemdServiceConfig KubeletSystemdServiceConfig `json:"kubeletSystemdServiceConfig"`
	TillerConfig                TillerConfig                `json:"tillerConfig"`

	CloudAccountName string        `json:"cloudAccountName" valid:"required, length(1|32)"`
	Timeout          time.Duration `json:"timeout"`
	runner.Runner    `json:"-"`
}
