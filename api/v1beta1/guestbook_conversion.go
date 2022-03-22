package v1beta1

import (
	webappv1 "my.domain/guestbook/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var guestbooklog = logf.Log.WithName("guestbook-resource")

// ConvertTo converts this Guestbook to the Hub version (v1).
func (src *Guestbook) ConvertTo(dstRaw conversion.Hub) error {
	//fmt.Println("v1beta1 -> v1")
	guestbooklog.Info("Converting v1beta1 -> v1", "name", src.Name)

	dst := dstRaw.(*webappv1.Guestbook)

	// Spec
	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Bar = "BetaBar"
	dst.Spec.Converted.BetaField = src.Spec.BetaField

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *Guestbook) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*webappv1.Guestbook)

	// var v1Spec GuestbookSpec

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.BetaField = src.Spec.Foo

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	return nil
}
