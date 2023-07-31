package logic

import (
	"context"
	"io"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jianxinliu/docker-starter/app/internal/svc"
	"github.com/jianxinliu/docker-starter/app/internal/types"
)

type AppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLogic {
	return &AppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppLogic) App(req *types.Request) (resp *types.Response, err error) {
	l.Logger.Infof("req: %s == [%s]", req.Name, l.svcCtx.Config.AA)
	resp = new(types.Response)
	resp.Message = req.Name

	// 读取本地文件内容，用于测试读取容器内变动的文件内容
	f, err := os.Open("etc/app-api.yaml")
	if err != nil {
		resp.Message = err.Error()
	}
	bytes, err := io.ReadAll(f)
	l.Logger.Infof("config content: %s", string(bytes))

	return
}
