package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/sundae/assist"
	"github.com/snivilised/sundae/locale"
	"github.com/spf13/pflag"
)

type ProfileParameterSet struct {
	Profile string
	Scheme  string
}

func (f *ProfileParameterSet) BindAll(
	parent *assist.ParamSet[ProfileParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --profile(P)
	//
	const (
		defaultProfile = ""
	)

	// should match: `^[\w-]+$`,
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.ProfileParamUsageTemplData{}),
			defaultProfile,
			flagSet...,
		),
		&parent.Native.Profile,
	)

	// -- scheme(S)
	//
	const (
		defaultScheme = ""
	)

	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SchemeParamUsageTemplData{}),
			defaultScheme,
			flagSet...,
		),
		&parent.Native.Scheme,
	)
}
