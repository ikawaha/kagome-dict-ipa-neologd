package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPmBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPm, nil
}

func dictMod1IpaNeologdDictPm() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
