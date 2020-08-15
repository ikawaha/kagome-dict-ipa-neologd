package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFxBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFx, nil
}

func dictMod0IpaNeologdDictFx() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
