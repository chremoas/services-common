package command

import (
	permsrv "github.com/chremoas/perms-srv/proto"
	"strings"
	"context"
)

type Permissions struct {
	Client permsrv.PermissionsClient
}

func (p Permissions) CanPerform(ctx context.Context, sender string, perms []string) (bool, error) {
	s := strings.Split(sender, ":")
	canPerform, err := p.Client.Perform(ctx,
		&permsrv.PermissionsRequest{
			User:            s[1],
			PermissionsList: perms,
		})

	if err != nil {
		return false, err
	}

	return canPerform.CanPerform, nil
}
