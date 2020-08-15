package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGlBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGl, nil
}

func dictMod0IpaNeologdDictGl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
