package daosnews

import (
	"hfxrbac/consts"
	"hfxrbac/daos"
	"hfxrbac/utils"
)

func DeleteNews(canShu map[string]interface{}) map[string]interface{} {
	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
	insertNeedCanShu := map[string]interface{}{
		consts.ShuJuKu:   consts.Hfxrbac,
		consts.ShuJuBiao: consts.News,
		consts.TiaoJians: map[string]interface{}{
			consts.Id: id,
		},
	}
	ret := daos.DeleteData(insertNeedCanShu)

	return ret
}
