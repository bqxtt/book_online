package tests

import (
	"context"
	"github.com/bqxtt/book_online/user/pkg/model"
	"github.com/bqxtt/book_online/user/pkg/sdk/userpb"
	"google.golang.org/grpc"
	"log"
	"reflect"
	"testing"
)

var (
	testClient clientTyp
)

func Test_All(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	ctx := context.TODO()

	testClient.UserServiceClient, testClient.ctx = client, ctx

	Test_clientTyp_register(t)
	Test_clientTyp_login(t)
	Test_clientTyp_get(t)
	Test_clientTyp_update(t)
	Test_clientTyp_get(t)
}

func Test_clientTyp_get(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal_get_01",
			args: args{
				ctx: testClient.ctx,
				id: 12,
			},
			wantErr: false,
			want: &model.User{
				UserID:    12,
				Name:      "bqx",
			},
		},
		{
			name: "normal_get_02",
			args: args{
				ctx: testClient.ctx,
				id: 1,
			},
			wantErr: false,
			want: &model.User{
				UserID:    1,
				Name:      "name",
			},
		},
		{
			name: "error_get_01",
			args: args{
				ctx: testClient.ctx,
				id: 3663,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clientTyp_login(t *testing.T) {
	type args struct {
		ctx       context.Context
		id        int64
		pwdDigest string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal_login_01",
			args: args{
				ctx: testClient.ctx,
				id: 12,
				pwdDigest: "123456",
			},
			wantErr: false,
		},
		{
			name: "normal_login_02",
			args: args{
				ctx: testClient.ctx,
				id: 1,
				pwdDigest: "password",
			},
			wantErr: false,
		},
		{
			name: "error_login_01",
			args: args{
				ctx: testClient.ctx,
				id: 1,
				pwdDigest: "123456",
			},
			wantErr: true,
		},
		{
			name: "error_login_02",
			args: args{
				ctx: testClient.ctx,
				id: 3663,
				pwdDigest: "password",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testClient.login(tt.args.ctx, tt.args.id, tt.args.pwdDigest); (err != nil) != tt.wantErr {
				t.Errorf("login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clientTyp_register(t *testing.T) {
	type args struct {
		ctx       context.Context
		id        int64
		name      string
		pwdDigest string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal_register_01",
			args: args{
				ctx: testClient.ctx,
				id: 12,
				name: "bqx",
				pwdDigest: "123456",
			},
			wantErr: false,
			want: &model.User{
				UserID:    12,
				Name:      "bqx",
			},
		},
		{
			name: "normal_register_02",
			args: args{
				ctx: testClient.ctx,
				id: 1,
				name: "name",
				pwdDigest: "password",
			},
			wantErr: false,
			want: &model.User{
				UserID:    1,
				Name:      "name",
			},
		},
		{
			name: "error_register_01",
			args: args{
				ctx: testClient.ctx,
				id: 12,
				name: "name",
				pwdDigest: "password",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.register(tt.args.ctx, tt.args.id, tt.args.name, tt.args.pwdDigest)
			if (err != nil) != tt.wantErr {
				t.Errorf("register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clientTyp_update(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   int64
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal_update_01",
			args: args{
				ctx: testClient.ctx,
				id: 12,
				name: "name",
			},
			wantErr: false,
			want: &model.User{
				UserID:    12,
				Name:      "name",
			},
		},
		{
			name: "normal_update_02",
			args: args{
				ctx: testClient.ctx,
				id: 12,
				name: "bqx",
			},
			wantErr: false,
			want: &model.User{
				UserID:    12,
				Name:      "bqx",
			},
		},
		{
			name: "normal_update_03",
			args: args{
				ctx: testClient.ctx,
				id: 1,
				name: "bqx",
			},
			wantErr: false,
			want: &model.User{
				UserID:    1,
				Name:      "bqx",
			},
		},
		{
			name: "normal_update_04",
			args: args{
				ctx: testClient.ctx,
				id: 1,
				name: "name",
			},
			wantErr: false,
			want: &model.User{
				UserID:    1,
				Name:      "name",
			},
		},
		{
			name: "error_update_01",
			args: args{
				ctx: testClient.ctx,
				id: 3663,
				name: "name",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.update(tt.args.ctx, tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
