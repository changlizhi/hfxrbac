package daosnews
//
//import (
//	"hfxrbac/consts"
//	"hfxrbac/daos"
//	"hfxrbac/utils"
//	"time"
//)
//
//func UpdateNews(canShu interface{}) interface{} {
//	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
//	biaoTi := utils.HuoQuZiFuZhi(canShu[consts.BiaoTi])
//	leiXing := utils.HuoQuZiFuZhi(canShu[consts.LeiXing])
//	jianJie := utils.HuoQuZiFuZhi(canShu[consts.JianJie])
//	suoLueTu := utils.HuoQuZiFuZhi(canShu[consts.SuoLueTu])
//	neiRong := utils.HuoQuZiFuZhi(canShu[consts.NeiRong])
//	gengXin := utils.Time2StringNyrSfm(time.Now())
//	shuJuZhis := interface{}{}
//	//不为空才更新
//	if biaoTi != "" {
//		shuJuZhis[consts.BiaoTi] = biaoTi
//	}
//	if leiXing != "" {
//		shuJuZhis[consts.LeiXing] = leiXing
//	}
//	if jianJie != "" {
//		shuJuZhis[consts.JianJie] = jianJie
//	}
//	if suoLueTu != "" {
//		shuJuZhis[consts.SuoLueTu] = suoLueTu
//	}
//	if neiRong != "" {
//		shuJuZhis[consts.NeiRong] = neiRong
//	}
//	if gengXin != "" {
//		shuJuZhis[consts.GengXin] = gengXin
//	}
//	insertNeedCanShu := interface{}{
//		consts.ShuJuKu:   consts.Hfxrbac,
//		consts.ShuJuBiao: consts.News,
//		consts.ShuJuZhis: shuJuZhis,
//		consts.TiaoJians: interface{}{
//			consts.Id: id,
//		},
//	}
//	ret := daos.UpdateData(insertNeedCanShu)
//	return ret
//
//}
