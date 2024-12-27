package v1beta

import (
	samplecontrollerv1alpha "github.com/naruse666/road-to-custom-controller/chapter-7/api/v1alpha"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// Convert Foo to Hub version (v1alpha)
func (src *Foo) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*samplecontrollerv1alpha.Foo)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.DeploymentName = src.Spec.DeploymentName
	dst.Spec.Replicas = src.Spec.Replicas

	// Status
	dst.Status.AvailableReplicas = src.Status.AvailableReplicas
	return nil
}

// Convert from Hub version to this version
func (dst *Foo) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*samplecontrollerv1alpha.Foo)

	// copy DeploymentName to Foo
	dst.Spec.Foo = src.Spec.DeploymentName

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.DeploymentName = src.Spec.DeploymentName
	dst.Spec.Replicas = src.Spec.Replicas

	// Status
	dst.Status.AvailableReplicas = src.Status.AvailableReplicas

	return nil
}
