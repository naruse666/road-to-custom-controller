/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	samplecontrolleriov1alpha "github.com/naruse666/road-to-custom-controller/chapter-7/api/v1alpha"
)

// nolint:unused
// log is for logging in this package.
var foolog = logf.Log.WithName("foo-resource")

// SetupFooWebhookWithManager registers the webhook for Foo in the manager.
func SetupFooWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&samplecontrolleriov1alpha.Foo{}).
		WithValidator(&FooCustomValidator{}).
		WithDefaulter(&FooCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-samplecontroller-io-v1alpha-foo,mutating=true,failurePolicy=fail,sideEffects=None,groups=samplecontroller.io,resources=foos,verbs=create;update,versions=v1alpha,name=mfoo-v1alpha.kb.io,admissionReviewVersions=v1

// FooCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Foo when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type FooCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &FooCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Foo.
func (d *FooCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	foo, ok := obj.(*samplecontrolleriov1alpha.Foo)

	if !ok {
		return fmt.Errorf("expected an Foo object but got %T", obj)
	}
	foolog.Info("Defaulting for Foo", "name", foo.GetName())

	if foo.Spec.Replicas == nil {
		foo.Spec.Replicas = new(int32)
		*foo.Spec.Replicas = 1
	}

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-samplecontroller-io-v1alpha-foo,mutating=false,failurePolicy=fail,sideEffects=None,groups=samplecontroller.io,resources=foos,verbs=create;update,versions=v1alpha,name=vfoo-v1alpha.kb.io,admissionReviewVersions=v1

// FooCustomValidator struct is responsible for validating the Foo resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type FooCustomValidator struct {
	//TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &FooCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Foo.
func (v *FooCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	foo, ok := obj.(*samplecontrolleriov1alpha.Foo)
	if !ok {
		return nil, fmt.Errorf("expected a Foo object but got %T", obj)
	}
	foolog.Info("Validation for Foo upon creation", "name", foo.GetName())

	return nil, validateFoo(foo)
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Foo.
func (v *FooCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	foo, ok := newObj.(*samplecontrolleriov1alpha.Foo)
	if !ok {
		return nil, fmt.Errorf("expected a Foo object for the newObj but got %T", newObj)
	}
	foolog.Info("Validation for Foo upon update", "name", foo.GetName())

	return nil, validateFoo(foo)
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Foo.
func (v *FooCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	foo, ok := obj.(*samplecontrolleriov1alpha.Foo)
	if !ok {
		return nil, fmt.Errorf("expected a Foo object but got %T", obj)
	}
	foolog.Info("Validation for Foo upon deletion", "name", foo.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}

func validateDeploymentName(foo *samplecontrolleriov1alpha.Foo) *field.Error {
	// object name must be no more than 253 characters.
	if len(foo.Spec.DeploymentName) > 253 {
		return field.Invalid(field.NewPath("spec").Child("deploymentName"),
			foo.Spec.DeploymentName, "must be no more than 253 characters")
	}
	return nil
}

func validateFoo(foo *samplecontrolleriov1alpha.Foo) error {
	var allErrs field.ErrorList
	if err := validateDeploymentName(foo); err != nil {
		allErrs = append(allErrs, err)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "samplecontroller", Kind: "Foo"},
		foo.Name, allErrs)
}
