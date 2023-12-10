package user

import "errors"

var (
	errConfigDirNotInited = errors.New("config directory was not inited")
	errPassedContentIsNil = errors.New("Can't fill resume.yaml with nil slice")
)
