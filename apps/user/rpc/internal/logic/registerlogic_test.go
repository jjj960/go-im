package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"goim/apps/user/rpc/internal/config"
	"goim/apps/user/rpc/internal/svc"
	"goim/apps/user/rpc/user"
	"path/filepath"
	"testing"
)

func init() {
	var c config.Config
	conf.MustLoad(filepath.Join("../../etc/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)
}

var svcCtx *svc.ServiceContext

func TestRegisterLogic_Register(t *testing.T) {
	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrint bool
		wantErr   bool
	}{
		{
			"1", args{in: &user.RegisterReq{
				Phone:    "13015694755",
				Nickname: "hys",
				Password: "123456",
				Avatar:   "dddpng.jpg",
				Sex:      1,
			}}, true, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantPrint {
				t.Log(tt.name, got)
			}
		})
	}
}
