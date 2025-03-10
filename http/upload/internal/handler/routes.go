// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"upload/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: UploadHandler(serverCtx),
			},
		},
	)
}
