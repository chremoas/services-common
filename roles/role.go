package roles

import (
	"bytes"
	"fmt"
	rolesrv "github.com/chremoas/role-srv/proto"
	common "github.com/chremoas/services-common/command"
	"context"
)

func ListRoles(ctx context.Context, roleClient rolesrv.RolesClient, sig bool) string {
	var buffer bytes.Buffer
	var roleList = make(map[string]string)
	roles, err := roleClient.GetRoles(ctx, &rolesrv.NilMessage{})

	if err != nil {
		return common.SendFatal(err.Error())
	}

	for role := range roles.Roles {
		if roles.Roles[role].Sig == sig {
			roleList[roles.Roles[role].ShortName] = roles.Roles[role].Name
		}
	}

	if len(roleList) == 0 {
		return common.SendError("No SIGs\n")
	}

	buffer.WriteString("SIGs:\n")
	for role := range roleList {
		buffer.WriteString(fmt.Sprintf("\t%s: %s\n", role, roleList[role]))
	}

	return fmt.Sprintf("```%s```", buffer.String())
}