package util

import (
	"log"
	"testing"
)

func TestXxx(t *testing.T) {
	sault := GetSault(5)
	log.Println(EncodingPwd(sault, "test"), len(EncodingPwd(sault, "test")))
}
