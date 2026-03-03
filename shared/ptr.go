package configure

import "github.com/hashicorp/terraform-plugin-framework/types"

func Int64ToUintPtr(v types.Int64) *int {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	i := int(v.ValueInt64())
	return &i
}

func ValueStringPointer(s types.String) *string {
	if s.IsNull() || s.IsUnknown() {
		return nil
	}
	return s.ValueStringPointer()
}

func ValueStringValue(s *string, current types.String) types.String {
	if s == nil {
		if current.IsUnknown() {
			return types.StringNull()
		}
		return current
	}
	return types.StringValue(*s)
}

func ValueBoolValue(b *bool, current types.Bool) types.Bool {
	if b == nil {
		if current.IsUnknown() {
			return types.BoolNull()
		}
		return current
	}
	return types.BoolValue(*b)
}

func ValueInt64Value(i *int, current types.Int64) types.Int64 {
	if i == nil {
		if current.IsUnknown() {
			return types.Int64Null()
		}
		return current
	}
	return types.Int64Value(int64(*i))
}
