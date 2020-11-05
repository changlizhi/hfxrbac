package daosnews
//
//import (
//	"hfxrbac/consts"
//	"hfxrbac/daos"
//	"hfxrbac/utils"
//)
//
//func DeleteNews(canShu interface{}) interface{} {
//	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
//	insertNeedCanShu := interface{}{
//		consts.ShuJuKu:   consts.Hfxrbac,
//		consts.ShuJuBiao: consts.News,
//		consts.TiaoJians: interface{}{
//			consts.Id: id,
//		},
//	}
//	ret := daos.DeleteData(insertNeedCanShu)
//
//	return ret
//}
