package consts

func Jg1JianXie() string {
	return Yh
}

func Jg1BiaoMing() string {
	return YongHu
}

func Jg1ZiDuan() []string {
	ret := []string{
		Yh + ZhuJian,
		Yh + MingCheng,
		Yh + MiMa,
		Yh + ShouJiHao,
		Yh + BianMa,
		Yh + YouXiang,
		Yh + NiCheng,
		Yh + QianMing,
	}
	return ret
}
