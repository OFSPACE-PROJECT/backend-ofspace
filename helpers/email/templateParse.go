package email

import (
	"bytes"
	"fmt"
	"html/template"
	"ofspace-be/features/booking"
	"strconv"
	"strings"
)

func templateParse(templateFileName string, data interface{}, bookData booking.Core) (string, error) {
	templatePath := fmt.Sprintf("helpers/email/message/%s", templateFileName)

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	r := strings.NewReplacer("#Name#", bookData.ConfirmedName, "#Price#", strconv.Itoa(int(bookData.Price)), "#Type#", bookData.Unit.UnitType, "#Date#", bookData.DealDate.String())
	bodyReplace := r.Replace(body)
	return bodyReplace, nil
}
