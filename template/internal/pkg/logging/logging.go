package logging

import "github.com/sirupsen/logrus"

func Init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.AddHook(NewContextHook([]string{
		"url_path",
		"method",
		"request_id",
		"workspace",
		"user_id",
	}))
}
