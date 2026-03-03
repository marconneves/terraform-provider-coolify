package configure

import (
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

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

func DiffString(plan, state types.String) *string {
	if plan.Equal(state) {
		return nil
	}
	if plan.IsNull() || plan.IsUnknown() {
		return nil
	}
	return plan.ValueStringPointer()
}

func DiffBase64String(plan, state types.String) *string {
	if plan.Equal(state) {
		return nil
	}
	if plan.IsNull() || plan.IsUnknown() {
		return nil
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(plan.ValueString()))
	return &encoded
}

func DiffBool(plan, state types.Bool) *bool {
	if plan.Equal(state) {
		return nil
	}
	if plan.IsNull() || plan.IsUnknown() {
		return nil
	}
	b := plan.ValueBool()
	return &b
}

func DiffInt64(plan, state types.Int64) *int {
	if plan.Equal(state) {
		return nil
	}
	if plan.IsNull() || plan.IsUnknown() {
		return nil
	}
	i := int(plan.ValueInt64())
	return &i
}

func Base64String(s *string, current types.String) types.String {
	if s == nil {
		if current.IsUnknown() {
			return types.StringNull()
		}
		return current
	}

	// Try to decode. If it fails, it's probably already plain text or not encoded at all.
	decoded, err := base64.StdEncoding.DecodeString(*s)
	if err == nil {
		return types.StringValue(string(decoded))
	}

	return types.StringValue(*s)
}

func Base64EncodeString(s types.String) *string {
	if s.IsNull() || s.IsUnknown() {
		return nil
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(s.ValueString()))
	return &encoded
}
