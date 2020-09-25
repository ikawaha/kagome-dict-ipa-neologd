package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictArBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAr, nil
}

func dictMod0IpaNeologdDictAr() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictArBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ar", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
