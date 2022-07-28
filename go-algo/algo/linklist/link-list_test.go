// DATE: 2022/1/16 11:58
// DESC: 测试

package linklist

import (
	"testing"
)

func TestBase(t *testing.T) {
	header := NewLinkNode()
	arr := []interface{}{1, 2, 1, 4, 2, 6, 4, 2}
	CreateLink(header, arr...)
	PrintLink(header)
	RemoveLinkDup(header)
	PrintLink(header)
}
