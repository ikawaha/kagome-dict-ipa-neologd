package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFe, nil
}

func dictMod0IpaNeologdDictFe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fe", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
