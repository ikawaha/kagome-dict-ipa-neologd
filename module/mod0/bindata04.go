package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAe, nil
}

func dictMod0IpaNeologdDictAe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ae", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
