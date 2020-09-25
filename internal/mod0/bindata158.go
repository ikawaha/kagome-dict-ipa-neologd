package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGc, nil
}

func dictMod0IpaNeologdDictGc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
