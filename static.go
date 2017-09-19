package logacef

import "strings"

func cef_fs_escape(s string, fs string) string {
	replacer := strings.NewReplacer(`\`, `\\`, fs, `\`+fs)
	e := replacer.Replace(s)

	return e
}
func cef_eao_escape(s string, fs string) string {
	replacer := strings.NewReplacer(fs, `\`+fs)
	e := replacer.Replace(s)

	return e
}
