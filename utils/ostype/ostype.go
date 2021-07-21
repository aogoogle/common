package ostype

const (
	ANDROID 		= "android"
	IOS 			= "ios"
	WEB 			= "web"
	MOBILE 			= "mobile"
	WIN 			= "pc"
)

func IsPhone(os string) bool {
	return os == ANDROID || os == IOS
}