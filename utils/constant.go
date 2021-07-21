package utils

const (
	ConfigEnv  		= "JSS_CONFIG"
	ConfigFile 		= "config.yaml"
)

// Pick1Of2
// @Description: 自定义3目运管，二选一
// @param flag
// @param sel1
// @param sel2
// @return interface{}
func Pick1Of2(flag bool, sel1, sel2 interface{}) interface{} {
	if flag {
		return sel1
	} else {
		return sel2
	}
}