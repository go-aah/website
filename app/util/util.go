package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strings"

	aah "aahframework.org/aah.v0"
	"aahframework.org/log.v0"
)

const (
	signaturePrefix = "sha1="

	// len(SignaturePrefix) + len(hex(sha1))
	signatureLength = 45
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
	computed := hmac.New(sha1.New, []byte(secret))
	_, _ = computed.Write(payload)

	return hmac.Equal(computed.Sum(nil), actual)
}
