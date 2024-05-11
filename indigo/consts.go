package indigo

import "reflect"

const (
	WEBARENA_INDIGO_ENDPOINT      = "WEBARENA_INDIGO_ENDPOINT"      //nolint:revive,stylecheck
	WEBARENA_INDIGO_CLIENT_ID     = "WEBARENA_INDIGO_CLIENT_ID"     //nolint:revive,stylecheck
	WEBARENA_INDIGO_CLIENT_SECRET = "WEBARENA_INDIGO_CLIENT_SECRET" //nolint:revive,stylecheck

	PathOAuthV1AccessTokens                    = "/oauth/v1/accesstokens" //nolint:gosec
	PathWebArenaIndigoV1VmSSHKey               = "/webarenaIndigo/v1/vm/sshkey"
	PathWebArenaIndigoV1VmSSHKeyActiveStatus   = "/webarenaIndigo/v1/vm/sshkey/active/status"
	PathWebArenaIndigoV1AuthCreateAPIKey       = "/webarenaIndigo/v1/auth/create/apikey"
	PathWebArenaIndigoV1AuthAPIKey             = "/webarenaIndigo/v1/auth/apikey"
	PathWebArenaIndigoV1VmInstanceTypes        = "/webarenaIndigo/v1/vm/instancetypes"
	PathWebArenaIndigoV1VmInstanceType         = "/webarenaIndigo/v1/vm/getregion"
	PathWebArenaIndigoV1VmOSList               = "/webarenaIndigo/v1/vm/oslist"
	PathWebArenaIndigoV1VmInstanceSpec         = "/webarenaIndigo/v1/vm/getinstancespec"
	PathWebArenaIndigoV1VmCreateInstance       = "/webarenaIndigo/v1/vm/createinstance"
	PathWebArenaIndigoV1VmGetInstanceList      = "/webarenaIndigo/v1/vm/getinstancelist"
	PathWebArenaIndigoV1VmInstanceStatusUpdate = "/webarenaIndigo/v1/vm/instance/statusupdate"
	PathWebArenaIndigoV1NwCreateFirewall       = "/webarenaIndigo/v1/nw/createfirewall"
	PathWebArenaIndigoV1NwGetFirewallList      = "/webarenaIndigo/v1/nw/getfirewalllist"
	PathWebArenaIndigoV1NwGetTemplate          = "/webarenaIndigo/v1/nw/gettemplate"
	PathWebArenaIndigoV1NwUpdateFirewall       = "/webarenaIndigo/v1/nw/updatefirewall"
	PathWebArenaIndigoV1NwAssign               = "/webarenaIndigo/v1/nw/assign"
	PathWebArenaIndigoV1NwDeleteFirewall       = "/webarenaIndigo/v1/nw/deletefirewall"
	PathWebArenaIndigoV1DiskTakeSnapshot       = "/webarenaIndigo/v1/disk/takesnapshot"
	PathWebArenaIndigoV1DiskSnapshotList       = "/webarenaIndigo/v1/disk/snapshotlist"
	PathWebArenaIndigoV1DiskRetakeSnapshot     = "/webarenaIndigo/v1/disk/retakesnapshot"
	PathWebArenaIndigoV1DiskRestoreSnapshot    = "/webarenaIndigo/v1/disk/restoresnapshot"
	PathWebArenaIndigoV1DiskDeleteSnapshot     = "/webarenaIndigo/v1/disk/deletesnapshot"
)

type empty struct{}

//nolint:gochecknoglobals
var pkgPath = reflect.TypeOf(empty{}).PkgPath()
