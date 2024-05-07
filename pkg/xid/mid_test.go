package xid_test

import (
	"fmt"
	"github.com/leiphp/alipay/pkg/xid"
	"testing"
)

func TestNewXID(t *testing.T) {
	fmt.Println(xid.NewMID())
}

func BenchmarkNewXID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xid.NewMID()
	}
}
