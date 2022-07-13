// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 05 Sep 2017 19:51:53 MSK.
// By https://git.io/c-for-go. DO NOT EDIT.

package vpx

/*
#cgo pkg-config: vpx
#include <vpx/vpx_encoder.h>
#include <vpx/vpx_decoder.h>
#include <vpx/vp8.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"
import "fmt"

// CodecEncInitVer function as declared in vpx-1.6.0/vpx_encoder.h:789
func CodecEncInitVer(ctx *CodecCtx, iface *CodecIface, cfg *CodecEncCfg, flags CodecFlags, ver int32) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer(iface)), cgoAllocsUnknown
	ccfg, _ := cfg.PassRef()
	cflags, _ := (C.vpx_codec_flags_t)(flags), cgoAllocsUnknown
	cver, _ := (C.int)(ver), cgoAllocsUnknown
	__ret := C.vpx_codec_enc_init_ver(cctx, ciface, ccfg, cflags, cver)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecEncInitMultiVer function as declared in vpx-1.6.0/vpx_encoder.h:824
func CodecEncInitMultiVer(ctx *CodecCtx, iface *CodecIface, cfg *CodecEncCfg, numEnc int32, flags CodecFlags, dsf *Rational, ver int32) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer(iface)), cgoAllocsUnknown
	ccfg, _ := cfg.PassRef()
	cnumEnc, _ := (C.int)(numEnc), cgoAllocsUnknown
	cflags, _ := (C.vpx_codec_flags_t)(flags), cgoAllocsUnknown
	cdsf, _ := dsf.PassRef()
	cver, _ := (C.int)(ver), cgoAllocsUnknown
	__ret := C.vpx_codec_enc_init_multi_ver(cctx, ciface, ccfg, cnumEnc, cflags, cdsf, cver)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecEncConfigDefault function as declared in vpx-1.6.0/vpx_encoder.h:861
func CodecEncConfigDefault(iface *CodecIface, cfg *CodecEncCfg, reserved uint32) CodecErr {
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer(iface)), cgoAllocsUnknown
	ccfg, _ := cfg.PassRef()
	creserved, _ := (C.uint)(reserved), cgoAllocsUnknown
	__ret := C.vpx_codec_enc_config_default(ciface, ccfg, creserved)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecEncConfigSet function as declared in vpx-1.6.0/vpx_encoder.h:880
func CodecEncConfigSet(ctx *CodecCtx, cfg *CodecEncCfg) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccfg, _ := cfg.PassRef()
	__ret := C.vpx_codec_enc_config_set(cctx, ccfg)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecGetGlobalHeaders function as declared in vpx-1.6.0/vpx_encoder.h:895
func CodecGetGlobalHeaders(ctx *CodecCtx) *FixedBuf {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.vpx_codec_get_global_headers(cctx)
	__v := NewFixedBufRef(unsafe.Pointer(__ret))
	return __v
}

// CodecEncode function as declared in vpx-1.6.0/vpx_encoder.h:940
func CodecEncode(ctx *CodecCtx, img *Image, pts CodecPts, duration uint, flags EncFrameFlags, deadline uint) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cimg, _ := img.PassRef()
	cpts, _ := (C.vpx_codec_pts_t)(pts), cgoAllocsUnknown
	cduration, _ := (C.ulong)(duration), cgoAllocsUnknown
	cflags, _ := (C.vpx_enc_frame_flags_t)(flags), cgoAllocsUnknown
	cdeadline, _ := (C.ulong)(deadline), cgoAllocsUnknown
	__ret := C.vpx_codec_encode(cctx, cimg, cpts, cduration, cflags, cdeadline)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecSetCxDataBuf function as declared in vpx-1.6.0/vpx_encoder.h:990
func CodecSetCxDataBuf(ctx *CodecCtx, buf *FixedBuf, padBefore uint32, padAfter uint32) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cbuf, _ := buf.PassRef()
	cpadBefore, _ := (C.uint)(padBefore), cgoAllocsUnknown
	cpadAfter, _ := (C.uint)(padAfter), cgoAllocsUnknown
	__ret := C.vpx_codec_set_cx_data_buf(cctx, cbuf, cpadBefore, cpadAfter)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecGetCxData function as declared in vpx-1.6.0/vpx_encoder.h:1019
func CodecGetCxData(ctx *CodecCtx, iter *CodecIter) *CodecCxPkt {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	citer, _ := (*C.vpx_codec_iter_t)(unsafe.Pointer(iter)), cgoAllocsUnknown
	__ret := C.vpx_codec_get_cx_data(cctx, citer)
	__v := NewCodecCxPktRef(unsafe.Pointer(__ret))
	return __v
}

// CodecGetPreviewFrame function as declared in vpx-1.6.0/vpx_encoder.h:1035
func CodecGetPreviewFrame(ctx *CodecCtx) *Image {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.vpx_codec_get_preview_frame(cctx)
	__v := NewImageRef(unsafe.Pointer(__ret))
	return __v
}

// CodecVersion function as declared in vpx-1.6.0/vpx_codec.h:242
func CodecVersion() int32 {
	__ret := C.vpx_codec_version()
	__v := (int32)(__ret)
	return __v
}

// CodecVersionStr function as declared in vpx-1.6.0/vpx_codec.h:264
func CodecVersionStr() string {
	__ret := C.vpx_codec_version_str()
	__v := packPCharString(__ret)
	return __v
}

// CodecVersionExtraStr function as declared in vpx-1.6.0/vpx_codec.h:273
func CodecVersionExtraStr() string {
	__ret := C.vpx_codec_version_extra_str()
	__v := packPCharString(__ret)
	return __v
}

// CodecBuildConfig function as declared in vpx-1.6.0/vpx_codec.h:282
func CodecBuildConfig() string {
	__ret := C.vpx_codec_build_config()
	__v := packPCharString(__ret)
	return __v
}

// CodecIfaceName function as declared in vpx-1.6.0/vpx_codec.h:292
func CodecIfaceName(iface *CodecIface) string {
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer(iface)), cgoAllocsUnknown
	__ret := C.vpx_codec_iface_name(ciface)
	__v := packPCharString(__ret)
	return __v
}

// CodecErrToString function as declared in vpx-1.6.0/vpx_codec.h:305
func CodecErrToString(err CodecErr) string {
	cerr, _ := (C.vpx_codec_err_t)(err), cgoAllocsUnknown
	__ret := C.vpx_codec_err_to_string(cerr)
	__v := packPCharString(__ret)
	return __v
}

// CodecGetError function as declared in vpx-1.6.0/vpx_codec.h:318
func CodecGetError(ctx *CodecCtx) string {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.vpx_codec_error(cctx)
	__v := packPCharString(__ret)
	return __v
}

// CodecErrorDetail function as declared in vpx-1.6.0/vpx_codec.h:331
func CodecErrorDetail(ctx *CodecCtx) string {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.vpx_codec_error_detail(cctx)
	__v := packPCharString(__ret)
	return __v
}

// CodecDestroy function as declared in vpx-1.6.0/vpx_codec.h:351
func CodecDestroy(ctx *CodecCtx) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.vpx_codec_destroy(cctx)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecGetCaps function as declared in vpx-1.6.0/vpx_codec.h:361
func CodecGetCaps(iface *CodecIface) CodecCaps {
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer(iface)), cgoAllocsUnknown
	__ret := C.vpx_codec_get_caps(ciface)
	__v := (CodecCaps)(__ret)
	return __v
}

// ImageAlloc function as declared in vpx-1.6.0/vpx_image.h:161
func ImageAlloc(img *Image, fmt ImageFormat, dW uint32, dH uint32, align uint32) *Image {
	cimg, _ := img.PassRef()
	cfmt, _ := (C.vpx_img_fmt_t)(fmt), cgoAllocsUnknown
	cdW, _ := (C.uint)(dW), cgoAllocsUnknown
	cdH, _ := (C.uint)(dH), cgoAllocsUnknown
	calign, _ := (C.uint)(align), cgoAllocsUnknown
	__ret := C.vpx_img_alloc(cimg, cfmt, cdW, cdH, calign)
	__v := NewImageRef(unsafe.Pointer(__ret))
	return __v
}

// ImageWrap function as declared in vpx-1.6.0/vpx_image.h:186
func ImageWrap(img *Image, fmt ImageFormat, dW uint32, dH uint32, align uint32, imgData *byte) *Image {
	cimg, _ := img.PassRef()
	cfmt, _ := (C.vpx_img_fmt_t)(fmt), cgoAllocsUnknown
	cdW, _ := (C.uint)(dW), cgoAllocsUnknown
	cdH, _ := (C.uint)(dH), cgoAllocsUnknown
	calign, _ := (C.uint)(align), cgoAllocsUnknown
	cimgData, _ := (*C.uchar)(unsafe.Pointer(imgData)), cgoAllocsUnknown
	__ret := C.vpx_img_wrap(cimg, cfmt, cdW, cdH, calign, cimgData)
	__v := NewImageRef(unsafe.Pointer(__ret))
	return __v
}

// ImageSetRect function as declared in vpx-1.6.0/vpx_image.h:207
func ImageSetRect(img *Image, x uint32, y uint32, w uint32, h uint32) int32 {
	cimg, _ := img.PassRef()
	cx, _ := (C.uint)(x), cgoAllocsUnknown
	cy, _ := (C.uint)(y), cgoAllocsUnknown
	cw, _ := (C.uint)(w), cgoAllocsUnknown
	ch, _ := (C.uint)(h), cgoAllocsUnknown
	__ret := C.vpx_img_set_rect(cimg, cx, cy, cw, ch)
	__v := (int32)(__ret)
	return __v
}

// ImageFlip function as declared in vpx-1.6.0/vpx_image.h:221
func ImageFlip(img *Image) {
	cimg, _ := img.PassRef()
	C.vpx_img_flip(cimg)
}

// ImageFree function as declared in vpx-1.6.0/vpx_image.h:229
func ImageFree(img *Image) {
	cimg, _ := img.PassRef()
	C.vpx_img_free(cimg)
}

// CodecDecInitVer function as declared in vpx-1.6.0/vpx_decoder.h:136
func CodecDecInitVer(ctx *CodecCtx, iface *CodecIface, cfg *CodecDecCfg, flags CodecFlags, ver int32) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer(iface)), cgoAllocsUnknown
	ccfg, _ := cfg.PassRef()
	cflags, _ := (C.vpx_codec_flags_t)(flags), cgoAllocsUnknown
	cver, _ := (C.int)(ver), cgoAllocsUnknown
	fmt.Printf("iface->abi_version:%s \n", *ciface.abi_version)
	__ret := C.vpx_codec_dec_init_ver2(cctx, ciface, ccfg, cflags, cver)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecPeekStreamInfo function as declared in vpx-1.6.0/vpx_decoder.h:167
func CodecPeekStreamInfo(iface *CodecIface, data string, dataSz uint32, si *CodecStreamInfo) CodecErr {
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer(iface)), cgoAllocsUnknown
	cdata, _ := unpackPUint8String(data)
	cdataSz, _ := (C.uint)(dataSz), cgoAllocsUnknown
	csi, _ := si.PassRef()
	__ret := C.vpx_codec_peek_stream_info(ciface, cdata, cdataSz, csi)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecGetStreamInfo function as declared in vpx-1.6.0/vpx_decoder.h:186
func CodecGetStreamInfo(ctx *CodecCtx, si *CodecStreamInfo) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	csi, _ := si.PassRef()
	__ret := C.vpx_codec_get_stream_info(cctx, csi)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecDecode function as declared in vpx-1.6.0/vpx_decoder.h:220
func CodecDecode(ctx *CodecCtx, data string, dataSz uint32, userPriv unsafe.Pointer, deadline int) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cdata, _ := unpackPUint8String(data)
	cdataSz, _ := (C.uint)(dataSz), cgoAllocsUnknown
	cuserPriv, _ := userPriv, cgoAllocsUnknown
	cdeadline, _ := (C.long)(deadline), cgoAllocsUnknown
	__ret := C.vpx_codec_decode(cctx, cdata, cdataSz, cuserPriv, cdeadline)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecGetFrame function as declared in vpx-1.6.0/vpx_decoder.h:242
func CodecGetFrame(ctx *CodecCtx, iter *CodecIter) *Image {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	citer, _ := (*C.vpx_codec_iter_t)(unsafe.Pointer(iter)), cgoAllocsUnknown
	__ret := C.vpx_codec_get_frame(cctx, citer)
	__v := NewImageRef(unsafe.Pointer(__ret))
	return __v
}

// CodecRegisterPutFrameCb function as declared in vpx-1.6.0/vpx_decoder.h:279
func CodecRegisterPutFrameCb(ctx *CodecCtx, cb CodecPutFrameCbFn, userPriv unsafe.Pointer) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccb, _ := cb.PassValue()
	cuserPriv, _ := userPriv, cgoAllocsUnknown
	__ret := C.vpx_codec_register_put_frame_cb(cctx, ccb, cuserPriv)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecRegisterPutSliceCb function as declared in vpx-1.6.0/vpx_decoder.h:321
func CodecRegisterPutSliceCb(ctx *CodecCtx, cb CodecPutSliceCbFn, userPriv unsafe.Pointer) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccb, _ := cb.PassValue()
	cuserPriv, _ := userPriv, cgoAllocsUnknown
	__ret := C.vpx_codec_register_put_slice_cb(cctx, ccb, cuserPriv)
	__v := (CodecErr)(__ret)
	return __v
}

// CodecSetFrameBufferFunctions function as declared in vpx-1.6.0/vpx_decoder.h:366
func CodecSetFrameBufferFunctions(ctx *CodecCtx, cbGet GetFrameBufferCbFn, cbRelease ReleaseFrameBufferCbFn, cbPriv unsafe.Pointer) CodecErr {
	cctx, _ := (*C.vpx_codec_ctx_t)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccbGet, _ := cbGet.PassValue()
	ccbRelease, _ := cbRelease.PassValue()
	ccbPriv, _ := cbPriv, cgoAllocsUnknown
	__ret := C.vpx_codec_set_frame_buffer_functions(cctx, ccbGet, ccbRelease, ccbPriv)
	__v := (CodecErr)(__ret)
	return __v
}