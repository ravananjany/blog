package util

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

func WithFields(pac string, file string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"package": pac,
		"go":      file,
	})
}

func StringTOI(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return id, nil
}
