package test

import (
	"cloud-disk/core/helper"
	"fmt"
	"testing"
)

func TestCod(t *testing.T) {

	fmt.Println("验证码:", helper.RandCode())
}
