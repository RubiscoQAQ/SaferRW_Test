package grpc

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/serviceconfig"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

func main() {
	r, rcleanup := manual.GenerateAndRegisterManualResolver()
	defer rcleanup()
	r.InitialState(resolver.State{ServiceConfig: &serviceconfig.ParseResult{Err: errors.New("resolver build err")}})

	cc, err := Dial(r.Scheme()+":///test.server", WithInsecure())
	if err != nil {
		//t.Fatalf("Dial(_, _) = _, %v; want _, nil", err)
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var dummy int
	const wantMsg = "error parsing service config"
	const wantCode = codes.Unavailable
	if err := cc.Invoke(ctx, "/foo/bar", &dummy, &dummy); status.Code(err) != wantCode || !strings.Contains(status.Convert(err).Message(), wantMsg) {
		//t.Fatalf("cc.Invoke(_, _, _, _) = %v; want status.Code()==%v, status.Message() contains %q", err, wantCode, wantMsg)
	}
}
