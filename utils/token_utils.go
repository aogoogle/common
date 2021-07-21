package utils

import (
	"github.com/aogoogle/common/utils/ostype"
	"github.com/aogoogle/common/utils/vtoken"
)

func GenarateAppTokenKey(userId string, clientOS string) (key string) {
	key = vtoken.TokenPrefix + userId +"_"
	if ostype.IsPhone(clientOS) { key += "Phone" } else { key += clientOS }
	return key
}