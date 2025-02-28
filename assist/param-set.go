package assist

import (
	"reflect"
	"strings"

	"github.com/snivilised/sundae/internal/third/lo"
	"github.com/snivilised/sundae/locale"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// FlagInfo collates together the parameters passed into the bind methods
// The Bind methods are just a wrapper around invoking the type based methods
// on the cobra flag set in order to define flags.
type FlagInfo struct {
	// Name of flag derived from the Usage
	//
	Name string

	// Usage provides a description of the flag, the first word should eb the name
	// of the flag.
	//
	Usage string

	// Short is the 1 letter character shortcut for the flag.
	//
	Short string

	// Default is the default value for the flag if the user does not provide a
	// value.
	//
	Default any

	// AlternativeFlagSet defines the flag set to use. Allows the user to specify which flag
	// to define this flag on. By default, it is on command.Flags()
	//
	AlternativeFlagSet *pflag.FlagSet
}

func extractNameFromUsage(usage string) string {
	name := usage

	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}

	return name
}

// NewFlagInfo factory function for FlagInfo. Use this function if the flag
// is to be defined on the default flag set, ie the one on command.Flags().
func NewFlagInfo(usage, short string, def any) *FlagInfo {
	return &FlagInfo{
		Name:    extractNameFromUsage(usage),
		Usage:   usage,
		Short:   short,
		Default: def,
	}
}

// NewFlagInfoOnFlagSet factory function for FlagInfo, with an alternative flag set.
// This function need only be usd to enable defining flags on the flag set
// other than that of command.Flags(), eg command.PersistentFlags().
func NewFlagInfoOnFlagSet(usage, short string,
	def any, alternativeFlagSet *pflag.FlagSet,
) *FlagInfo {
	return &FlagInfo{
		Name:               extractNameFromUsage(usage),
		Usage:              usage,
		Short:              short,
		Default:            def,
		AlternativeFlagSet: alternativeFlagSet,
	}
}

// FlagName returns the name of the flag derived from the Usage.
func (info *FlagInfo) FlagName() string {
	return info.Name
}

// ParamSet represents a set of flags/options/positional args for a command.
// The term 'parameter set' is used really to distinguish from other established
// abstractions (flags/options/positional args, otherwise to be referred to as
// inputs). The ParamSet is used to ensure that all these inputs are collated
// into a single entity that the application can refer to as required. A command
// can have multiple parameter sets associated with it, but will probably best
// be used with a single parameter set, where inputs not provided by the end user
// are defaulted, perhaps from config. If its essential to distinguish between
// different activation scenarios (ie which set of parameters that the user provides)
// then the client can define multiple parameter sets to reflect this.
//
// The binder methods are defined explicitly for each type as 'go' does not
// allow for generic parameters defined at the method level as opposed to
// being defined on the receiver struct.
//
// The generic parameter N represents the client defined native parameter set. Eg:
//
//	type WidgetParameterSet struct {
//		 Directory string
//		 Output    string
//		 Format    OutputFormatEnum
//		 Shape     InfexionShapeEnum
//		 Concise   bool
//		 Strategy  TraversalStrategyEnum
//		 Overwrite bool
//		 Pattern   string}
//
// ... is known as the 'native' parameter set for a 'widget' command which
// would be used to instantiate ParamSet in a declaration as follows:
//
// var paramSet *ParamSet[WidgetParameterSet].
type ParamSet[N any] struct {
	// Native is the native client defined parameter set instance, which
	// must be a struct.
	//
	Native *N

	// FlagSet is the default Cobra FlagSet
	//
	FlagSet *pflag.FlagSet

	// Command is the cobra command that the parameter set is bound to.
	//
	Command *cobra.Command
}

// NewParamSet is the factory function, which creates a 'parameter set' for
// a command. Each command can have multiple command sets, reflecting the
// different ways a command can be used
//
// paramSet = NewParamSet[WidgetParameterSet](widgetCommand)
//
// The default flag set is defined, ie command.Flags(). If an alternative
// flag set is required, then the client should use
//
// The generic parameter N represents the client defined native parameter set.
func NewParamSet[N any](command *cobra.Command) (ps *ParamSet[N]) {
	ps = new(ParamSet[N])
	ps.FlagSet = command.Flags()
	ps.Native = new(N)
	ps.Command = command

	if reflect.TypeOf(*ps.Native).Kind() != reflect.Struct {
		typeName := reflect.TypeOf(*ps.Native).Name()

		panic(
			locale.NewParamSetObjectMustBeStructNativeError(command.Name(), typeName),
		)
	}

	return ps
}

// ResolveFlagSet resolves between the default flag set on the param set
// and the optional one defined on the FlagInfo. If there is no default
// flag set, then there must be one on the flag info, otherwise a panic
// will occur due dereferencing a nil pointer.
func (params *ParamSet[N]) ResolveFlagSet(info *FlagInfo) *pflag.FlagSet {
	return lo.Ternary(info.AlternativeFlagSet == nil,
		params.FlagSet, info.AlternativeFlagSet,
	)
}
