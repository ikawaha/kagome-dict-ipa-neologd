package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEc, nil
}

func dictMod0IpaNeologdDictEc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ec", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
