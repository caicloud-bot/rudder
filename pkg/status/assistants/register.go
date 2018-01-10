package assistants

import "github.com/caicloud/release-controller/pkg/status"

// Register registers all assistant for an umpire.
func Register(umpire status.Umpire) {
	umpire.Employ(gvkService, ServiceAssistant)
	umpire.Employ(gvkDeployment, DeploymentAssistant)
	umpire.Employ(gvkStatefulSet, ServiceAssistant)
	umpire.Employ(gvkDaemonSet, ServiceAssistant)
	umpire.Employ(gvkJob, ServiceAssistant)
	umpire.Employ(gvkCronJob, ServiceAssistant)
	umpire.Employ(gvkPVC, ServiceAssistant)
}