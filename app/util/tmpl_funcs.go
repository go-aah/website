package util

import (
	"fmt"
	"html/template"
	"strings"
)

// TmplAbsReqURL method create absolute URL
func TmplAbsReqURL(viewArgs map[string]interface{}) template.URL {
	return template.URL(fmt.Sprintf("%s://%s%s", viewArgs["Scheme"], viewArgs["Host"], viewArgs["RequestPath"]))
}

// TmplVerGtEq method compare two versions
func TmplVerGtEq(currentVersion, expectedVersion string) bool {
	return VersionGtEq(currentVersion, expectedVersion)
}

// TmplDVerDis method creates display version string for dropdown
func TmplDVerDis(verVal string) string {
	if strings.HasSuffix(verVal, "-edge") {
		return verVal
	}
	return verVal + ".x"
}
