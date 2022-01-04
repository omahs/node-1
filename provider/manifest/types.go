package manifest

import (
	manifest "github.com/ovrclk/akash/manifest/v2beta1"
	dtypes "github.com/ovrclk/akash/x/deployment/types/v1beta2"
)

// Status is the data structure
type Status struct {
	Deployments uint32 `json:"deployments"`
}

type submitRequest struct {
	Deployment dtypes.DeploymentID `json:"deployment"`
	Manifest   manifest.Manifest   `json:"manifest"`
}
