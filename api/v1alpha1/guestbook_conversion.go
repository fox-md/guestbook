package v1alpha1

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	webappv1 "my.domain/guestbook/api/v1"
)

// ConvertTo converts this Guestbook to the Hub version (v1).
func (src *Guestbook) ConvertTo(dstRaw conversion.Hub) error {
	fmt.Println("v1alpha1 -> v1")

	dst := dstRaw.(*webappv1.Guestbook)

	// Spec
	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Bar = "AlphaBar"
	dst.Spec.Converted.AlphaField = src.Spec.AlphaField

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *Guestbook) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*webappv1.Guestbook)

	// var v1Spec GuestbookSpec

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.AlphaField = src.Spec.Converted.AlphaField

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	return nil
}
