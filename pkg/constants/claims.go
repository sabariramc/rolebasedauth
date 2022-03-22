package constants

type Claim string

const (
	TenantCreate    Claim = "tenant.create"
	RoleCreate      Claim = "role.create"
	RoleUpdate      Claim = "role.update"
	RoleDelete      Claim = "role.delete"
	RoleGet         Claim = "role.get"
	RoleList        Claim = "role.list"
	UserGroupCreate Claim = "usergroup.create"
	UserGroupUpdate Claim = "usergroup.update"
	UserGroupDelete Claim = "usergroup.delete"
	UserGroupGet    Claim = "usergroup.get"
	UserGroupList   Claim = "usergroup.list"
	UserCreate      Claim = "user.create"
	UserUpdate      Claim = "user.update"
	UserDelete      Claim = "user.delete"
	UserGet         Claim = "user.get"
	UserList        Claim = "user.list"
)
