package consts

const (
	Jg3YeWus = "Jg3YeWus" //业务结构
	Yw       = "Yw"
)
const (
	YwZhuJian   = Yw + ZhuJian
	YwMingCheng = Yw + MingCheng
	YwBianMa    = Yw + BianMa
	YwJiBie     = Yw + JiBie   //接口级别还是数据级别，如果是数据级别必须要有一个用YongHuZhuJian+当前表名的字段进行匹配，如果是接口级别则不需要。
	YwLianJie   = Yw + LianJie //这个主要是为了方便网关的，实际上所有的业务只会走一个handler
	YwZhuBiao   = Yw + ZhuBiao //业务主表就是即将操作的主要的表，这个表中如果确定了是一个数据级别的表那么必须有一个字段是（YongHuZhuJian+当前表名），比如操作活动表，更改标题时必须要活动创建人才可以更改，这时候主表就是活动表，其他表如果需要操作则从活动表进行关联，
)
