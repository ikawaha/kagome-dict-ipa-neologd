package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCqBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCq, nil
}

func dictMod0IpaNeologdDictCq() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
