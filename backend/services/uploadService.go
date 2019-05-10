package services

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/CanobbioE/reelo/backend/api"
)

// ParseFileWithInfo reads the content of a buffer and verifies its correctness
// and tries to parse each line into an entity to be saved in the database
// appending the data defined in the upload info
// TODO: implements it
func ParseFileWithInfo(buf bytes.Buffer, info api.UploadInfo) error {
	year, err := strconv.Atoi(info.Year)
	if err != nil {
		return err
	}
	category := strings.ToUpper(info.Category)
	format := info.Format
	isParis := info.IsParis

	// TODO: remove this
	fmt.Println(year, category, format, isParis)

	contents := buf.String()
	fmt.Println(contents)
	buf.Reset()
	return nil
}
