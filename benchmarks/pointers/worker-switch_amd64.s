#include "textflag.h"

TEXT ·storePointer(SB),NOSPLIT,$0-16
	MOVQ	destPtr+0(FP), BP
	MOVQ	newPtr +8(FP), AX
	XCHGQ	AX, 0(BP)
	RET
