package u_http_client

import (
	"regexp"
	"testing"

	"github.com/tikivn/ultrago/u_prometheus"
)

func Test_httpExecutor_cleanUpPath(t *testing.T) {
	prometheusHttpConfig := map[*regexp.Regexp]string{
		regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})\\/"): "/<id>/",
		regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})"):    "/<id>",
		regexp.MustCompile("\\/ts\\/tikimsp\\/[a-z0-9A-Z/]+\\.png"):                                            "/ts/tikimsp/<path.png>",
		regexp.MustCompile("\\/ts\\/tikimsp\\/[a-z0-9A-Z/]+\\.(jpg|JPG)"):                                      "/ts/tikimsp/<path.jpg>",
		regexp.MustCompile("\\/v2\\/banners\\/[0-9]+"):                                                         "/v2/banners/<id>",
		regexp.MustCompile("\\/v2\\/banner_groups\\/[0-9]+"):                                                   "/v2/banner_groups/<id>",
		regexp.MustCompile("\\/v1\\/customers\\/[0-9]+"):                                                       "/v1/customers/<id>",
	}

	type fields struct {
		prometheusHttpConfig *u_prometheus.HttpConfig
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				prometheusHttpConfig: &u_prometheus.HttpConfig{
					PathCleanUpMap: prometheusHttpConfig,
				},
			},
			args: args{
				path: "/v2/banners/104407",
			},
			want: "/v2/banners/<id>",
		},
		{
			name: "Test case 2",
			fields: fields{
				prometheusHttpConfig: &u_prometheus.HttpConfig{
					PathCleanUpMap: prometheusHttpConfig,
				},
			},
			args: args{
				path: "/v2/banners/104407/upload_banners",
			},
			want: "/v2/banners/<id>/upload_banners",
		},
		{
			name: "Test case 3",
			fields: fields{
				prometheusHttpConfig: &u_prometheus.HttpConfig{
					PathCleanUpMap: prometheusHttpConfig,
				},
			},
			args: args{
				path: "/v2/banner_groups/46951/add_banners",
			},
			want: "/v2/banner_groups/<id>/add_banners",
		},
		{
			name: "Test case 4",
			fields: fields{
				prometheusHttpConfig: &u_prometheus.HttpConfig{
					PathCleanUpMap: prometheusHttpConfig,
				},
			},
			args: args{
				path: "/ts/tikimsp/ff/e3/dc/ba2f7f4eea078dd3af327179c924c622.jpg",
			},
			want: "/ts/tikimsp/<path.jpg>",
		},
		{
			name: "Test case 5",
			fields: fields{
				prometheusHttpConfig: &u_prometheus.HttpConfig{
					PathCleanUpMap: prometheusHttpConfig,
				},
			},
			args: args{
				path: "/ts/tikimsp/ff/e3/dc/ba2f7f4eea078dd3af327179c924c622.JPG",
			},
			want: "/ts/tikimsp/<path.jpg>",
		},
		{
			name: "Test case 6",
			fields: fields{
				prometheusHttpConfig: &u_prometheus.HttpConfig{
					PathCleanUpMap: prometheusHttpConfig,
				},
			},
			args: args{
				path: "/ts/tikimsp/ff/44/a3/dd252086d156d95d907c8bb4da67f02c.png",
			},
			want: "/ts/tikimsp/<path.png>",
		},
		{
			name: "Test case 7",
			fields: fields{
				prometheusHttpConfig: &u_prometheus.HttpConfig{
					PathCleanUpMap: prometheusHttpConfig,
				},
			},
			args: args{
				path: "/v1/customers/26815664",
			},
			want: "/v1/customers/<id>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &httpExecutor{
				prometheusHttpConfig: tt.fields.prometheusHttpConfig,
			}
			if got := c.cleanUpPath(tt.args.path); got != tt.want {
				t.Errorf("httpExecutor.cleanUpPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
