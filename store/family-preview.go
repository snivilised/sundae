package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/sundae/assist"
	"github.com/snivilised/sundae/locale"
	"github.com/spf13/pflag"
)

type PreviewParameterSet struct {
	DryRun bool
}

func (f *PreviewParameterSet) BindAll(
	parent *assist.ParamSet[PreviewParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --dry-run(D)
	//
	const (
		defaultDryRun = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.DryRunParamUsageTemplData{}),
			defaultDryRun,
			flagSet...,
		),
		&parent.Native.DryRun,
	)
}
