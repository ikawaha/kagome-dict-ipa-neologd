package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGmBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGm, nil
}

func dictMod0IpaNeologdDictGm() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
