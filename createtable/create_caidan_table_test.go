package createtable
//
//import (
//	"hfxrbac/consts"
//	"hfxrbac/daos"
//	"log"
//	"testing"
//)
//
//func TestCreateCaiDan(t *testing.T){
//	db := daos.HuoQuLianJieChi()
//
//	log.Println("db-----", db)
//	canShu := interface{}{
//		consts.ShuJuKu:   consts.Hfxrbac,
//		consts.ShuJuBiao: consts.CaiDans,
//		consts.ZhuJian:   consts.CaiDanId,
//		consts.SuoYin:    consts.CaiDanBianMa,
//		consts.ZiDuans: []interface{}{
//			interface{}{
//				consts.BianMa:   consts.CaiDanId,
//				consts.LeiXing:  consts.VARCHAR,
//				consts.ChangDu:  "50",
//				consts.MoRenZhi: "'feitian'",
//			},
//			interface{}{
//				consts.BianMa:   consts.CaiDanBianMa,
//				consts.LeiXing:  consts.VARCHAR,
//				consts.ChangDu:  "50",
//				consts.MoRenZhi: "'feitian'",
//			},
//			interface{}{
//				consts.BianMa:   consts.CaiDanMingCheng,
//				consts.LeiXing:  consts.VARCHAR,
//				consts.ChangDu:  "50",
//				consts.MoRenZhi: "'feitian'",
//			},
//			interface{}{
//				consts.BianMa:   consts.CaiDanUrl,
//				consts.LeiXing:  consts.VARCHAR,
//				consts.ChangDu:  "50",
//				consts.MoRenZhi: "'feitian'",
//			},
//			interface{}{
//				consts.BianMa:   consts.JueSeIdCaiDan,
//				consts.LeiXing:  consts.VARCHAR,
//				consts.ChangDu:  "50",
//				consts.MoRenZhi: "'feitian'",
//			},
//		},
//	}
//	daos.ChuangJianBiao(canShu)
//}