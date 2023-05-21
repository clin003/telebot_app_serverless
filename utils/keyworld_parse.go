package utils

import (
	"strings"
)

func KeyworldListParse(keyworldListStr string) []string {
	if len(keyworldListStr) <= 0 {
		return nil
	}

	retArr := make([]string, 0)
	listText := keyworldListStr
	list := strings.Split(listText, "|")
	for _, v := range list {
		vc := v
		if len(vc) > 0 {
			retArr = append(retArr, vc)
		}
	}
	return retArr
}

// 检查屏蔽关键词，关键词，存在屏蔽词(keyworldFilter) 返回false，存在订阅关键词(keyworldList)或无订阅关键词词 返回true
func KeyworldCheck(msgText, keyworldFilter, keyworldList string) (retText, retFilter string, retBool bool) {
	retBool = true

	keyworldFilter := KeyworldListParse(keyworldFilter)
	for _, v := range keyworldFilter {
		vc := v
		if strings.Contains(msgText, vc) {
			retFilter = vc
			retBool = false
			return
		}
	}

	keyworldList := KeyworldListParse(keyworldList)
	if len(keyworldList) <= 0 {
		retText = "无订阅词限定"
		retBool = true
		return
	} else {
		retBool = false
	}
	for _, v := range keyworldList {
		vc := v
		if strings.Contains(msgText, vc) {
			retText = vc
			retBool = true
			return
		}
	}
	return
}
