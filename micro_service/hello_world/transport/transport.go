package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/lindep/golang_exp/micro_service/hello_world/datastruct"
	"github.com/lindep/golang_exp/micro_service/hello_world/logging"
	"github.com/lindep/golang_exp/micro_service/hello_world/service"
)

// AphService have a function
// that will have implementation that call our business logic function
type AphService interface {
	// HelloWorldService string will save the NAME input param
	HelloWorldService(context.Context, string) string
}

// aphService struct of the interface
type aphService struct{}

// ErrEmpty gen empty error string
var ErrEmpty = errors.New("empty string")

// HelloWorldService this will start call to business logic
func (aphService) HelloWorldService(_ context.Context, name string) string {
	logging.Log(fmt.Sprintf("Func HelloWorldService"))
	return callServiceHelloWorld(name)
}

// callServiceHelloWorld call business logic function
func callServiceHelloWorld(name string) string {
	messageResponse := service.HelloWorld(name)
	return messageResponse
}

func makeHelloWorldEndpoint(aph AphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.HelloWorldRequest)
		logging.Log(fmt.Sprintf("Name Request %s", req.NAME))
		v := aph.HelloWorldService(ctx, req.NAME)
		logging.Log(fmt.Sprintf("Response Final Message %s", v))
		return datastruct.HelloWorldResponse{STATUS: 200, MESSAGE: v}, nil
	}
}

func decodeHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.HelloWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// RegisterHTTPSServicesAndStartListener ...
func RegisterHTTPSServicesAndStartListener() {
	aph := aphService{}
	logging.Log(fmt.Sprintf("Setting up HelloWorldHandler"))
	HelloWorldHandler := httptransport.NewServer(
		makeHelloWorldEndpoint(aph),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/HelloWorld", HelloWorldHandler)
}
