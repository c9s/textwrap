package textwrap

import "testing"


func TestWrap(t *testing.T) {
	str := Wrap(`Mauris mauris ante, blandit et, ultrices a, suscipit eget, quam. Integer ut neque. Vivamus nisi metus, molestie vel, gravida in, condimentum sit amet, nunc. Nam a nibh. Donec suscipit eros. Nam mi. Proin viverra leo ut odio. Curabitur malesuada. Vestibulum a velit eu ante scelerisque vulputate.`, 50, "  ", "  ")
	t.Log(str)
}



