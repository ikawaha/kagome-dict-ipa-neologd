package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCc, nil
}

func dictMod0IpaNeologdDictCc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
