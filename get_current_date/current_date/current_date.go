package current_date

import (
	"fmt"
	"time"
)

func Get_current_date() (date string) {
	now := time.Now().Local()

	fmt.Println("%T", now)

	return ""
}
