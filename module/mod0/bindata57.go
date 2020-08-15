package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCf, nil
}

func dictMod0IpaNeologdDictCf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
