package flag

import (
	"bufio"
	"os"
	"strings"
)

func SetPackages(flags *Flags, apath string, defaults []string) []string {
	packages := []string{}

	if apath != "" {
		exists, err := IsFile(apath)
		if err != nil {
			panic(err)
		}

		if exists {
			o, err := os.OpenFile(apath, os.O_RDONLY, 0o600)
			if err != nil {
				panic(err)
			}

			l, err := Scanner(o)
			if err != nil {
				panic(err)
			}

			packages = Parse(l, flags)
		}
	} else {
		packages = defaults
	}

	return packages
}

func Parse(list []string, flags *Flags) []string {
	replacements := map[string]string{
		"{HOST}":               flags.Host,
		"{TARGETX64}":          flags.Targetx64,
		"{TARGETX86}":          flags.Targetx86,
		"{TARGETARM}":          flags.Targetarm,
		"{TARGETARM_UPPER}":    strings.ToUpper(flags.Targetarm),
		"{TARGETARM64}":        flags.Targetarm64,
		"{TARGETARM64_UPPER}":  strings.ToUpper(flags.Targetarm64),
		"{MSVC_VERSION}":       flags.MsvcVer,
		"{MSVC_VERSION_LOCAL}": flags.MsvcVerLocal,
		"{WINSDK_VERSION}":     flags.WinSDKVer,
	}

	parsed := []string{}

	for _, item := range list {
		for placeholder, value := range replacements {
			item = strings.ReplaceAll(item, placeholder, value)
		}

		parsed = append(parsed, item)
	}

	return parsed
}

func Scanner(file *os.File) ([]string, error) {
	entries := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}

		entries = append(entries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func IsFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return fileInfo.Mode().IsRegular(), nil
}
