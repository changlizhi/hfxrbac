package service
//
//import (
//	"hfxrbac/consts"
//	"hfxrbac/daosnews"
//	"hfxrbac/utils"
//	"log"
//)
//
//func JiaoYanFw002(canShu interface{}) interface{} {
//	ret := interface{}{}
//	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
//	log.Println("id-----", id)
//	if utils.ShiFouKongZiFu(id) {
//		ret[consts.ZhuangTai] = consts.ShiBai
//		ret[consts.ShuoMing] = "失败:Id主键不能为空"
//		return ret
//	}
//
//	return ret
//}
//func DelNewsService(canShu interface{}) interface{} {
//	retJiaoYan := JiaoYanFw002(canShu)
//	if retJiaoYan[consts.ZhuangTai] == consts.ShiBai {
//		return retJiaoYan
//	}
//	//可以设置表的标题字段为唯一索引就不需要在插入之前校验了
//	ret := daosnews.DeleteNews(canShu)
//	return ret
//}
