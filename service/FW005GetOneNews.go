package service

import "hfxrbac/daosnews"

func GetOneNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := daosnews.QueryOneNews(canShu)
	return ret
}
