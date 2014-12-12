package goasp

// objects.go contains the encoding's primitives

type otype byte

var view_state_preamble = []byte{0xff, 0x01}

const (
	tUINT             otype = 0x02
	tBOOL_ARRAY             = 0x03
	tSTRING                 = 0x05
	tRGBA                   = 0x09
	tSTR_NULLTERM           = 0x0B
	tPAIR                   = 0x0f
	tTRIPLET                = 0x10
	tSTRING_ARRAY           = 0x15
	tOBJECT_CONTAINER       = 0x16
	tCONTROL_STATE          = 0x18
	tUNIT                   = 0x1b
	tSTRING2                = 0x1e
	tSTRING_REF             = 0x1f
	tUID                    = 0x24
	tEMPTY                  = 0x64
	tSTRING_EMPTY           = 0x65
	tNUM0                   = 0x66
	tTRUE                   = 0x67
	tFALSE                  = 0x68
)

// obj is a parse token subtree.
type pobject struct {
	token otype
	value interface{}
}
