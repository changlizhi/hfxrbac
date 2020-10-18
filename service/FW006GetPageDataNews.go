package service

import "hfxrbac/daosnews"

func GetPageDataNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := daosnews.InsertNewsT(canShu)
	return ret
}
