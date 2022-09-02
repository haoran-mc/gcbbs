package column

import (
	"database/sql/driver"
	"encoding/json"
)

// SA string array type
type SA []string

// Value ...
func (c SA) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan ...
func (c *SA) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &c)
}
