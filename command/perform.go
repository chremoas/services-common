package command

import (
	"context"
	permsrv "github.com/chremoas/perms-srv/proto"
	"strings"
)

// Should perform.go setup its own client?
type Permissions struct {
	Client          permsrv.PermissionsClient
	PermissionsList []string
}

func (p Permissions) CanPerform(ctx context.Context, sender string) (bool, error) {
	s := strings.Split(sender, ":")
	canPerform, err := p.Client.Perform(ctx,
		&permsrv.PermissionsRequest{
			User:            s[1],
			PermissionsList: p.PermissionsList,
		})

	if err != nil {
		return false, err
	}

	return canPerform.CanPerform, nil
}
