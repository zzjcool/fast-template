package render

import "testing"

func TestStr(t *testing.T) {
	type args struct {
		config string
		tpl    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "render",
			args: args{
				config: `foo: bar`,
				tpl: "{{ .foo }}",
			},
			want:    "bar",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Str(tt.args.config, tt.args.tpl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Str() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Str() = %v, want %v", got, tt.want)
			}
		})
	}
}
