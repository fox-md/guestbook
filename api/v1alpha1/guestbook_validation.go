package v1alpha1

import (
	"strings"

	validationutils "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"

	//logf "sigs.k8s.io/controller-runtime/pkg/log"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (r *Guestbook) validateGuestbook() error {
	var allErrs field.ErrorList
	if err := r.validateFoo(); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := r.ValidateName(); err != nil {
		allErrs = append(allErrs, err)
	}
	// if err := r.validateBeta(); err != nil {
	// 	allErrs = append(allErrs, err)
	// }

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "webapp.my.domain", Kind: "Guestbook"},
		r.Name, allErrs)
}

func (gb *Guestbook) validateFoo() *field.Error {
	guestbooklog.Info("Validating foo", "foo", gb.Spec.Foo)
	if strings.HasPrefix(gb.Spec.Foo, "BetaAGU") {
		return nil
	} else {
		fldPath := field.NewPath("spec").Child("foo")
		return field.Invalid(fldPath, gb.Spec.Foo, "'foo' field does not start with 'AGU'")
	}
}

// func (gb *Guestbook) validateBeta() *field.Error {
// 	guestbooklog.Info("Validating foo", "foo", gb.Spec.Foo)
// 	if strings.HasPrefix(gb.Spec.BetaField, "Beta") {
// 		return nil
// 	} else {
// 		fldPath := field.NewPath("spec").Child("foo")
// 		return field.Invalid(fldPath, gb.Spec.Foo, "'betafield' field does not start with 'Beta'")
// 	}
// }

func (gb *Guestbook) ValidateName() *field.Error {
	guestbooklog.Info("Validating object", "name", gb.ObjectMeta.Name)
	if len(gb.ObjectMeta.Name) > validationutils.DNS1035LabelMaxLength-11 {
		return field.Invalid(field.NewPath("metadata").Child("name"), gb.ObjectMeta.Name, "must be no more than 52 characters")
	}
	return nil
}
