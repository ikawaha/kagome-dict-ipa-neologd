package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAsBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAs, nil
}

func dictMod0IpaNeologdDictAs() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.as", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
