// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package grpc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ultravioletrs/cocos-ai/manager"
)

func runEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(runReq)

		if err := req.validate(); err != nil {
			return runRes{}, err
		}

		id, err := svc.Run(ctx, req.Computation)
		if err != nil {
			return runRes{}, err
		}

		return runRes{ID: id}, nil
	}
}