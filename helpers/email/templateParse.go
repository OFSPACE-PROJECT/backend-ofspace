package email

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"ofspace-be/features/booking"
	"path/filepath"
	"strconv"
	"strings"
)

func templateParse(templateFileName string, data interface{}, bookData booking.Core) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("helpers/email/message/%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}
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
