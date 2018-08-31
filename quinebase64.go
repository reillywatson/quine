// works by base64-encoding this file, but with the base64-encoded string replaced with something else.
package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	str := "Ly8gd29ya3MgYnkgYmFzZTY0LWVuY29kaW5nIHRoaXMgZmlsZSwgYnV0IHdpdGggdGhlIGJhc2U2NC1lbmNvZGVkIHN0cmluZyByZXBsYWNlZCB3aXRoIHNvbWV0aGluZyBlbHNlLgpwYWNrYWdlIG1haW4KCmltcG9ydCAoCgkiZW5jb2RpbmcvYmFzZTY0IgoJImZtdCIKCSJzdHJpbmdzIgopCgpmdW5jIG1haW4oKSB7CglzdHIgOj0gInF1aW5lIgoJYiwgXyA6PSBiYXNlNjQuU3RkRW5jb2RpbmcuRGVjb2RlU3RyaW5nKHN0cikKCWZtdC5QcmludChzdHJpbmdzLlJlcGxhY2Uoc3RyaW5nKGIpLCAicXVpbmUiLCBzdHIsIDEpKQp9Cg=="
	b, _ := base64.StdEncoding.DecodeString(str)
	fmt.Print(strings.Replace(string(b), "quine", str, 1))
}
