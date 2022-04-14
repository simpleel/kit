package strings_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.simpleel.com/kit/strings"
)

func TestToSnakeCase(t *testing.T) {
	s := strings.ToSnakeCase("LoginEmailAbbCbbA-A")
	fmt.Println(s)
	if !assert.Equal(t, s, "login_email_abb_cbb") {
		t.Fatal()
	}
}
