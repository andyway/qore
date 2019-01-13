// +build enterprise

package migrations

import "github.com/andyway/qore/app/enterprise"

func init() {
	AutoMigrate(&enterprise.QorMicroSite{})
}
