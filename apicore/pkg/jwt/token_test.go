package jwt

import "testing"

func TestToken(t *testing.T) {
	type args struct {
		login string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test_token_1", args{"login@teste.com"}, ``, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { //checa de o jwt foi gerado
			got, err := Token(tt.args.login)
			if err != nil {
				t.Errorf("Token () error %v, got %v", err, got)
			}
		})
	}
}
