package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCsBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCs, nil
}

func dictMod0IpaNeologdDictCs() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cs", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
