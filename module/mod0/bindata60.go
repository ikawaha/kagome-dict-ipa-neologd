package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCi, nil
}

func dictMod0IpaNeologdDictCi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ci", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
