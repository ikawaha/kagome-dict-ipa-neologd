package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGpBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGp, nil
}

func dictMod0IpaNeologdDictGp() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
