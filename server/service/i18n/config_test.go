package i18n

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

var X = `
REPLACE INTO sys_i18n_localize_config (lang_tag, message_id, template_data, category_id, error_code, created_at, updated_at, deleted_at)
VALUES
	('%s', '%s', '%s', '0', '0', '%v', '%v', '%v');
`

func TestSysI18nLocalizeConfigService_CreateSysI18nLocalizeConfig(t *testing.T) {
	f, _ := os.Open("a.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	type field struct {
		a string
		b string
	}
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "   ")
		var c field
		for i := range a {
			if len(a[i]) > 0 && len(c.a) == 0 {
				c.a = a[i]
				continue
			}
			if len(a[i]) > 0 && len(c.b) == 0 {
				c.b = a[i]
			}
		}
		if c.a == "" {
			continue
		}
		s := fmt.Sprintf(X, "en", c.a, c.b, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
		println(s)
	}
}
