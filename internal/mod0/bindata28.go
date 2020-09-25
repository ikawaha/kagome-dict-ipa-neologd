package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBc, nil
}

func dictMod0IpaNeologdDictBc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
