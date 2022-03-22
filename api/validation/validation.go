package validation

import (
	//"fmt"

	validationutils "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	//webappv1 "my.domain/guestbook/api/v1"
)

var guestbooklog = logf.Log.WithName("guestbook-resource")

func ValidateName(name string) *field.Error {
	guestbooklog.Info("Validating object", "name", name)
	if len(name) > validationutils.DNS1035LabelMaxLength-11 {
		return field.Invalid(field.NewPath("metadata").Child("name"), name, "must be no more than 52 characters")
	}
	return nil
}
