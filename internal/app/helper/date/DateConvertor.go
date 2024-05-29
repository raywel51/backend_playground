package date

import "time"

func ConvertToCustomFormat(inputDateString string) (string, error) {
	layout := "2006-01-02T15:04:05.999999-07:00"
	parsedTime, err := time.Parse(layout, inputDateString)
	if err != nil {
		return "", err
	}

	outputFormat := "02-01-2006 | 15:04"
	outputDateString := parsedTime.Format(outputFormat)

	return outputDateString, nil
}
