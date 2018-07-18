package current_date

import (
	"fmt"
	"reflect"
	"time"
)

// GetCurrentDate returns Time.now()
func GetCurrentDate() (date string) {
	now := time.Now().Local()

	fmt.Println(reflect.TypeOf(now))

	return ""
}
