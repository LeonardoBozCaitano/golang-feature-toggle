package feature_test

import (
	"context"
	"testing"

	"github.com/schedule-api/pkg/docker"
	"github.com/schedule-api/pkg/feature"
	"github.com/schedule-api/pkg/server"
	"github.com/schedule-api/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestService_Save(t *testing.T) {
	pgDb := docker.NewPostgres().WithTestPort(t)
	pgDb.Start(t)
	defer pgDb.Stop()

	db, _ := server.NewTestDatabase(pgDb.GetPort())

	userService := user.NewService(db)
	savedUserId, _ := userService.Save(context.Background(), user.User{Email: "test4@test.com", Password: "Abc123@", Type: "USER"})
	tests := []struct {
		name         string
		want         int
		toSave       feature.Feature
		wantError    bool
		errorMessage string
	}{
		{
			name:      "should save",
			want:      1,
			toSave:    feature.Feature{Name: "feature-1", Responsible: savedUserId},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := feature.NewService(db)
			got, err := s.Save(context.Background(), tt.toSave)

			if !tt.wantError {
				assert.Equal(t, tt.want, got)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.errorMessage, err.Error())
			}
		})
	}
}
