package v1beta1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var vmsinglelog = logf.Log.WithName("vmsingle-resource")

func (r *VMSingle) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,admissionReviewVersions=v1,sideEffects=none,path=/validate-operator-victoriametrics-com-v1beta1-vmsingle,mutating=false,failurePolicy=fail,groups=operator.victoriametrics.com,resources=vmsingles,versions=v1beta1,name=vvmsingle.kb.io

var _ webhook.Validator = &VMSingle{}

func (r *VMSingle) sanityCheck() error {
	if r.Spec.VMBackup != nil {
		if err := r.Spec.VMBackup.sanityCheck(r.Spec.License); err != nil {
			return err
		}
	}

	return nil
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VMSingle) ValidateCreate() (admission.Warnings, error) {
	if r.Spec.ParsingError != "" {
		return nil, fmt.Errorf(r.Spec.ParsingError)
	}
	if mustSkipValidation(r) {
		return nil, nil
	}
	if err := r.sanityCheck(); err != nil {
		return nil, err
	}
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VMSingle) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	if r.Spec.ParsingError != "" {
		return nil, fmt.Errorf(r.Spec.ParsingError)
	}
	if mustSkipValidation(r) {
		return nil, nil
	}
	if err := r.sanityCheck(); err != nil {
		return nil, err
	}
	// todo check for negative storage resize.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VMSingle) ValidateDelete() (admission.Warnings, error) {
	// no-op
	return nil, nil
}
