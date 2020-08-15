package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCxBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCx, nil
}

func dictMod0IpaNeologdDictCx() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
