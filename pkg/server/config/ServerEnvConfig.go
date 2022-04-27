package serverEnvConfig

import "github.com/caarlos0/env"

type ServerEnvConfig struct {
	CanServerUpdate                      bool   `env:"CAN_SERVER_UPDATE" envDefault:"true"` // default true
	InstallerCrdObjectGroupName          string `env:"INSTALLER_CRD_OBJECT_GROUP_NAME" envDefault:"installer.devtron.ai"`
	InstallerCrdObjectVersion            string `env:"INSTALLER_CRD_OBJECT_VERSION" envDefault:"v1alpha1"`
	InstallerCrdObjectResource           string `env:"INSTALLER_CRD_OBJECT_RESOURCE" envDefault:"installers"`
	InstallerCrdNamespace                string `env:"INSTALLER_CRD_NAMESPACE" envDefault:"devtroncd"`
	DevtronHelmRepoName                  string `env:"DEVTRON_HELM_REPO_NAME" envDefault:"devtron"`
	DevtronHelmRepoUrl                   string `env:"DEVTRON_HELM_REPO_URL" envDefault:"https://helm.devtron.ai"`
	DevtronHelmReleaseName               string `env:"DEVTRON_HELM_RELEASE_NAME" envDefault:"devtron"`
	DevtronHelmReleaseNamespace          string `env:"DEVTRON_HELM_RELEASE_NAMESPACE" envDefault:"devtroncd"`
	DevtronHelmReleaseChartName          string `env:"DEVTRON_HELM_RELEASE_CHART_NAME" envDefault:"devtron-operator"`
	DevtronVersionIdentifierInHelmValues string `env:"DEVTRON_VERSION_IDENTIFIER_IN_HELM_VALUES" envDefault:"installer.release"`
	DevtronModulesIdentifierInHelmValues string `env:"DEVTRON_MODULES_IDENTIFIER_IN_HELM_VALUES" envDefault:"installer.modules"`
}

func ParseServerEnvConfig() *ServerEnvConfig {
	cfg := &ServerEnvConfig{}
	err := env.Parse(cfg)
	if err != nil {
		panic("failed to parse server env config: " + err.Error())
	}
	return cfg
}