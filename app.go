package Hongbao_system

import (
	_ "account/apis/web"
	_ "account/core/accounts"
	"github.com/zhangxinling2/infra"
	"github.com/zhangxinling2/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDataBaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&base.IrisApplicationStarter{})
	infra.Register(&infra.WebStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}