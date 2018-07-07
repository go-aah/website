package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"aahframework.org/aah.v0"
	"aahframework.org/ahttp.v0"
	"aahframework.org/log.v0"
	"github.com/hashicorp/go-version"
)

const (
	signaturePrefix = "sha1="

	// len(SignaturePrefix) + len(hex(sha1))
	signatureLength = 45
)

// Version number cleaner
var (
	VerRep    = strings.NewReplacer("v", "", ".x", "", "-edge", "")
	VerKeyRep = strings.NewReplacer(".x", "", "-edge", "", ".", "")
)

// ExecCmd method to execute command line arguments.
func ExecCmd(cmdName string, args []string, stdout bool) (string, error) {
	cmd := exec.Command(cmdName, args...)
	log.Info("Executing ", strings.Join(cmd.Args, " "))

	if stdout {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return "", err
		}
		_ = cmd.Wait()
	} else {
		bytes, err := cmd.CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("\n%s\n%s", string(bytes), err)
		}

		return string(bytes), nil
	}

	return "", nil
}

// IsValidHubSignature method returns true the signature is correct otherwise false.
func IsValidHubSignature(signature string, payload []byte) bool {
	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	_, _ = hex.Decode(actual, []byte(signature[5:]))

	secret := aah.AppConfig().StringDefault("docs.github_secert", "")
	log.Info("Github Hook: ", secret)
	computed := hmac.New(sha1.New, []byte(secret))
	_, _ = computed.Write(payload)

	return hmac.Equal(computed.Sum(nil), actual)
}

// AllowAllOriginForStaticFiles adds `AccessControlAllowOrigin: *` header on static file response.
func AllowAllOriginForStaticFiles(e *aah.Event) {
	ctx := e.Data.(*aah.Context)
	if ctx.IsStaticRoute() {
		ctx.Res.Header().Set(ahttp.HeaderAccessControlAllowOrigin, "*")
	}
}

// VersionGtEq method compare two semantic versions
func VersionGtEq(currentVersion, expectedVersion string) bool {
	cv, err := version.NewVersion(VerRep.Replace(currentVersion))
	if err != nil {
		log.Error("Current version: ", currentVersion, err)
		return false
	}

	ev, err := version.NewVersion(VerRep.Replace(expectedVersion))
	if err != nil {
		log.Error("Expected version: ", expectedVersion, err)
		return false
	}
	return (cv.Equal(ev) || cv.GreaterThan(ev))
}

// VersionLtEq method compare two semantic versions
func VersionLtEq(currentVersion, expectedVersion string) bool {
	cv, err := version.NewVersion(VerRep.Replace(currentVersion))
	if err != nil {
		log.Error("Current version: ", currentVersion, err)
		return false
	}

	ev, err := version.NewVersion(VerRep.Replace(expectedVersion))
	if err != nil {
		log.Error("Expected version: ", expectedVersion, err)
		return false
	}
	return (cv.Equal(ev) || cv.LessThan(ev))
}

// IsVersionNo method returns true if given string is vaild version no
// otherwise false.
func IsVersionNo(v string) bool {
	_, err := version.NewVersion(VerRep.Replace(v))
	return err == nil
}
