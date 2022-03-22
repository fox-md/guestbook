package v1

import (
	"strings"

	//validationutils "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"

	//logf "sigs.k8s.io/controller-runtime/pkg/log"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apivalidation "my.domain/guestbook/api/validation"
)

func (r *Guestbook) validateGuestbook() error {
	var allErrs field.ErrorList
	if err := r.validateFoo(); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := apivalidation.ValidateName(r.Name); err != nil {
		allErrs = append(allErrs, err)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "webapp.my.domain", Kind: "Guestbook"},
		r.Name, allErrs)
}

func (gb *Guestbook) validateFoo() *field.Error {
	guestbooklog.Info("Validating foo", "foo", gb.Spec.Foo)
	if strings.HasPrefix(gb.Spec.Foo, "AGU") {
		return nil
	} else {
		fldPath := field.NewPath("spec").Child("foo")
		return field.Invalid(fldPath, gb.Spec.Foo, "'foo' field does not start with 'AGU'")
	}
}
