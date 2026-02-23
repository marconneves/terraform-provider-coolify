package configure

func Int64ToUintPtr(v *int64) *int {
	if v == nil {
		return nil
	}
	i := int(*v)
	return &i
}
