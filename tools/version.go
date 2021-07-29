package tools

import "encoding/json"

var (
	Name = "shorten url service"

	Version = "Not Provide"

	BuildTime = "Not Provide"

	GitSHA1 = "Not Provide"
)

func PrintVersion() []byte {
	version := map[string]string{
		"Name":      Name,
		"Version":   Version,
		"BuildTime": BuildTime,
		"GitSHA1":   GitSHA1,
	}

	data, _ := json.MarshalIndent(version, "", "   ")
	return data
}
