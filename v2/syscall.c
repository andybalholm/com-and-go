#include <runtime.h>
#include <cgocall.h>
void runtime·asmstdcall(void *c);

void ·cSyscall(LibCall *c) {
	runtime·cgocall(runtime·asmstdcall, c);
}
