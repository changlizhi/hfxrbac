package service
//
//import (
//	"hfxrbac/consts"
//	"hfxrbac/daosnews"
//	"hfxrbac/utils"
//)
//
//func JiaoYanFw003(canShu interface{}) interface{} {
//	ret := interface{}{}
//	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
//	if utils.ShiFouKongZiFu(id) {
//		ret[consts.ZhuangTai] = consts.ShiBai
//		ret[consts.ShuoMing] = "失败:Id主键不能为空"
//		return ret
//	}
//
//	return ret
//}
//func UpdateNewsService(canShu interface{}) interface{} {
//	retJiaoYan := JiaoYanFw003(canShu)
//	if retJiaoYan[consts.ZhuangTai] == consts.ShiBai {
//		return retJiaoYan
//	}
//	ret := daosnews.UpdateNews(canShu)
//	return ret
//}
