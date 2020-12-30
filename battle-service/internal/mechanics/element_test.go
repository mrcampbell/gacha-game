package mechanics

import (
	"testing"

	"github.com/mrcampbell/gacha-game/battle-service/internal/app"
)

func TestElementalAdvantageFactor(t *testing.T) {
	type args struct {
		attacker app.Element
		defender app.Element
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "fire vs fire = normal",
			args: args{
				attacker: app.ElementFire,
				defender: app.ElementFire,
			},
			want:    1.0,
			wantErr: false,
		},
		{
			name: "fire vs grass = effective",
			args: args{
				attacker: app.ElementFire,
				defender: app.ElementGrass,
			},
			want:    2.0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ElementalAdvantageFactor(tt.args.attacker, tt.args.defender)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElementalAdvantageFactor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ElementalAdvantageFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
