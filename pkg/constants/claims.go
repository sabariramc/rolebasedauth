package constants

type Claim string

const (
	TenantCreate    Claim = "tenant.create"
	TenantUpdate    Claim = "tenant.update"
	TenantDelete    Claim = "tenant.delete"
	TenantList      Claim = "tenant.list"
	TenantGet       Claim = "tenant.get"
	TenantSearch    Claim = "tenant.search"
	UserRoleCreate  Claim = "userrole.create"
	UserRoleUpdate  Claim = "userrole.update"
	UserRoleDelete  Claim = "userrole.delete"
	UserRoleGet     Claim = "userrole.get"
	UserRoleList    Claim = "userrole.list"
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

var RoleTenantAdmin = []Claim{
	TenantUpdate,
	TenantSearch,
	TenantGet,
	UserRoleCreate,
	UserRoleUpdate,
	UserRoleDelete,
	UserRoleGet,
	UserRoleList,
	UserGroupCreate,
	UserGroupUpdate,
	UserGroupDelete,
	UserGroupGet,
	UserGroupList,
	UserCreate,
	UserUpdate,
	UserDelete,
	UserGet,
	UserList,
}
