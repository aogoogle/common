package request

type JWT struct {
	SignKey 	string 	`mapstructure:"sign_key" json:"SignKey" yaml:"sign_key"`
	ExpiresTime int 	`mapstructure:"expries_time" json:"ExpriesTime" yaml:"expries_time"`
	BufferTime 	int 	`mapstructure:"buffer_time" json:"BufferTime" yaml:"buffer_time"`
}
