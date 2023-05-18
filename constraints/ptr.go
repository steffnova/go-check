package constraints

type Ptr struct {
	NilFrequency uint64
}

func PtrDefault() Ptr {
	return Ptr{
		NilFrequency: 5,
	}
}
