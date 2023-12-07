package strings

import (
	"MyCodeArchive_Go/logging"
	"strconv"
)

func ConvertPageToInt(pageNumber, pageSize string) (int, int, error) {
	logging.Log.Infof("start covert page number and page size to int: %s %s", pageNumber, pageSize)
	number, err := strconv.Atoi(pageNumber)
	if err != nil {
		return -1, -1, err
	}
	size, err := strconv.Atoi(pageSize)
	if err != nil {
		return -1, -1, err
	}
	return number, size, nil
}
