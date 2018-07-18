package current_date

import (
	"fmt"
	"time"
)

func Get_current_date() {
	now := time.Now().Local()

	fmt.Println("%T", now)
}
