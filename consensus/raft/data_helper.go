package raft

// dataBackupHelper helps making and rotating backups from a folder.
// it will name them <folderName>.old.0, .old.1... and so on.
// when a new backup is made, the old.0 is renamed to old.1 and so on.
// when the "keep" number is reached, the oldest is always
// discarded.
type dataBackupHelper struct {
	baseDir    string
	folderName string
	keep       int
}

func newDataBackupHelper(dataFolder string, keep int) *dataBackupHelper {
	_ = "STUB: not implemented"
	return nil
}

func (dbh *dataBackupHelper) makeName(i int) string { _ = "STUB: not implemented"; return "" }

func (dbh *dataBackupHelper) listBackups() []string { _ = "STUB: not implemented"; return nil }

func (dbh *dataBackupHelper) makeBackup() error { _ = "STUB: not implemented"; return nil }

// nothing to backup

// make sure config folder exists

// list all backups in it

// remove last / oldest. Ex. if max is five, remove name.old.4

// append new backup folder. Ex, if 2 exist: add name.old.2

// increase number for all backups folders.
// If there are 3: 1->2, 0->1.
// Note in all cases the last backup in the list does not exist
// (either removed or not created, just added to this list)

// save new as name.old.0
