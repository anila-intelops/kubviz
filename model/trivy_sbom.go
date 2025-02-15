package model

import (
	"github.com/aquasecurity/trivy/pkg/sbom/cyclonedx"
)

type Sbom struct {
	ID     string
	Report cyclonedx.BOM
}

type SbomData struct {
	ID               string
	ClusterName string
	ComponentName    string
	PackageName string
	PackageUrl       string
	BomRef           string
	SerialNumber     string
	CycloneDxVersion int
	BomFormat        string
}


