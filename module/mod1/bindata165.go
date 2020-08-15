package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPyBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPy, nil
}

func dictMod1IpaNeologdDictPy() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.py", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
