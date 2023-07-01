package pathutils

import "os/user"

func GetHomeDir() string {
	u, err := user.Current()
	if err != nil {
		return "/tmp"
	}
	return u.HomeDir
}
