package model

type PermissionPolicy struct {
	Model
	PolicyName string
	Remark     string
}

type PermissionAccess struct {
	Model
	AccessRule string
	Remark     string
	AccessType string
}

type PermissionAssociation struct {
	PolicyId int
	AccessId int
}
