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

func TestService_GetById(t *testing.T) {
	pgDb := docker.NewPostgres().WithTestPort(t)
	pgDb.Start(t)
	defer pgDb.Stop()

	db, _ := server.NewTestDatabase(pgDb.GetPort())

	userService := user.NewService(db)
	savedUserId, _ := userService.Save(context.Background(), user.User{ID: 1, Email: "test@test", Password: "Abc123@", Type: "USER"})

	s := feature.NewService(db)
	savedFeature, _ := s.Save(context.Background(), feature.Feature{Name: "feature-1", Responsible: savedUserId})

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "should get saved feature",
			id:      savedFeature,
			wantErr: false,
		},
		{
			name:    "should not find user by id",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetById(context.Background(), tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.id, got.ID)
		})
	}
}
