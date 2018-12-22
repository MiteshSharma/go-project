package setting

import "encoding/json"

type Setting struct {
	BuildNo   string
	Version   string
	Branch    string
	Commit    string
	StartTime string
}

func NewSetting(buildNo string, version string, branch string, commit string, startTime string) *Setting {
	setting := &Setting{
		BuildNo:   buildNo,
		Version:   version,
		Branch:    branch,
		Commit:    commit,
		StartTime: startTime,
	}
	return setting
}

func (s *Setting) ToJson() string {
	json, _ := json.Marshal(s)
	return string(json)
}
