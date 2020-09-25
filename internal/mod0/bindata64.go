package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCmBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCm, nil
}

func dictMod0IpaNeologdDictCm() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
