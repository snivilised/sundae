package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/sundae/assist"
	"github.com/snivilised/sundae/locale"
	"github.com/spf13/pflag"
)

type I18nParameterSet struct {
	Language string
}

func (f *I18nParameterSet) BindAll(
	parent *assist.ParamSet[I18nParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --language
	//
	const (
		defaultLanguage = ""
	)

	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.LanguageParamUsageTemplData{}),
			defaultLanguage,
			flagSet...,
		),
		&parent.Native.Language,
	)
}
