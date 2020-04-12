package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

func Md5(data string) string {
	t := md5.New();
	io.WriteString(t, data);
	return fmt.Sprintf("%x", t.Sum(nil));
}

func SHA1(data string) string {
	t := sha1.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}
