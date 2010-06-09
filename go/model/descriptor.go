package model

type PackageName string
type VersionName string
type ConfigName string

// A Descriptor is a reference to a particular configuration, usually
// specified as part of an "include" statement.
type Descriptor struct {
	PackageName PackageName
	VersionName VersionName
	ConfigName  ConfigName
}

func (d *Descriptor) String() string {
	return string(d.PackageName) + "/" + string(d.VersionName) + ":" + string(d.ConfigName)
}
