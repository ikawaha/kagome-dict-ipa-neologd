package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCe, nil
}

func dictMod0IpaNeologdDictCe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ce", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
