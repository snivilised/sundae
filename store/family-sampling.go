package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/sundae/assist"
	"github.com/snivilised/sundae/locale"
	"github.com/spf13/pflag"
)

type SamplingParameterSet struct {
	IsSampling bool
	NoFiles    uint
	NoFolders  uint
	Last       bool
}

func (f *SamplingParameterSet) BindAll(
	parent *assist.ParamSet[SamplingParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	const (
		defIsSampling = false
	)

	// --sample
	//
	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SamplingSampleUsageTemplData{}),
			defIsSampling,
			flagSet...,
		),
		&parent.Native.IsSampling,
	)

	const (
		defFSItems = uint(3)
		minFSItems = uint(1)
		maxFSItems = uint(128)
	)

	// --no-files
	//
	parent.BindUint(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SamplingNoFilesUsageTemplData{}),
			defFSItems,
			flagSet...,
		),
		&parent.Native.NoFiles,
	)

	// --no-folders
	//
	parent.BindUint(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SamplingNoFoldersUsageTemplData{}),
			defFSItems,
			flagSet...,
		),
		&parent.Native.NoFolders,
	)

	const (
		defIsLast = false
	)

	// --last
	//
	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SamplingLastUsageTemplData{}),
			defIsLast,
			flagSet...,
		),
		&parent.Native.Last,
	)
}
