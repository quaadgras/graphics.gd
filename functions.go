//go:build !generate

package gd

/*
	#include <gdnative_interface.h>

	extern uint8_t get_virtual_index(void *p_userdata, const char *p_name);
	extern void call_virtual_index(uint8_t index, uintptr_t instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret);

	void call_virtual_index1(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)1, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index2(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)2, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index3(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)3, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index4(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)4, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index5(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)5, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index6(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)6, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index7(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)7, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index8(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)8, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index9(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)9, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index10(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)10, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index11(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)11, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index12(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)12, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index13(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)13, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index14(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)14, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index15(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)15, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index16(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)16, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index17(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)17, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index18(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)18, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index19(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)19, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index20(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)20, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index21(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)21, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index22(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)22, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index23(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)23, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index24(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)24, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index25(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)25, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index26(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)26, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index27(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)27, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index28(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)28, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index29(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)29, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index30(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)30, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index31(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)31, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index32(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)32, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index33(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)33, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index34(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)34, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index35(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)35, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index36(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)36, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index37(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)37, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index38(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)38, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index39(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)39, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index40(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)40, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index41(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)41, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index42(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)42, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index43(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)43, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index44(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)44, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index45(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)45, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index46(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)46, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index47(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)47, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index48(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)48, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index49(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)49, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index50(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)50, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index51(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)51, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index52(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)52, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index53(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)53, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index54(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)54, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index55(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)55, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index56(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)56, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index57(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)57, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index58(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)58, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index59(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)59, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index60(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)60, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index61(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)61, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index62(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)62, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index63(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)63, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index64(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)64, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index65(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)65, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index66(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)66, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index67(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)67, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index68(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)68, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index69(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)69, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index70(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)70, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index71(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)71, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index72(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)72, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index73(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)73, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index74(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)74, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index75(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)75, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index76(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)76, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index77(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)77, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index78(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)78, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index79(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)79, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index80(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)80, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index81(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)81, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index82(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)82, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index83(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)83, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index84(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)84, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index85(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)85, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index86(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)86, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index87(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)87, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index88(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)88, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index89(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)89, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index90(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)90, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index91(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)91, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index92(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)92, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index93(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)93, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index94(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)94, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index95(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)95, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index96(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)96, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index97(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)97, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index98(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)98, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index99(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)99, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index100(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)100, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index101(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)101, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index102(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)102, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index103(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)103, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index104(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)104, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index105(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)105, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index106(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)106, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index107(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)107, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index108(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)108, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index109(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)109, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index110(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)110, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index111(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)111, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index112(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)112, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index113(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)113, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index114(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)114, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index115(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)115, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index116(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)116, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index117(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)117, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index118(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)118, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index119(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)119, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index120(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)120, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index121(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)121, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index122(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)122, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index123(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)123, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index124(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)124, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index125(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)125, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index126(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)126, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index127(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)127, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index128(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)128, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index129(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)129, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index130(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)130, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index131(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)131, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index132(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)132, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index133(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)133, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index134(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)134, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index135(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)135, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index136(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)136, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index137(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)137, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index138(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)138, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index139(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)139, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index140(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)140, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index141(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)141, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index142(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)142, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index143(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)143, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index144(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)144, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index145(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)145, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index146(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)146, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index147(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)147, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index148(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)148, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index149(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)149, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index150(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)150, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index151(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)151, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index152(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)152, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index153(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)153, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index154(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)154, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index155(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)155, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index156(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)156, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index157(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)157, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index158(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)158, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index159(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)159, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index160(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)160, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index161(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)161, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index162(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)162, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index163(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)163, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index164(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)164, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index165(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)165, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index166(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)166, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index167(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)167, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index168(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)168, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index169(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)169, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index170(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)170, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index171(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)171, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index172(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)172, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index173(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)173, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index174(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)174, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index175(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)175, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index176(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)176, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index177(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)177, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index178(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)178, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index179(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)179, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index180(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)180, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index181(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)181, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index182(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)182, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index183(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)183, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index184(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)184, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index185(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)185, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index186(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)186, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index187(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)187, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index188(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)188, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index189(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)189, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index190(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)190, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index191(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)191, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index192(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)192, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index193(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)193, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index194(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)194, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index195(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)195, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index196(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)196, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index197(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)197, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index198(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)198, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index199(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)199, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index200(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)200, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index201(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)201, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index202(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)202, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index203(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)203, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index204(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)204, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index205(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)205, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index206(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)206, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index207(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)207, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index208(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)208, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index209(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)209, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index210(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)210, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index211(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)211, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index212(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)212, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index213(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)213, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index214(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)214, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index215(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)215, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index216(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)216, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index217(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)217, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index218(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)218, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index219(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)219, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index220(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)220, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index221(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)221, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index222(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)222, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index223(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)223, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index224(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)224, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index225(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)225, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index226(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)226, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index227(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)227, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index228(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)228, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index229(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)229, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index230(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)230, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index231(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)231, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index232(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)232, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index233(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)233, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index234(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)234, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index235(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)235, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index236(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)236, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index237(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)237, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index238(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)238, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index239(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)239, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index240(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)240, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index241(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)241, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index242(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)242, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index243(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)243, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index244(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)244, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index245(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)245, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index246(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)246, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index247(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)247, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index248(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)248, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index249(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)249, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index250(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)250, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index251(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)251, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index252(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)252, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index253(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)253, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index254(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)254, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index255(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)255, (uintptr_t)p_instance,  p_args, r_ret);}
	GDNativeExtensionClassCallVirtual get_virtual_func(void *p_userdata, const char *p_name) {
	switch (get_virtual_index(p_userdata, p_name)) {
		case 1: return call_virtual_index1;
		case 2: return call_virtual_index2;
		case 3: return call_virtual_index3;
		case 4: return call_virtual_index4;
		case 5: return call_virtual_index5;
		case 6: return call_virtual_index6;
		case 7: return call_virtual_index7;
		case 8: return call_virtual_index8;
		case 9: return call_virtual_index9;
		case 10: return call_virtual_index10;
		case 11: return call_virtual_index11;
		case 12: return call_virtual_index12;
		case 13: return call_virtual_index13;
		case 14: return call_virtual_index14;
		case 15: return call_virtual_index15;
		case 16: return call_virtual_index16;
		case 17: return call_virtual_index17;
		case 18: return call_virtual_index18;
		case 19: return call_virtual_index19;
		case 20: return call_virtual_index20;
		case 21: return call_virtual_index21;
		case 22: return call_virtual_index22;
		case 23: return call_virtual_index23;
		case 24: return call_virtual_index24;
		case 25: return call_virtual_index25;
		case 26: return call_virtual_index26;
		case 27: return call_virtual_index27;
		case 28: return call_virtual_index28;
		case 29: return call_virtual_index29;
		case 30: return call_virtual_index30;
		case 31: return call_virtual_index31;
		case 32: return call_virtual_index32;
		case 33: return call_virtual_index33;
		case 34: return call_virtual_index34;
		case 35: return call_virtual_index35;
		case 36: return call_virtual_index36;
		case 37: return call_virtual_index37;
		case 38: return call_virtual_index38;
		case 39: return call_virtual_index39;
		case 40: return call_virtual_index40;
		case 41: return call_virtual_index41;
		case 42: return call_virtual_index42;
		case 43: return call_virtual_index43;
		case 44: return call_virtual_index44;
		case 45: return call_virtual_index45;
		case 46: return call_virtual_index46;
		case 47: return call_virtual_index47;
		case 48: return call_virtual_index48;
		case 49: return call_virtual_index49;
		case 50: return call_virtual_index50;
		case 51: return call_virtual_index51;
		case 52: return call_virtual_index52;
		case 53: return call_virtual_index53;
		case 54: return call_virtual_index54;
		case 55: return call_virtual_index55;
		case 56: return call_virtual_index56;
		case 57: return call_virtual_index57;
		case 58: return call_virtual_index58;
		case 59: return call_virtual_index59;
		case 60: return call_virtual_index60;
		case 61: return call_virtual_index61;
		case 62: return call_virtual_index62;
		case 63: return call_virtual_index63;
		case 64: return call_virtual_index64;
		case 65: return call_virtual_index65;
		case 66: return call_virtual_index66;
		case 67: return call_virtual_index67;
		case 68: return call_virtual_index68;
		case 69: return call_virtual_index69;
		case 70: return call_virtual_index70;
		case 71: return call_virtual_index71;
		case 72: return call_virtual_index72;
		case 73: return call_virtual_index73;
		case 74: return call_virtual_index74;
		case 75: return call_virtual_index75;
		case 76: return call_virtual_index76;
		case 77: return call_virtual_index77;
		case 78: return call_virtual_index78;
		case 79: return call_virtual_index79;
		case 80: return call_virtual_index80;
		case 81: return call_virtual_index81;
		case 82: return call_virtual_index82;
		case 83: return call_virtual_index83;
		case 84: return call_virtual_index84;
		case 85: return call_virtual_index85;
		case 86: return call_virtual_index86;
		case 87: return call_virtual_index87;
		case 88: return call_virtual_index88;
		case 89: return call_virtual_index89;
		case 90: return call_virtual_index90;
		case 91: return call_virtual_index91;
		case 92: return call_virtual_index92;
		case 93: return call_virtual_index93;
		case 94: return call_virtual_index94;
		case 95: return call_virtual_index95;
		case 96: return call_virtual_index96;
		case 97: return call_virtual_index97;
		case 98: return call_virtual_index98;
		case 99: return call_virtual_index99;
		case 100: return call_virtual_index100;
		case 101: return call_virtual_index101;
		case 102: return call_virtual_index102;
		case 103: return call_virtual_index103;
		case 104: return call_virtual_index104;
		case 105: return call_virtual_index105;
		case 106: return call_virtual_index106;
		case 107: return call_virtual_index107;
		case 108: return call_virtual_index108;
		case 109: return call_virtual_index109;
		case 110: return call_virtual_index110;
		case 111: return call_virtual_index111;
		case 112: return call_virtual_index112;
		case 113: return call_virtual_index113;
		case 114: return call_virtual_index114;
		case 115: return call_virtual_index115;
		case 116: return call_virtual_index116;
		case 117: return call_virtual_index117;
		case 118: return call_virtual_index118;
		case 119: return call_virtual_index119;
		case 120: return call_virtual_index120;
		case 121: return call_virtual_index121;
		case 122: return call_virtual_index122;
		case 123: return call_virtual_index123;
		case 124: return call_virtual_index124;
		case 125: return call_virtual_index125;
		case 126: return call_virtual_index126;
		case 127: return call_virtual_index127;
		case 128: return call_virtual_index128;
		case 129: return call_virtual_index129;
		case 130: return call_virtual_index130;
		case 131: return call_virtual_index131;
		case 132: return call_virtual_index132;
		case 133: return call_virtual_index133;
		case 134: return call_virtual_index134;
		case 135: return call_virtual_index135;
		case 136: return call_virtual_index136;
		case 137: return call_virtual_index137;
		case 138: return call_virtual_index138;
		case 139: return call_virtual_index139;
		case 140: return call_virtual_index140;
		case 141: return call_virtual_index141;
		case 142: return call_virtual_index142;
		case 143: return call_virtual_index143;
		case 144: return call_virtual_index144;
		case 145: return call_virtual_index145;
		case 146: return call_virtual_index146;
		case 147: return call_virtual_index147;
		case 148: return call_virtual_index148;
		case 149: return call_virtual_index149;
		case 150: return call_virtual_index150;
		case 151: return call_virtual_index151;
		case 152: return call_virtual_index152;
		case 153: return call_virtual_index153;
		case 154: return call_virtual_index154;
		case 155: return call_virtual_index155;
		case 156: return call_virtual_index156;
		case 157: return call_virtual_index157;
		case 158: return call_virtual_index158;
		case 159: return call_virtual_index159;
		case 160: return call_virtual_index160;
		case 161: return call_virtual_index161;
		case 162: return call_virtual_index162;
		case 163: return call_virtual_index163;
		case 164: return call_virtual_index164;
		case 165: return call_virtual_index165;
		case 166: return call_virtual_index166;
		case 167: return call_virtual_index167;
		case 168: return call_virtual_index168;
		case 169: return call_virtual_index169;
		case 170: return call_virtual_index170;
		case 171: return call_virtual_index171;
		case 172: return call_virtual_index172;
		case 173: return call_virtual_index173;
		case 174: return call_virtual_index174;
		case 175: return call_virtual_index175;
		case 176: return call_virtual_index176;
		case 177: return call_virtual_index177;
		case 178: return call_virtual_index178;
		case 179: return call_virtual_index179;
		case 180: return call_virtual_index180;
		case 181: return call_virtual_index181;
		case 182: return call_virtual_index182;
		case 183: return call_virtual_index183;
		case 184: return call_virtual_index184;
		case 185: return call_virtual_index185;
		case 186: return call_virtual_index186;
		case 187: return call_virtual_index187;
		case 188: return call_virtual_index188;
		case 189: return call_virtual_index189;
		case 190: return call_virtual_index190;
		case 191: return call_virtual_index191;
		case 192: return call_virtual_index192;
		case 193: return call_virtual_index193;
		case 194: return call_virtual_index194;
		case 195: return call_virtual_index195;
		case 196: return call_virtual_index196;
		case 197: return call_virtual_index197;
		case 198: return call_virtual_index198;
		case 199: return call_virtual_index199;
		case 200: return call_virtual_index200;
		case 201: return call_virtual_index201;
		case 202: return call_virtual_index202;
		case 203: return call_virtual_index203;
		case 204: return call_virtual_index204;
		case 205: return call_virtual_index205;
		case 206: return call_virtual_index206;
		case 207: return call_virtual_index207;
		case 208: return call_virtual_index208;
		case 209: return call_virtual_index209;
		case 210: return call_virtual_index210;
		case 211: return call_virtual_index211;
		case 212: return call_virtual_index212;
		case 213: return call_virtual_index213;
		case 214: return call_virtual_index214;
		case 215: return call_virtual_index215;
		case 216: return call_virtual_index216;
		case 217: return call_virtual_index217;
		case 218: return call_virtual_index218;
		case 219: return call_virtual_index219;
		case 220: return call_virtual_index220;
		case 221: return call_virtual_index221;
		case 222: return call_virtual_index222;
		case 223: return call_virtual_index223;
		case 224: return call_virtual_index224;
		case 225: return call_virtual_index225;
		case 226: return call_virtual_index226;
		case 227: return call_virtual_index227;
		case 228: return call_virtual_index228;
		case 229: return call_virtual_index229;
		case 230: return call_virtual_index230;
		case 231: return call_virtual_index231;
		case 232: return call_virtual_index232;
		case 233: return call_virtual_index233;
		case 234: return call_virtual_index234;
		case 235: return call_virtual_index235;
		case 236: return call_virtual_index236;
		case 237: return call_virtual_index237;
		case 238: return call_virtual_index238;
		case 239: return call_virtual_index239;
		case 240: return call_virtual_index240;
		case 241: return call_virtual_index241;
		case 242: return call_virtual_index242;
		case 243: return call_virtual_index243;
		case 244: return call_virtual_index244;
		case 245: return call_virtual_index245;
		case 246: return call_virtual_index246;
		case 247: return call_virtual_index247;
		case 248: return call_virtual_index248;
		case 249: return call_virtual_index249;
		case 250: return call_virtual_index250;
		case 251: return call_virtual_index251;
		case 252: return call_virtual_index252;
		case 253: return call_virtual_index253;
		case 254: return call_virtual_index254;
		case 255: return call_virtual_index255;
		default: return NULL;}
	}

	void method0(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result) {
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, NULL, (GDNativeTypePtr*)result);
	}
	void method1(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a) {
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&a), (GDNativeTypePtr*)result);
	}
	void method2(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b) {
		GDNativeTypePtr stack[2] = {a, b};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method3(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c) {
		GDNativeTypePtr stack[3] = {a, b, c};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method4(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d) {
		GDNativeTypePtr stack[4] = {a, b, c, d};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method5(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e) {
		GDNativeTypePtr stack[5] = {a, b, c, d, e};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method6(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f) {
		GDNativeTypePtr stack[6] = {a, b, c, d, e, f};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method7(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g) {
		GDNativeTypePtr stack[7] = {a, b, c, d, e, f, g};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method8(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h) {
		GDNativeTypePtr stack[8] = {a, b, c, d, e, f, g, h};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method9(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i) {
		GDNativeTypePtr stack[9] = {a, b, c, d, e, f, g, h, i};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method10(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i, void *j) {
		GDNativeTypePtr stack[10] = {a, b, c, d, e, f, g, h, i, j};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method11(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i, void *j, void *k) {
		GDNativeTypePtr stack[11] = {a, b, c, d, e, f, g, h, i, j, k};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method12(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i, void *j, void *k, void *l) {
		GDNativeTypePtr stack[12] = {a, b, c, d, e, f, g, h, i, j, k, l};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method13(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i, void *j, void *k, void *l, void *m) {
		GDNativeTypePtr stack[13] = {a, b, c, d, e, f, g, h, i, j, k, l, m};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}

	void utility0(uintptr_t utility, void *result) {
		((GDNativePtrUtilityFunction)utility)(result, NULL, 0);
	}
	void utility1(uintptr_t utility, void *r_ret, void *a) {
		((GDNativePtrUtilityFunction)(utility))(r_ret, &a, 1);
	}
	void utility2(uintptr_t utility, void *r_ret, void *a, void *b) {
		GDNativeTypePtr stack[2] = { a, b };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 2);
	}
	void utility3(uintptr_t utility, void *r_ret, void *a, void *b, void *c) {
		GDNativeTypePtr stack[3] = { a, b, c };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 3);
	}
	void utility4(uintptr_t utility, void *r_ret, void *a, void *b, void *c, void *d) {
		GDNativeTypePtr stack[4] = { a, b, c, d };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 4);
	}
	void utility5(uintptr_t utility, void *r_ret, void *a, void *b, void *c, void *d, void *e) {
		GDNativeTypePtr stack[5] = { a, b, c, d, e };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 5);
	}

	void builtin0(void *p_base, uintptr_t builtin, void *r_ret) {
		((GDNativePtrBuiltInMethod)builtin)(p_base, NULL, r_ret, 0);
	}
	void builtin1(void *p_base, uintptr_t builtin, void *r_ret, void *a) {
		((GDNativePtrBuiltInMethod)(builtin))(p_base, &a, r_ret, 1);
	}
	void builtin2(void *p_base, uintptr_t builtin, void *r_ret, void *a, void *b) {
		GDNativeTypePtr stack[2] = { a, b };
		((GDNativePtrBuiltInMethod)(builtin))(p_base, &stack[0], r_ret, 2);
	}
	void builtin3(void *p_base, uintptr_t builtin, void *r_ret, void *a, void *b, void *c) {
		GDNativeTypePtr stack[3] = { a, b, c };
		((GDNativePtrBuiltInMethod)(builtin))(p_base, &stack[0], r_ret, 3);
	}
	void builtin4(void *p_base, uintptr_t builtin, void *r_ret, void *a, void *b, void *c, void *d) {
		GDNativeTypePtr stack[4] = { a, b, c, d };
		((GDNativePtrBuiltInMethod)(builtin))(p_base, &stack[0], r_ret, 4);
	}
	void builtin5(void *p_base, uintptr_t builtin, void *r_ret, void *a, void *b, void *c, void *d, void *e) {
		GDNativeTypePtr stack[5] = { a, b, c, d, e };
		((GDNativePtrBuiltInMethod)(builtin))(p_base, &stack[0], r_ret, 5);
	}
	void builtin6(void *p_base, uintptr_t builtin, void *r_ret, void *a, void *b, void *c, void *d, void *e, void *f) {
		GDNativeTypePtr stack[6] = { a, b, c, d, e, f };
		((GDNativePtrBuiltInMethod)(builtin))(p_base, &stack[0], r_ret, 6);
	}
	void builtin7(void *p_base, uintptr_t builtin, void *r_ret, void *a, void *b, void *c, void *d, void *e, void *f, void *g) {
		GDNativeTypePtr stack[7] = { a, b, c, d, e, f, g };
		((GDNativePtrBuiltInMethod)(builtin))(p_base, &stack[0], r_ret, 7);
	}
	void builtin8(void *p_base, uintptr_t builtin, void *r_ret, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h) {
		GDNativeTypePtr stack[8] = { a, b, c, d, e, f, g, h };
		((GDNativePtrBuiltInMethod)(builtin))(p_base, &stack[0], r_ret, 8);
	}

	void constructor0(void *p_base, uintptr_t constructor) {
		((GDNativePtrConstructor)constructor)(p_base, NULL);
	}
	void constructor1(void *p_base, uintptr_t constructor, void *a) {
		((GDNativePtrConstructor)(constructor))(p_base, &a);
	}
	void constructor2(void *p_base, uintptr_t constructor, void *a, void *b) {
		GDNativeTypePtr stack[2] = { a, b };
		((GDNativePtrConstructor)(constructor))(p_base, &stack[0]);
	}
	void constructor3(void *p_base, uintptr_t constructor, void *a, void *b, void *c) {
		GDNativeTypePtr stack[3] = { a, b, c };
		((GDNativePtrConstructor)(constructor))(p_base, &stack[0]);
	}
	void constructor4(void *p_base, uintptr_t constructor, void *a, void *b, void *c, void *d) {
		GDNativeTypePtr stack[4] = { a, b, c, d };
		((GDNativePtrConstructor)(constructor))(p_base, &stack[0]);
	}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func cVariantTypeOf(t reflect.Type) cVariantType {
	switch t.Kind() {
	case reflect.Bool:
		return C.GDNATIVE_VARIANT_TYPE_BOOL
	case reflect.Int:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Int8:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Int16:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Int32:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Int64:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Uint:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Uint8:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Uint16:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Uint32:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Uint64:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Float32:
		return C.GDNATIVE_VARIANT_TYPE_FLOAT
	case reflect.Float64:
		return C.GDNATIVE_VARIANT_TYPE_FLOAT
	case reflect.String:
		return C.GDNATIVE_VARIANT_TYPE_STRING
	case reflect.Struct:
		return C.GDNATIVE_VARIANT_TYPE_DICTIONARY
	case reflect.Map:
		return C.GDNATIVE_VARIANT_TYPE_DICTIONARY
	case reflect.Slice:
		return C.GDNATIVE_VARIANT_TYPE_ARRAY
	case reflect.Array:
		return C.GDNATIVE_VARIANT_TYPE_ARRAY
	case reflect.Interface:
		return C.GDNATIVE_VARIANT_TYPE_OBJECT
	case reflect.Ptr:
		return C.GDNATIVE_VARIANT_TYPE_OBJECT
	default:
		panic("unsupported type")
	}
}

func toUnsafe(val any) (ptr unsafe.Pointer, free func()) {
	switch v := val.(type) {
	case *struct{}:
		return nil, nil
	case *bool:
		return unsafe.Pointer(v), nil
	case *int64:
		return unsafe.Pointer(v), nil
	case *float64:
		return unsafe.Pointer(v), nil
	case *cVector2:
		return unsafe.Pointer(v), nil
	case *string:
		var native cString
		if *v != "" {
			native.with_utf8_chars(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypeString.get_destructor().call(unsafe.Pointer(&native))
		}
	case *cPackedStringArray:
		return unsafe.Pointer(v), nil
	case *cName:
		var native cStringName
		var s cString
		if *v != "" {
			s.with_utf8_chars(string(*v))
		}
		constructorCall(unsafe.Pointer(&native), cVariantTypeStringName.get_constructor(2), unsafe.Pointer(&s))
		return unsafe.Pointer(&native), func() {
			cVariantTypeStringName.get_destructor().call(unsafe.Pointer(&native))
			cVariantTypeString.get_destructor().call(unsafe.Pointer(&s))
		}
	case *cNodePath:
		return unsafe.Pointer(v), nil
	case *cRID:
		return unsafe.Pointer(v), nil
	case *cObject:
		return unsafe.Pointer(v), nil
	case *cCallable:
		return unsafe.Pointer(v), nil
	case *cSignal:
		return unsafe.Pointer(v), nil
	case *cDictionary:
		return unsafe.Pointer(v), nil
	case *cArray:
		return unsafe.Pointer(v), nil
	case *[]byte:
		var native cPackedByteArray
		if len(*v) > 0 {
			native.from(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypePackedByteArray.get_destructor().call(unsafe.Pointer(&native))
		}
	case *[]int32:
		var native cPackedInt32Array
		if len(*v) > 0 {
			native.from(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypePackedInt32Array.get_destructor().call(unsafe.Pointer(&native))
		}
	case *[]int64:
		var native cPackedInt64Array
		if len(*v) > 0 {
			native.from(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypePackedInt64Array.get_destructor().call(unsafe.Pointer(&native))
		}
	case *[]float32:
		var native cPackedFloat32Array
		if len(*v) > 0 {
			native.from(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypePackedFloat32Array.get_destructor().call(unsafe.Pointer(&native))
		}
	case *[]float64:
		var native cPackedFloat64Array
		if len(*v) > 0 {
			native.from(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypePackedFloat64Array.get_destructor().call(unsafe.Pointer(&native))
		}
	case *[]string:
		var native cPackedStringArray

		var converted []cString
		for _, s := range *v {
			var c cString
			if s != "" {
				c.with_utf8_chars(s)
			}
			converted = append(converted, c)
		}
		if len(converted) > 0 {
			native.from(converted)
		}

		return unsafe.Pointer(&native), func() {
			cVariantTypePackedStringArray.get_destructor().call(unsafe.Pointer(&native))
		}
	case *[]cVector2:
		var native cPackedVector2Array
		if len(*v) > 0 {
			native.from(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypePackedVector2Array.get_destructor().call(unsafe.Pointer(&native))
		}
	case *[]cVector3:
		var native cPackedVector3Array
		if len(*v) > 0 {
			native.from(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypePackedVector3Array.get_destructor().call(unsafe.Pointer(&native))
		}
	case *[]cColor:
		var native cPackedColorArray
		if len(*v) > 0 {
			native.from(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypePackedColorArray.get_destructor().call(unsafe.Pointer(&native))
		}
	default:
		panic("toUnsafe: unsupported type " + reflect.TypeOf(val).String())
	}
}

func toResult(val any, ptr unsafe.Pointer) {
	switch v := val.(type) {
	case struct{}:
		return
	case bool:
		*(*bool)(ptr) = v
	case *int64:
		*(*int64)(ptr) = *v
	case *float64:
		*(*float64)(ptr) = *v
	case *cVector2:
		*(*cVector2)(ptr) = *v
	case *string:
		var native = (*cString)(ptr)
		if *v != "" {
			native.with_utf8_chars(*v)
		}
	case *cPackedStringArray:
		*(*cPackedStringArray)(ptr) = *v
	case *cName:
		var native = (*cStringName)(ptr)
		var s cString
		if *v != "" {
			s.with_utf8_chars(string(*v))
		}
		constructorCall(unsafe.Pointer(&native), cVariantTypeStringName.get_constructor(2), unsafe.Pointer(&s))
		// FIXME does cString need to be freed?
	case *cNodePath:
		*(*cNodePath)(ptr) = *v
	case *cRID:
		*(*cRID)(ptr) = *v
	case *cObject:
		*(*cObject)(ptr) = *v
	case *cCallable:
		*(*cCallable)(ptr) = *v
	case *cSignal:
		*(*cSignal)(ptr) = *v
	case *cDictionary:
		*(*cDictionary)(ptr) = *v
	case *cArray:
		*(*cArray)(ptr) = *v
	case *[]byte:
		var native = (*cPackedByteArray)(ptr)
		if len(*v) > 0 {
			native.from(*v)
		}
	case *[]int32:
		var native = (*cPackedInt32Array)(ptr)
		if len(*v) > 0 {
			native.from(*v)
		}
	case *[]int64:
		var native = (*cPackedInt64Array)(ptr)
		if len(*v) > 0 {
			native.from(*v)
		}
	case *[]float32:
		var native = (*cPackedFloat32Array)(ptr)
		if len(*v) > 0 {
			native.from(*v)
		}
	case *[]float64:
		var native = (*cPackedFloat64Array)(ptr)
		if len(*v) > 0 {
			native.from(*v)
		}
	case *[]string:
		var native = (*cPackedStringArray)(ptr)

		var converted []cString
		for _, s := range *v {
			var c cString
			if s != "" {
				c.with_utf8_chars(s)
			}
			converted = append(converted, c)
		}
		if len(converted) > 0 {
			native.from(converted)
		}

		// FIXME do cStrings need to be freed?
	case *[]cVector2:
		var native = (*cPackedVector2Array)(ptr)
		if len(*v) > 0 {
			native.from(*v)
		}
	case *[]cVector3:
		var native = (*cPackedVector3Array)(ptr)
		if len(*v) > 0 {
			native.from(*v)
		}
	case *[]cColor:
		var native = (*cPackedColorArray)(ptr)
		if len(*v) > 0 {
			native.from(*v)
		}
	default:
		panic("toResult: unsupported type " + reflect.TypeOf(val).String())
	}
}

func into(result any, ptr unsafe.Pointer) {
	switch v := result.(type) {
	case *struct{}:
		return
	case *bool:
		*v = *(*bool)(ptr)
	case *int64:
		*v = *(*int64)(ptr)
	case *float64:
		*v = *(*float64)(ptr)
	case *string:
		*v = (*cString)(ptr).to_utf8_chars()
	case *cVector2:
		*v = *(*cVector2)(ptr)
	case *cVector2i:
		*v = *(*cVector2i)(ptr)
	case *cRect2:
		*v = *(*cRect2)(ptr)
	case *cRect2i:
		*v = *(*cRect2i)(ptr)
	case *cVector3:
		*v = *(*cVector3)(ptr)
	case *cVector3i:
		*v = *(*cVector3i)(ptr)
	case *cTransform2D:
		*v = *(*cTransform2D)(ptr)
	case *cVector4:
		*v = *(*cVector4)(ptr)
	case *cVector4i:
		*v = *(*cVector4i)(ptr)
	case *cPlane:
		*v = *(*cPlane)(ptr)
	case *cQuaternion:
		*v = *(*cQuaternion)(ptr)
	case *cAABB:
		*v = *(*cAABB)(ptr)
	case *cBasis:
		*v = *(*cBasis)(ptr)
	case *cTransform3D:
		*v = *(*cTransform3D)(ptr)
	case *cProjection:
		*v = *(*cProjection)(ptr)
	case *cColor:
		*v = *(*cColor)(ptr)
	case *[]string:
		packed := (*cPackedStringArray)(ptr).slice()
		for i := range packed {
			*v = append(*v, packed[i].to_utf8_chars())
		}
	default:
		panic("into: unsupported type " + reflect.TypeOf(result).String())
	}
}

func methodCall[Result any](object cObject, method cMethodBind, args ...any) Result {
	var result Result
	c_result, free := toUnsafe(&result)

	var c_args = make([]unsafe.Pointer, len(args))
	for i := range args {
		c_args[i], free = toUnsafe(args[i])
		if free != nil {
			defer free()
		}
	}

	method.unsafeReturn(c_result, object, c_args...)

	if free != nil {
		into(&result, c_result)
		free()
	}

	return result
}

func (method cMethodBind) unsafeReturn(result unsafe.Pointer, object cObject, args ...unsafe.Pointer) {
	switch len(args) {
	case 0:
		C.method0(api, C.uintptr_t(method), C.uintptr_t(object), result)
	case 1:
		C.method1(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0])
	case 2:
		C.method2(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1])
	case 3:
		C.method3(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2])
	case 4:
		C.method4(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3])
	case 5:
		C.method5(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4])
	case 6:
		C.method6(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5])
	case 7:
		C.method7(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6])
	case 8:
		C.method8(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7])
	case 9:
		C.method9(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8])
	case 10:
		C.method10(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8], args[9])
	case 11:
		C.method11(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8], args[9], args[10])
	case 12:
		C.method12(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8], args[9], args[10], args[11])
	case 13:
		C.method13(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8], args[9], args[10], args[11],
			args[12])
	default:
		panic("too many arguments")
	}
}

func utilityCall[Result any](utility cUtilityFunction, args ...any) Result {
	var result Result
	c_result, free := toUnsafe(&result)

	var c_args = make([]unsafe.Pointer, len(args))
	for i := range args {
		c_args[i], free = toUnsafe(args[i])
		if free != nil {
			defer free()
		}
	}

	utility.unsafeReturn(c_result, c_args...)

	if free != nil {
		into(&result, c_result)
		free()
	}

	return result
}

// Return the result of calling the given utility function with the given arguments, no type-checking
// is performed. So the arguments must be passed in as the correct type.
func (utility cUtilityFunction) unsafeReturn(result unsafe.Pointer, args ...unsafe.Pointer) {
	switch len(args) {
	case 0:
		C.utility0(C.uintptr_t(utility), result)
	case 1:
		C.utility1(C.uintptr_t(utility), result, args[0])
	case 2:
		C.utility2(C.uintptr_t(utility), result, args[0],
			args[1])
	case 3:
		C.utility3(C.uintptr_t(utility), result, args[0],
			args[1], args[2])
	case 4:
		C.utility4(C.uintptr_t(utility), result, args[0],
			args[1], args[2], args[3])
	case 5:
		C.utility5(C.uintptr_t(utility), result, args[0],
			args[1], args[2], args[3], args[4])
	default:
		panic("too many arguments")
	}
}

func builtinCall[Result any](base any, builtin cBuiltInMethod, args ...any) Result {
	var result Result
	c_result, free := toUnsafe(&result)

	var c_args = make([]unsafe.Pointer, len(args))
	for i := range args {
		c_args[i], free = toUnsafe(args[i])
		if free != nil {
			defer free()
		}
	}

	c_base, free := toUnsafe(base)
	if free != nil {
		defer free()
	}

	builtin.unsafeReturn(c_base, c_result, c_args...)

	if free != nil {
		into(&result, c_result)
		free()
	}

	return result
}

func (builtin cBuiltInMethod) unsafeReturn(base unsafe.Pointer, result unsafe.Pointer, args ...unsafe.Pointer) {
	switch len(args) {
	case 0:
		C.builtin0(base, C.uintptr_t(builtin), result)
	case 1:
		C.builtin1(base, C.uintptr_t(builtin), result, args[0])
	case 2:
		C.builtin2(base, C.uintptr_t(builtin), result, args[0],
			args[1])
	case 3:
		C.builtin3(base, C.uintptr_t(builtin), result, args[0],
			args[1], args[2])
	case 4:
		C.builtin4(base, C.uintptr_t(builtin), result, args[0],
			args[1], args[2], args[3])
	case 5:
		C.builtin5(base, C.uintptr_t(builtin), result, args[0],
			args[1], args[2], args[3], args[4])
	case 6:
		C.builtin6(base, C.uintptr_t(builtin), result, args[0],
			args[1], args[2], args[3], args[4], args[5])
	case 7:
		C.builtin7(base, C.uintptr_t(builtin), result, args[0],
			args[1], args[2], args[3], args[4], args[5], args[6])
	case 8:
		C.builtin8(base, C.uintptr_t(builtin), result, args[0],
			args[1], args[2], args[3], args[4], args[5], args[6],
			args[7])
	default:
		panic("too many arguments")
	}
}

func constructorCall(base any, constructor cConstructor, args ...any) {
	var free func()
	var c_args = make([]unsafe.Pointer, len(args))
	for i := range args {
		c_args[i], free = toUnsafe(args[i])
		if free != nil {
			defer free()
		}
	}

	c_base, free := toUnsafe(base)
	if free != nil {
		defer free()
	}

	constructor.unsafeReturn(c_base, c_args...)
}

func (constructor cConstructor) unsafeReturn(base unsafe.Pointer, args ...unsafe.Pointer) {
	switch len(args) {
	case 0:
		C.constructor0(base, C.uintptr_t(constructor))
	case 1:
		C.constructor1(base, C.uintptr_t(constructor), args[0])
	case 2:
		C.constructor2(base, C.uintptr_t(constructor), args[0],
			args[1])
	case 3:
		C.constructor3(base, C.uintptr_t(constructor), args[0],
			args[1], args[2])
	case 4:
		C.constructor4(base, C.uintptr_t(constructor), args[0],
			args[1], args[2], args[3])
	default:
		panic("too many arguments")
	}
}
