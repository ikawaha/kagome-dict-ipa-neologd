package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFc, nil
}

func dictMod0IpaNeologdDictFc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
