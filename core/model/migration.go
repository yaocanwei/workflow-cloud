package model

import (
	db "flowpipe-server/core/db"
)

func Migrate() {
	db.WfDb.AutoMigrate(&WfArc{})
	db.WfDb.AutoMigrate(&WfCaseAssignment{})
	db.WfDb.AutoMigrate(&WfCase{})
	db.WfDb.AutoMigrate(&WfComment{})
	db.WfDb.AutoMigrate(&WfDemoTarget{})
	db.WfDb.AutoMigrate(&WfEntry{})
	db.WfDb.AutoMigrate(&WfFieldValue{})
	db.WfDb.AutoMigrate(&WfField{})
	db.WfDb.AutoMigrate(&WfForm{})
	db.WfDb.AutoMigrate(&WfGroup{})
	db.WfDb.AutoMigrate(&WfGuard{})
	db.WfDb.AutoMigrate(&WfParty{})
	db.WfDb.AutoMigrate(&WfPlace{})
	db.WfDb.AutoMigrate(&WfToken{})
	db.WfDb.AutoMigrate(&WfTranstion{})
	db.WfDb.AutoMigrate(&WfTransitionStaticAssignment{})
	db.WfDb.AutoMigrate(&WfUser{})
	db.WfDb.AutoMigrate(&WfWorkflow{})
	db.WfDb.AutoMigrate(&WfWorkitem{})
	db.WfDb.AutoMigrate(&WfWorkitemAssignment{})
	db.WfDb.AutoMigrate(&WfDepartment{})
	db.WfDb.AutoMigrate(&WfPosition{})
	db.WfDb.AutoMigrate(&WfRole{})
	db.WfDb.AutoMigrate(&WfTag{})
}
