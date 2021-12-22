package auth

type permissionList struct {
	ManageExecutors    string
	ManageAccounts     string
	ManageServerGroups string
	ManageTemplate     string
}

var Permissions = permissionList{
	ManageExecutors:    "executors.managae",
	ManageAccounts:     "accounts.manage",
	ManageServerGroups: "servergroups.manage",
	ManageTemplate:     "template.manage",
}
