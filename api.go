package cu

// #include <cuda.h>
import "C"
import "unsafe"

// This file was generated by the genlib program. DO NOT EDIT

func GetDevice(ordinal int) (device Device, err error) {
	Cordinal := C.int(ordinal)
	var Cdevice C.CUdevice
	err = result(C.cuDeviceGet(&Cdevice, Cordinal))
	device = Device(Cdevice)

	return
}

func NumDevices() (count int, err error) {
	var Ccount C.int
	err = result(C.cuDeviceGetCount(&Ccount))
	count = int(Ccount)

	return
}

func (dev Device) TotalMem() (bytes int64, err error) {
	Cdev := C.CUdevice(dev)
	var Cbytes C.size_t
	err = result(C.cuDeviceTotalMem(&Cbytes, Cdev))
	bytes = int64(Cbytes)

	return
}

func (dev Device) Attribute(attrib DeviceAttribute) (pi int, err error) {
	Cdev := C.CUdevice(dev)
	Cattrib := C.CUdevice_attribute(attrib)
	var Cpi C.int
	err = result(C.cuDeviceGetAttribute(&Cpi, Cattrib, Cdev))
	pi = int(Cpi)

	return
}

func (dev Device) ReleasePrimaryCtx() (err error) {
	Cdev := C.CUdevice(dev)
	return result(C.cuDevicePrimaryCtxRelease(Cdev))
}

func (dev Device) SetPrimaryCtxFlags(flags ContextFlags) (err error) {
	Cdev := C.CUdevice(dev)
	Cflags := C.uint(flags)
	return result(C.cuDevicePrimaryCtxSetFlags(Cdev, Cflags))
}

func (dev Device) PrimaryCtxState() (flags ContextFlags, active int, err error) {
	Cdev := C.CUdevice(dev)
	var Cflags C.uint
	var Cactive C.int
	err = result(C.cuDevicePrimaryCtxGetState(Cdev, &Cflags, &Cactive))
	flags = ContextFlags(Cflags)

	active = int(Cactive)

	return
}

func (dev Device) ResetPrimaryCtx() (err error) {
	Cdev := C.CUdevice(dev)
	return result(C.cuDevicePrimaryCtxReset(Cdev))
}

func PushCurrentCtx(ctx CUContext) (err error) {
	Cctx := ctx.c()
	return result(C.cuCtxPushCurrent(Cctx))
}

func PopCurrentCtx() (pctx CUContext, err error) {
	var Cpctx C.CUcontext
	err = result(C.cuCtxPopCurrent(&Cpctx))
	pctx = makeContext(Cpctx)
	return
}

func SetCurrentContext(ctx CUContext) (err error) {
	Cctx := ctx.c()
	return result(C.cuCtxSetCurrent(Cctx))
}

func CurrentContext() (pctx CUContext, err error) {
	var Cpctx C.CUcontext
	err = result(C.cuCtxGetCurrent(&Cpctx))
	pctx = makeContext(Cpctx)
	return
}

func CurrentDevice() (device Device, err error) {
	var Cdevice C.CUdevice
	err = result(C.cuCtxGetDevice(&Cdevice))
	device = Device(Cdevice)

	return
}

func CurrentFlags() (flags ContextFlags, err error) {
	var Cflags C.uint
	err = result(C.cuCtxGetFlags(&Cflags))
	flags = ContextFlags(Cflags)

	return
}

func Synchronize() (err error) {
	return result(C.cuCtxSynchronize())
}

func SetLimit(limit Limit, value int64) (err error) {
	Climit := C.CUlimit(limit)
	Cvalue := C.size_t(value)
	return result(C.cuCtxSetLimit(Climit, Cvalue))
}

func Limits(limit Limit) (pvalue int64, err error) {
	Climit := C.CUlimit(limit)
	var Cpvalue C.size_t
	err = result(C.cuCtxGetLimit(&Cpvalue, Climit))
	pvalue = int64(Cpvalue)

	return
}

func CurrentCacheConfig() (pconfig FuncCacheConfig, err error) {
	var Cpconfig C.CUfunc_cache
	err = result(C.cuCtxGetCacheConfig(&Cpconfig))
	pconfig = FuncCacheConfig(Cpconfig)

	return
}

func SetCurrentCacheConfig(config FuncCacheConfig) (err error) {
	Cconfig := C.CUfunc_cache(config)
	return result(C.cuCtxSetCacheConfig(Cconfig))
}

func SharedMemConfig() (pConfig SharedConfig, err error) {
	var CpConfig C.CUsharedconfig
	err = result(C.cuCtxGetSharedMemConfig(&CpConfig))
	pConfig = SharedConfig(CpConfig)

	return
}

func SetSharedMemConfig(config SharedConfig) (err error) {
	Cconfig := C.CUsharedconfig(config)
	return result(C.cuCtxSetSharedMemConfig(Cconfig))
}

func (ctx CUContext) APIVersion() (version uint, err error) {
	Cctx := ctx.c()
	var Cversion C.uint
	err = result(C.cuCtxGetApiVersion(Cctx, &Cversion))
	version = uint(Cversion)

	return
}

func StreamPriorityRange() (leastPriority int, greatestPriority int, err error) {
	var CleastPriority C.int
	var CgreatestPriority C.int
	err = result(C.cuCtxGetStreamPriorityRange(&CleastPriority, &CgreatestPriority))
	leastPriority = int(CleastPriority)

	greatestPriority = int(CgreatestPriority)

	return
}

func Unload(hmod Module) (err error) {
	Chmod := hmod.c()
	return result(C.cuModuleUnload(Chmod))
}

func MemInfo() (free int64, total int64, err error) {
	var Cfree C.size_t
	var Ctotal C.size_t
	err = result(C.cuMemGetInfo(&Cfree, &Ctotal))
	free = int64(Cfree)

	total = int64(Ctotal)

	return
}

func MemAlloc(bytesize int64) (dptr DevicePtr, err error) {
	Cbytesize := C.size_t(bytesize)
	var Cdptr C.CUdeviceptr
	err = result(C.cuMemAlloc(&Cdptr, Cbytesize))
	dptr = DevicePtr(Cdptr)

	return
}

func MemAllocPitch(WidthInBytes int64, Height int64, ElementSizeBytes uint) (dptr DevicePtr, pPitch int64, err error) {
	CWidthInBytes := C.size_t(WidthInBytes)
	CHeight := C.size_t(Height)
	CElementSizeBytes := C.uint(ElementSizeBytes)
	var Cdptr C.CUdeviceptr
	var CpPitch C.size_t
	err = result(C.cuMemAllocPitch(&Cdptr, &CpPitch, CWidthInBytes, CHeight, CElementSizeBytes))
	dptr = DevicePtr(Cdptr)

	pPitch = int64(CpPitch)

	return
}

func MemFree(dptr DevicePtr) (err error) {
	Cdptr := C.CUdeviceptr(dptr)
	return result(C.cuMemFree(Cdptr))
}

func MemFreeHost(p unsafe.Pointer) (err error) {
	Cp := p
	return result(C.cuMemFreeHost(Cp))
}

func MemAllocManaged(bytesize int64, flags MemAttachFlags) (dptr DevicePtr, err error) {
	Cbytesize := C.size_t(bytesize)
	Cflags := C.uint(flags)
	var Cdptr C.CUdeviceptr
	err = result(C.cuMemAllocManaged(&Cdptr, Cbytesize, Cflags))
	dptr = DevicePtr(Cdptr)

	return
}

func Memcpy(dst DevicePtr, src DevicePtr, ByteCount int64) (err error) {
	Cdst := C.CUdeviceptr(dst)
	Csrc := C.CUdeviceptr(src)
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpy(Cdst, Csrc, CByteCount))
}

func MemcpyPeer(dstDevice DevicePtr, dstContext CUContext, srcDevice DevicePtr, srcContext CUContext, ByteCount int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CdstContext := dstContext.c()
	CsrcDevice := C.CUdeviceptr(srcDevice)
	CsrcContext := srcContext.c()
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyPeer(CdstDevice, CdstContext, CsrcDevice, CsrcContext, CByteCount))
}

func MemcpyHtoD(dstDevice DevicePtr, srcHost unsafe.Pointer, ByteCount int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CsrcHost := srcHost
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyHtoD(CdstDevice, CsrcHost, CByteCount))
}

func MemcpyDtoH(dstHost unsafe.Pointer, srcDevice DevicePtr, ByteCount int64) (err error) {
	CdstHost := dstHost
	CsrcDevice := C.CUdeviceptr(srcDevice)
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyDtoH(CdstHost, CsrcDevice, CByteCount))
}

func MemcpyDtoD(dstDevice DevicePtr, srcDevice DevicePtr, ByteCount int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CsrcDevice := C.CUdeviceptr(srcDevice)
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyDtoD(CdstDevice, CsrcDevice, CByteCount))
}

func MemcpyDtoA(dstArray Array, dstOffset int64, srcDevice DevicePtr, ByteCount int64) (err error) {
	CdstArray := dstArray.c()
	CdstOffset := C.size_t(dstOffset)
	CsrcDevice := C.CUdeviceptr(srcDevice)
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyDtoA(CdstArray, CdstOffset, CsrcDevice, CByteCount))
}

func MemcpyAtoD(dstDevice DevicePtr, srcArray Array, srcOffset int64, ByteCount int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CsrcArray := srcArray.c()
	CsrcOffset := C.size_t(srcOffset)
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyAtoD(CdstDevice, CsrcArray, CsrcOffset, CByteCount))
}

func MemcpyHtoA(dstArray Array, dstOffset int64, srcHost unsafe.Pointer, ByteCount int64) (err error) {
	CdstArray := dstArray.c()
	CdstOffset := C.size_t(dstOffset)
	CsrcHost := srcHost
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyHtoA(CdstArray, CdstOffset, CsrcHost, CByteCount))
}

func MemcpyAtoH(dstHost unsafe.Pointer, srcArray Array, srcOffset int64, ByteCount int64) (err error) {
	CdstHost := dstHost
	CsrcArray := srcArray.c()
	CsrcOffset := C.size_t(srcOffset)
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyAtoH(CdstHost, CsrcArray, CsrcOffset, CByteCount))
}

func MemcpyAtoA(dstArray Array, dstOffset int64, srcArray Array, srcOffset int64, ByteCount int64) (err error) {
	CdstArray := dstArray.c()
	CdstOffset := C.size_t(dstOffset)
	CsrcArray := srcArray.c()
	CsrcOffset := C.size_t(srcOffset)
	CByteCount := C.size_t(ByteCount)
	return result(C.cuMemcpyAtoA(CdstArray, CdstOffset, CsrcArray, CsrcOffset, CByteCount))
}

func Memcpy2D(pCopy Memcpy2dParam) (err error) {
	CpCopy := pCopy.c()
	return result(C.cuMemcpy2D(CpCopy))
}

func Memcpy2DUnaligned(pCopy Memcpy2dParam) (err error) {
	CpCopy := pCopy.c()
	return result(C.cuMemcpy2DUnaligned(CpCopy))
}

func Memcpy3D(pCopy Memcpy3dParam) (err error) {
	CpCopy := pCopy.c()
	return result(C.cuMemcpy3D(CpCopy))
}

func Memcpy3DPeer(pCopy Memcpy3dPeerParam) (err error) {
	CpCopy := pCopy.c()
	return result(C.cuMemcpy3DPeer(CpCopy))
}

func MemcpyAsync(dst DevicePtr, src DevicePtr, ByteCount int64, hStream Stream) (err error) {
	Cdst := C.CUdeviceptr(dst)
	Csrc := C.CUdeviceptr(src)
	CByteCount := C.size_t(ByteCount)
	ChStream := hStream.c()
	return result(C.cuMemcpyAsync(Cdst, Csrc, CByteCount, ChStream))
}

func MemcpyPeerAsync(dstDevice DevicePtr, dstContext CUContext, srcDevice DevicePtr, srcContext CUContext, ByteCount int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CdstContext := dstContext.c()
	CsrcDevice := C.CUdeviceptr(srcDevice)
	CsrcContext := srcContext.c()
	CByteCount := C.size_t(ByteCount)
	ChStream := hStream.c()
	return result(C.cuMemcpyPeerAsync(CdstDevice, CdstContext, CsrcDevice, CsrcContext, CByteCount, ChStream))
}

func MemcpyHtoDAsync(dstDevice DevicePtr, srcHost unsafe.Pointer, ByteCount int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CsrcHost := srcHost
	CByteCount := C.size_t(ByteCount)
	ChStream := hStream.c()
	return result(C.cuMemcpyHtoDAsync(CdstDevice, CsrcHost, CByteCount, ChStream))
}

func MemcpyDtoHAsync(dstHost unsafe.Pointer, srcDevice DevicePtr, ByteCount int64, hStream Stream) (err error) {
	CdstHost := dstHost
	CsrcDevice := C.CUdeviceptr(srcDevice)
	CByteCount := C.size_t(ByteCount)
	ChStream := hStream.c()
	return result(C.cuMemcpyDtoHAsync(CdstHost, CsrcDevice, CByteCount, ChStream))
}

func MemcpyDtoDAsync(dstDevice DevicePtr, srcDevice DevicePtr, ByteCount int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CsrcDevice := C.CUdeviceptr(srcDevice)
	CByteCount := C.size_t(ByteCount)
	ChStream := hStream.c()
	return result(C.cuMemcpyDtoDAsync(CdstDevice, CsrcDevice, CByteCount, ChStream))
}

func MemcpyHtoAAsync(dstArray Array, dstOffset int64, srcHost unsafe.Pointer, ByteCount int64, hStream Stream) (err error) {
	CdstArray := dstArray.c()
	CdstOffset := C.size_t(dstOffset)
	CsrcHost := srcHost
	CByteCount := C.size_t(ByteCount)
	ChStream := hStream.c()
	return result(C.cuMemcpyHtoAAsync(CdstArray, CdstOffset, CsrcHost, CByteCount, ChStream))
}

func MemcpyAtoHAsync(dstHost unsafe.Pointer, srcArray Array, srcOffset int64, ByteCount int64, hStream Stream) (err error) {
	CdstHost := dstHost
	CsrcArray := srcArray.c()
	CsrcOffset := C.size_t(srcOffset)
	CByteCount := C.size_t(ByteCount)
	ChStream := hStream.c()
	return result(C.cuMemcpyAtoHAsync(CdstHost, CsrcArray, CsrcOffset, CByteCount, ChStream))
}

func Memcpy2DAsync(pCopy Memcpy2dParam, hStream Stream) (err error) {
	CpCopy := pCopy.c()
	ChStream := hStream.c()
	return result(C.cuMemcpy2DAsync(CpCopy, ChStream))
}

func Memcpy3DAsync(pCopy Memcpy3dParam, hStream Stream) (err error) {
	CpCopy := pCopy.c()
	ChStream := hStream.c()
	return result(C.cuMemcpy3DAsync(CpCopy, ChStream))
}

func Memcpy3DPeerAsync(pCopy Memcpy3dPeerParam, hStream Stream) (err error) {
	CpCopy := pCopy.c()
	ChStream := hStream.c()
	return result(C.cuMemcpy3DPeerAsync(CpCopy, ChStream))
}

func MemsetD8(dstDevice DevicePtr, uc byte, N int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	Cuc := C.uchar(uc)
	CN := C.size_t(N)
	return result(C.cuMemsetD8(CdstDevice, Cuc, CN))
}

func MemsetD16(dstDevice DevicePtr, us uint16, N int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	Cus := C.ushort(us)
	CN := C.size_t(N)
	return result(C.cuMemsetD16(CdstDevice, Cus, CN))
}

func MemsetD32(dstDevice DevicePtr, ui uint32, N int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	Cui := C.uint(ui)
	CN := C.size_t(N)
	return result(C.cuMemsetD32(CdstDevice, Cui, CN))
}

func MemsetD2D8(dstDevice DevicePtr, dstPitch int64, uc byte, Width int64, Height int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CdstPitch := C.size_t(dstPitch)
	Cuc := C.uchar(uc)
	CWidth := C.size_t(Width)
	CHeight := C.size_t(Height)
	return result(C.cuMemsetD2D8(CdstDevice, CdstPitch, Cuc, CWidth, CHeight))
}

func MemsetD2D16(dstDevice DevicePtr, dstPitch int64, us uint16, Width int64, Height int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CdstPitch := C.size_t(dstPitch)
	Cus := C.ushort(us)
	CWidth := C.size_t(Width)
	CHeight := C.size_t(Height)
	return result(C.cuMemsetD2D16(CdstDevice, CdstPitch, Cus, CWidth, CHeight))
}

func MemsetD2D32(dstDevice DevicePtr, dstPitch int64, ui uint, Width int64, Height int64) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CdstPitch := C.size_t(dstPitch)
	Cui := C.uint(ui)
	CWidth := C.size_t(Width)
	CHeight := C.size_t(Height)
	return result(C.cuMemsetD2D32(CdstDevice, CdstPitch, Cui, CWidth, CHeight))
}

func MemsetD8Async(dstDevice DevicePtr, uc byte, N int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	Cuc := C.uchar(uc)
	CN := C.size_t(N)
	ChStream := hStream.c()
	return result(C.cuMemsetD8Async(CdstDevice, Cuc, CN, ChStream))
}

func MemsetD16Async(dstDevice DevicePtr, us uint16, N int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	Cus := C.ushort(us)
	CN := C.size_t(N)
	ChStream := hStream.c()
	return result(C.cuMemsetD16Async(CdstDevice, Cus, CN, ChStream))
}

func MemsetD32Async(dstDevice DevicePtr, ui uint, N int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	Cui := C.uint(ui)
	CN := C.size_t(N)
	ChStream := hStream.c()
	return result(C.cuMemsetD32Async(CdstDevice, Cui, CN, ChStream))
}

func MemsetD2D8Async(dstDevice DevicePtr, dstPitch int64, uc byte, Width int64, Height int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CdstPitch := C.size_t(dstPitch)
	Cuc := C.uchar(uc)
	CWidth := C.size_t(Width)
	CHeight := C.size_t(Height)
	ChStream := hStream.c()
	return result(C.cuMemsetD2D8Async(CdstDevice, CdstPitch, Cuc, CWidth, CHeight, ChStream))
}

func MemsetD2D16Async(dstDevice DevicePtr, dstPitch int64, us uint16, Width int64, Height int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CdstPitch := C.size_t(dstPitch)
	Cus := C.ushort(us)
	CWidth := C.size_t(Width)
	CHeight := C.size_t(Height)
	ChStream := hStream.c()
	return result(C.cuMemsetD2D16Async(CdstDevice, CdstPitch, Cus, CWidth, CHeight, ChStream))
}

func MemsetD2D32Async(dstDevice DevicePtr, dstPitch int64, ui uint, Width int64, Height int64, hStream Stream) (err error) {
	CdstDevice := C.CUdeviceptr(dstDevice)
	CdstPitch := C.size_t(dstPitch)
	Cui := C.uint(ui)
	CWidth := C.size_t(Width)
	CHeight := C.size_t(Height)
	ChStream := hStream.c()
	return result(C.cuMemsetD2D32Async(CdstDevice, CdstPitch, Cui, CWidth, CHeight, ChStream))
}

func (hArray Array) Descriptor() (pArrayDescriptor ArrayDesc, err error) {
	ChArray := hArray.c()
	var CpArrayDescriptor C.CUDA_ARRAY_DESCRIPTOR
	err = result(C.cuArrayGetDescriptor(&CpArrayDescriptor, ChArray))
	pArrayDescriptor = goArrayDesc(&CpArrayDescriptor)
	return
}

func DestroyArray(hArray Array) (err error) {
	ChArray := hArray.c()
	return result(C.cuArrayDestroy(ChArray))
}

func (hArray Array) Descriptor3() (pArrayDescriptor Array3Desc, err error) {
	ChArray := hArray.c()
	var CpArrayDescriptor C.CUDA_ARRAY3D_DESCRIPTOR
	err = result(C.cuArray3DGetDescriptor(&CpArrayDescriptor, ChArray))
	pArrayDescriptor = goArray3Desc(&CpArrayDescriptor)
	return
}

func (hStream Stream) Priority() (priority int, err error) {
	ChStream := hStream.c()
	var Cpriority C.int
	err = result(C.cuStreamGetPriority(ChStream, &Cpriority))
	priority = int(Cpriority)

	return
}

func (hStream Stream) Flags() (flags StreamFlags, err error) {
	ChStream := hStream.c()
	var Cflags C.uint
	err = result(C.cuStreamGetFlags(ChStream, &Cflags))
	flags = StreamFlags(Cflags)

	return
}

func (hStream Stream) Wait(hEvent Event, Flags uint) (err error) {
	ChStream := hStream.c()
	ChEvent := hEvent.c()
	CFlags := C.uint(Flags)
	return result(C.cuStreamWaitEvent(ChStream, ChEvent, CFlags))
}

func (hStream Stream) AttachMemAsync(dptr DevicePtr, length int64, flags uint) (err error) {
	ChStream := hStream.c()
	Cdptr := C.CUdeviceptr(dptr)
	Clength := C.size_t(length)
	Cflags := C.uint(flags)
	return result(C.cuStreamAttachMemAsync(ChStream, Cdptr, Clength, Cflags))
}

func (hStream Stream) Query() (err error) {
	ChStream := hStream.c()
	return result(C.cuStreamQuery(ChStream))
}

func (hStream Stream) Synchronize() (err error) {
	ChStream := hStream.c()
	return result(C.cuStreamSynchronize(ChStream))
}

func (hEvent Event) Record(hStream Stream) (err error) {
	ChEvent := hEvent.c()
	ChStream := hStream.c()
	return result(C.cuEventRecord(ChEvent, ChStream))
}

func (hEvent Event) Query() (err error) {
	ChEvent := hEvent.c()
	return result(C.cuEventQuery(ChEvent))
}

func (hEvent Event) Synchronize() (err error) {
	ChEvent := hEvent.c()
	return result(C.cuEventSynchronize(ChEvent))
}

func (hStart Event) Elapsed(hEnd Event) (pMilliseconds float64, err error) {
	ChStart := hStart.c()
	ChEnd := hEnd.c()
	var CpMilliseconds C.float
	err = result(C.cuEventElapsedTime(&CpMilliseconds, ChStart, ChEnd))
	pMilliseconds = float64(CpMilliseconds)

	return
}

func (stream Stream) WaitOnValue32(addr DevicePtr, value uint32, flags uint) (err error) {
	Cstream := stream.c()
	Caddr := C.CUdeviceptr(addr)
	Cvalue := C.cuuint32_t(value)
	Cflags := C.uint(flags)
	return result(C.cuStreamWaitValue32(Cstream, Caddr, Cvalue, Cflags))
}

func (stream Stream) WriteValue32(addr DevicePtr, value uint32, flags uint) (err error) {
	Cstream := stream.c()
	Caddr := C.CUdeviceptr(addr)
	Cvalue := C.cuuint32_t(value)
	Cflags := C.uint(flags)
	return result(C.cuStreamWriteValue32(Cstream, Caddr, Cvalue, Cflags))
}

func (fn Function) Attribute(attrib FunctionAttribute) (pi int, err error) {
	Cfn := fn.c()
	Cattrib := C.CUfunction_attribute(attrib)
	var Cpi C.int
	err = result(C.cuFuncGetAttribute(&Cpi, Cattrib, Cfn))
	pi = int(Cpi)

	return
}

func (fn Function) SetCacheConfig(config FuncCacheConfig) (err error) {
	Cfn := fn.c()
	Cconfig := C.CUfunc_cache(config)
	return result(C.cuFuncSetCacheConfig(Cfn, Cconfig))
}

func (fn Function) SetSharedMemConfig(config SharedConfig) (err error) {
	Cfn := fn.c()
	Cconfig := C.CUsharedconfig(config)
	return result(C.cuFuncSetSharedMemConfig(Cfn, Cconfig))
}

func (hTexRef TexRef) SetArray(hArray Array, Flags uint) (err error) {
	ChTexRef := hTexRef.c()
	ChArray := hArray.c()
	CFlags := C.uint(Flags)
	return result(C.cuTexRefSetArray(ChTexRef, ChArray, CFlags))
}

func (hTexRef TexRef) SetAddress(dptr DevicePtr, bytes int64) (ByteOffset int64, err error) {
	ChTexRef := hTexRef.c()
	Cdptr := C.CUdeviceptr(dptr)
	Cbytes := C.size_t(bytes)
	var CByteOffset C.size_t
	err = result(C.cuTexRefSetAddress(&CByteOffset, ChTexRef, Cdptr, Cbytes))
	ByteOffset = int64(CByteOffset)

	return
}

func (hTexRef TexRef) SetAddress2D(desc ArrayDesc, dptr DevicePtr, Pitch int64) (err error) {
	ChTexRef := hTexRef.c()
	Cdesc := desc.c()
	Cdptr := C.CUdeviceptr(dptr)
	CPitch := C.size_t(Pitch)
	return result(C.cuTexRefSetAddress2D(ChTexRef, Cdesc, Cdptr, CPitch))
}

func (hTexRef TexRef) SetFormat(fmt Format, NumPackedComponents int) (err error) {
	ChTexRef := hTexRef.c()
	Cfmt := C.CUarray_format(fmt)
	CNumPackedComponents := C.int(NumPackedComponents)
	return result(C.cuTexRefSetFormat(ChTexRef, Cfmt, CNumPackedComponents))
}

func (hTexRef TexRef) SetAddressMode(dim int, am AddressMode) (err error) {
	ChTexRef := hTexRef.c()
	Cdim := C.int(dim)
	Cam := C.CUaddress_mode(am)
	return result(C.cuTexRefSetAddressMode(ChTexRef, Cdim, Cam))
}

func (hTexRef TexRef) SetFilterMode(fm FilterMode) (err error) {
	ChTexRef := hTexRef.c()
	Cfm := C.CUfilter_mode(fm)
	return result(C.cuTexRefSetFilterMode(ChTexRef, Cfm))
}

func (hTexRef TexRef) SetMipmapFilterMode(fm FilterMode) (err error) {
	ChTexRef := hTexRef.c()
	Cfm := C.CUfilter_mode(fm)
	return result(C.cuTexRefSetMipmapFilterMode(ChTexRef, Cfm))
}

func (hTexRef TexRef) SetMipmapLevelBias(bias float64) (err error) {
	ChTexRef := hTexRef.c()
	Cbias := C.float(bias)
	return result(C.cuTexRefSetMipmapLevelBias(ChTexRef, Cbias))
}

func (hTexRef TexRef) SetMipmapLevelClamp(minMipmapLevelClamp float64, maxMipmapLevelClamp float64) (err error) {
	ChTexRef := hTexRef.c()
	CminMipmapLevelClamp := C.float(minMipmapLevelClamp)
	CmaxMipmapLevelClamp := C.float(maxMipmapLevelClamp)
	return result(C.cuTexRefSetMipmapLevelClamp(ChTexRef, CminMipmapLevelClamp, CmaxMipmapLevelClamp))
}

func (hTexRef TexRef) SetMaxAnisotropy(maxAniso uint) (err error) {
	ChTexRef := hTexRef.c()
	CmaxAniso := C.uint(maxAniso)
	return result(C.cuTexRefSetMaxAnisotropy(ChTexRef, CmaxAniso))
}

func (hTexRef TexRef) SetBorderColor(pBorderColor [3]float32) (err error) {
	ChTexRef := hTexRef.c()
	CpBorderColor := (*C.float)(unsafe.Pointer(&pBorderColor[0]))
	return result(C.cuTexRefSetBorderColor(ChTexRef, CpBorderColor))
}

func (hTexRef TexRef) SetFlags(Flags TexRefFlags) (err error) {
	ChTexRef := hTexRef.c()
	CFlags := C.uint(Flags)
	return result(C.cuTexRefSetFlags(ChTexRef, CFlags))
}

func (hTexRef TexRef) Address() (pdptr DevicePtr, err error) {
	ChTexRef := hTexRef.c()
	var Cpdptr C.CUdeviceptr
	err = result(C.cuTexRefGetAddress(&Cpdptr, ChTexRef))
	pdptr = DevicePtr(Cpdptr)

	return
}

func (hTexRef TexRef) Array() (phArray Array, err error) {
	ChTexRef := hTexRef.c()
	var CphArray C.CUarray
	err = result(C.cuTexRefGetArray(&CphArray, ChTexRef))
	phArray = goArray(&CphArray)
	return
}

func (hTexRef TexRef) AddressMode(dim int) (pam AddressMode, err error) {
	ChTexRef := hTexRef.c()
	Cdim := C.int(dim)
	var Cpam C.CUaddress_mode
	err = result(C.cuTexRefGetAddressMode(&Cpam, ChTexRef, Cdim))
	pam = AddressMode(Cpam)

	return
}

func (hTexRef TexRef) FilterMode() (pfm FilterMode, err error) {
	ChTexRef := hTexRef.c()
	var Cpfm C.CUfilter_mode
	err = result(C.cuTexRefGetFilterMode(&Cpfm, ChTexRef))
	pfm = FilterMode(Cpfm)

	return
}

func (hTexRef TexRef) Format() (pFormat Format, pNumChannels int, err error) {
	ChTexRef := hTexRef.c()
	var CpFormat C.CUarray_format
	var CpNumChannels C.int
	err = result(C.cuTexRefGetFormat(&CpFormat, &CpNumChannels, ChTexRef))
	pFormat = Format(CpFormat)

	pNumChannels = int(CpNumChannels)

	return
}

func (hTexRef TexRef) MaxAnisotropy() (pmaxAniso int, err error) {
	ChTexRef := hTexRef.c()
	var CpmaxAniso C.int
	err = result(C.cuTexRefGetMaxAnisotropy(&CpmaxAniso, ChTexRef))
	pmaxAniso = int(CpmaxAniso)

	return
}

func (hTexRef TexRef) BorderColor() (pBorderColor [3]float32, err error) {
	ChTexRef := hTexRef.c()
	CpBorderColor := (*C.float)(unsafe.Pointer(&pBorderColor[0]))
	err = result(C.cuTexRefGetBorderColor(CpBorderColor, ChTexRef))
	return
}

func (hTexRef TexRef) Flags() (pFlags TexRefFlags, err error) {
	ChTexRef := hTexRef.c()
	var CpFlags C.uint
	err = result(C.cuTexRefGetFlags(&CpFlags, ChTexRef))
	pFlags = TexRefFlags(CpFlags)

	return
}

func (hSurfRef SurfRef) SetArray(hArray Array, Flags uint) (err error) {
	ChSurfRef := hSurfRef.c()
	ChArray := hArray.c()
	CFlags := C.uint(Flags)
	return result(C.cuSurfRefSetArray(ChSurfRef, ChArray, CFlags))
}

func (hSurfRef SurfRef) GetArray() (phArray Array, err error) {
	ChSurfRef := hSurfRef.c()
	var CphArray C.CUarray
	err = result(C.cuSurfRefGetArray(&CphArray, ChSurfRef))
	phArray = goArray(&CphArray)
	return
}

func (dev Device) CanAccessPeer(peerDev Device) (canAccessPeer int, err error) {
	Cdev := C.CUdevice(dev)
	CpeerDev := C.CUdevice(peerDev)
	var CcanAccessPeer C.int
	err = result(C.cuDeviceCanAccessPeer(&CcanAccessPeer, Cdev, CpeerDev))
	canAccessPeer = int(CcanAccessPeer)

	return
}

func (srcDevice Device) P2PAttribute(attrib P2PAttribute, dstDevice Device) (value int, err error) {
	CsrcDevice := C.CUdevice(srcDevice)
	Cattrib := C.CUdevice_P2PAttribute(attrib)
	CdstDevice := C.CUdevice(dstDevice)
	var Cvalue C.int
	err = result(C.cuDeviceGetP2PAttribute(&Cvalue, Cattrib, CsrcDevice, CdstDevice))
	value = int(Cvalue)

	return
}

func (peerContext CUContext) EnablePeerAccess(Flags uint) (err error) {
	CpeerContext := peerContext.c()
	CFlags := C.uint(Flags)
	return result(C.cuCtxEnablePeerAccess(CpeerContext, CFlags))
}

func (peerContext CUContext) DisablePeerAccess() (err error) {
	CpeerContext := peerContext.c()
	return result(C.cuCtxDisablePeerAccess(CpeerContext))
}
