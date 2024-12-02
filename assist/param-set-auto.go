package assist

import (
	"net"
	"time"
)

// ----> auto generated(Build-ParamSet/gen-ps)

// BindBool binds bool slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindBool(info *FlagInfo, to *bool) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.BoolVar(to, info.FlagName(), info.Default.(bool), info.Usage)
	} else {
		flagSet.BoolVarP(to, info.FlagName(), info.Short, info.Default.(bool), info.Usage)
	}

	return params
}

// BindBoolSlice binds []bool slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindBoolSlice(info *FlagInfo, to *[]bool) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.BoolSliceVar(to, info.FlagName(), info.Default.([]bool), info.Usage)
	} else {
		flagSet.BoolSliceVarP(to, info.FlagName(), info.Short, info.Default.([]bool), info.Usage)
	}

	return params
}

// BindDuration binds time.Duration slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindDuration(info *FlagInfo, to *time.Duration) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.DurationVar(to, info.FlagName(), info.Default.(time.Duration), info.Usage)
	} else {
		flagSet.DurationVarP(to, info.FlagName(), info.Short, info.Default.(time.Duration), info.Usage)
	}

	return params
}

// BindDurationSlice binds []time.Duration slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindDurationSlice(info *FlagInfo, to *[]time.Duration) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.DurationSliceVar(to, info.FlagName(), info.Default.([]time.Duration), info.Usage)
	} else {
		flagSet.DurationSliceVarP(to, info.FlagName(), info.Short, info.Default.([]time.Duration), info.Usage)
	}

	return params
}

// BindEnum binds enum slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
// Note that normally the client would bind to a member of the native parameter
// set. However, since there is a discrepancy between the type of the native int
// based pseudo enum member and the equivalent acceptable string value typed by
// the user on the command line (idiomatically stored on the enum info), the
// client needs to extract the enum value from the enum info, something like this:
//
// paramSet.Native.Format = OutputFormatEnumInfo.Value()
//
// The best place to put this would be inside the PreRun/PreRunE function, assuming the
// param set and the enum info are both in scope. Actually, every int based enum
// flag, would need to have this assignment performed.
func (params *ParamSet[N]) BindEnum(info *FlagInfo, to *string) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.StringVar(to, info.FlagName(), info.Default.(string), info.Usage)
	} else {
		flagSet.StringVarP(to, info.FlagName(), info.Short, info.Default.(string), info.Usage)
	}

	return params
}

// BindFloat32 binds float32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindFloat32(info *FlagInfo, to *float32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Float32Var(to, info.FlagName(), info.Default.(float32), info.Usage)
	} else {
		flagSet.Float32VarP(to, info.FlagName(), info.Short, info.Default.(float32), info.Usage)
	}

	return params
}

// BindFloat32Slice binds []float32 slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindFloat32Slice(info *FlagInfo, to *[]float32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Float32SliceVar(to, info.FlagName(), info.Default.([]float32), info.Usage)
	} else {
		flagSet.Float32SliceVarP(to, info.FlagName(), info.Short, info.Default.([]float32), info.Usage)
	}

	return params
}

// BindFloat64 binds float64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindFloat64(info *FlagInfo, to *float64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Float64Var(to, info.FlagName(), info.Default.(float64), info.Usage)
	} else {
		flagSet.Float64VarP(to, info.FlagName(), info.Short, info.Default.(float64), info.Usage)
	}

	return params
}

// BindFloat64Slice binds []float64 slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindFloat64Slice(info *FlagInfo, to *[]float64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Float64SliceVar(to, info.FlagName(), info.Default.([]float64), info.Usage)
	} else {
		flagSet.Float64SliceVarP(to, info.FlagName(), info.Short, info.Default.([]float64), info.Usage)
	}

	return params
}

// BindInt binds int slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindInt(info *FlagInfo, to *int) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.IntVar(to, info.FlagName(), info.Default.(int), info.Usage)
	} else {
		flagSet.IntVarP(to, info.FlagName(), info.Short, info.Default.(int), info.Usage)
	}

	return params
}

// BindIntSlice binds []int slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindIntSlice(info *FlagInfo, to *[]int) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.IntSliceVar(to, info.FlagName(), info.Default.([]int), info.Usage)
	} else {
		flagSet.IntSliceVarP(to, info.FlagName(), info.Short, info.Default.([]int), info.Usage)
	}

	return params
}

// BindInt16 binds int16 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindInt16(info *FlagInfo, to *int16) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int16Var(to, info.FlagName(), info.Default.(int16), info.Usage)
	} else {
		flagSet.Int16VarP(to, info.FlagName(), info.Short, info.Default.(int16), info.Usage)
	}

	return params
}

// BindInt32 binds int32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindInt32(info *FlagInfo, to *int32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int32Var(to, info.FlagName(), info.Default.(int32), info.Usage)
	} else {
		flagSet.Int32VarP(to, info.FlagName(), info.Short, info.Default.(int32), info.Usage)
	}

	return params
}

// BindInt32Slice binds []int32 slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindInt32Slice(info *FlagInfo, to *[]int32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int32SliceVar(to, info.FlagName(), info.Default.([]int32), info.Usage)
	} else {
		flagSet.Int32SliceVarP(to, info.FlagName(), info.Short, info.Default.([]int32), info.Usage)
	}

	return params
}

// BindInt64 binds int64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindInt64(info *FlagInfo, to *int64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int64Var(to, info.FlagName(), info.Default.(int64), info.Usage)
	} else {
		flagSet.Int64VarP(to, info.FlagName(), info.Short, info.Default.(int64), info.Usage)
	}

	return params
}

// BindInt64Slice binds []int64 slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindInt64Slice(info *FlagInfo, to *[]int64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int64SliceVar(to, info.FlagName(), info.Default.([]int64), info.Usage)
	} else {
		flagSet.Int64SliceVarP(to, info.FlagName(), info.Short, info.Default.([]int64), info.Usage)
	}

	return params
}

// BindInt8 binds int8 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindInt8(info *FlagInfo, to *int8) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Int8Var(to, info.FlagName(), info.Default.(int8), info.Usage)
	} else {
		flagSet.Int8VarP(to, info.FlagName(), info.Short, info.Default.(int8), info.Usage)
	}

	return params
}

// BindIPMask binds net.IPMask slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindIPMask(info *FlagInfo, to *net.IPMask) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.IPMaskVar(to, info.FlagName(), info.Default.(net.IPMask), info.Usage)
	} else {
		flagSet.IPMaskVarP(to, info.FlagName(), info.Short, info.Default.(net.IPMask), info.Usage)
	}

	return params
}

// BindIPNet binds net.IPNet slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindIPNet(info *FlagInfo, to *net.IPNet) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.IPNetVar(to, info.FlagName(), info.Default.(net.IPNet), info.Usage)
	} else {
		flagSet.IPNetVarP(to, info.FlagName(), info.Short, info.Default.(net.IPNet), info.Usage)
	}

	return params
}

// BindString binds string slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindString(info *FlagInfo, to *string) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.StringVar(to, info.FlagName(), info.Default.(string), info.Usage)
	} else {
		flagSet.StringVarP(to, info.FlagName(), info.Short, info.Default.(string), info.Usage)
	}

	return params
}

// BindStringSlice binds []string slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindStringSlice(info *FlagInfo, to *[]string) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.StringSliceVar(to, info.FlagName(), info.Default.([]string), info.Usage)
	} else {
		flagSet.StringSliceVarP(to, info.FlagName(), info.Short, info.Default.([]string), info.Usage)
	}

	return params
}

// BindUint16 binds uint16 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindUint16(info *FlagInfo, to *uint16) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Uint16Var(to, info.FlagName(), info.Default.(uint16), info.Usage)
	} else {
		flagSet.Uint16VarP(to, info.FlagName(), info.Short, info.Default.(uint16), info.Usage)
	}

	return params
}

// BindUint32 binds uint32 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindUint32(info *FlagInfo, to *uint32) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Uint32Var(to, info.FlagName(), info.Default.(uint32), info.Usage)
	} else {
		flagSet.Uint32VarP(to, info.FlagName(), info.Short, info.Default.(uint32), info.Usage)
	}

	return params
}

// BindUint64 binds uint64 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindUint64(info *FlagInfo, to *uint64) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Uint64Var(to, info.FlagName(), info.Default.(uint64), info.Usage)
	} else {
		flagSet.Uint64VarP(to, info.FlagName(), info.Short, info.Default.(uint64), info.Usage)
	}

	return params
}

// BindUint8 binds uint8 slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindUint8(info *FlagInfo, to *uint8) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.Uint8Var(to, info.FlagName(), info.Default.(uint8), info.Usage)
	} else {
		flagSet.Uint8VarP(to, info.FlagName(), info.Short, info.Default.(uint8), info.Usage)
	}

	return params
}

// BindUint binds uint slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
func (params *ParamSet[N]) BindUint(info *FlagInfo, to *uint) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.UintVar(to, info.FlagName(), info.Default.(uint), info.Usage)
	} else {
		flagSet.UintVarP(to, info.FlagName(), info.Short, info.Default.(uint), info.Usage)
	}

	return params
}

// BindUintSlice binds []uint slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) BindUintSlice(info *FlagInfo, to *[]uint) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.UintSliceVar(to, info.FlagName(), info.Default.([]uint), info.Usage)
	} else {
		flagSet.UintSliceVarP(to, info.FlagName(), info.Short, info.Default.([]uint), info.Usage)
	}

	return params
}

// <---- end of auto generated
