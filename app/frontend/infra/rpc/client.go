package rpc

import (
	"context"
	"os"
	"sync"

	"github.com/baiyutang/gomall/app/cart/kitex_gen/cart/cartservice"
	"github.com/baiyutang/gomall/app/frontend/infra/mtl"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product/productcatalogservice"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/user/userservice"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client
	CartClient    cartservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
		initUserClient()
		initCartClient()
	})
}

func initProductClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		frontendutils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8881"))
	}
	p := provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(mtl.TracerProvider))
	defer p.Shutdown(context.Background())
	opts = append(opts, client.WithSuite(tracing.NewClientSuite()))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendutils.MustHandleError(err)
}

func initUserClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		frontendutils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8882"))
	}

	UserClient, err = userservice.NewClient("user", opts...)
	frontendutils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		frontendutils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8883"))
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendutils.MustHandleError(err)
}
