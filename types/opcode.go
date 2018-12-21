package types

type Opcode uint16

const (
	OPCunreachable                Opcode = 0x0000
	OPCbr                         Opcode = 0x000c
	OPCbr_if                      Opcode = 0x000d
	OPCbr_table                   Opcode = 0x000e
	OPCreturn_                    Opcode = 0x000f
	OPCcall                       Opcode = 0x0010
	OPCcall_indirect              Opcode = 0x0011
	OPCdrop                       Opcode = 0x001a
	OPCselect                     Opcode = 0x001b
	OPCget_local                  Opcode = 0x0020
	OPCset_local                  Opcode = 0x0021
	OPCtee_local                  Opcode = 0x0022
	OPCget_global                 Opcode = 0x0023
	OPCset_global                 Opcode = 0x0024
	OPCtable_get                  Opcode = 0x0025
	OPCtable_set                  Opcode = 0x0026
	OPCthrow_                     Opcode = 0xfb00
	OPCrethrow                    Opcode = 0xfb01
	OPCnop                        Opcode = 0x0001
	OPCi32_load                   Opcode = 0x0028
	OPCi64_load                   Opcode = 0x0029
	OPCf32_load                   Opcode = 0x002a
	OPCf64_load                   Opcode = 0x002b
	OPCi32_load8_s                Opcode = 0x002c
	OPCi32_load8_u                Opcode = 0x002d
	OPCi32_load16_s               Opcode = 0x002e
	OPCi32_load16_u               Opcode = 0x002f
	OPCi64_load8_s                Opcode = 0x0030
	OPCi64_load8_u                Opcode = 0x0031
	OPCi64_load16_s               Opcode = 0x0032
	OPCi64_load16_u               Opcode = 0x0033
	OPCi64_load32_s               Opcode = 0x0034
	OPCi64_load32_u               Opcode = 0x0035
	OPCi32_store                  Opcode = 0x0036
	OPCi64_store                  Opcode = 0x0037
	OPCf32_store                  Opcode = 0x0038
	OPCf64_store                  Opcode = 0x0039
	OPCi32_store8                 Opcode = 0x003a
	OPCi32_store16                Opcode = 0x003b
	OPCi64_store8                 Opcode = 0x003c
	OPCi64_store16                Opcode = 0x003d
	OPCi64_store32                Opcode = 0x003e
	OPCmemory_size                Opcode = 0x003f
	OPCmemory_grow                Opcode = 0x0040
	OPCi32_const                  Opcode = 0x0041
	OPCi64_const                  Opcode = 0x0042
	OPCf32_const                  Opcode = 0x0043
	OPCf64_const                  Opcode = 0x0044
	OPCi32_eqz                    Opcode = 0x0045
	OPCi32_eq                     Opcode = 0x0046
	OPCi32_ne                     Opcode = 0x0047
	OPCi32_lt_s                   Opcode = 0x0048
	OPCi32_lt_u                   Opcode = 0x0049
	OPCi32_gt_s                   Opcode = 0x004a
	OPCi32_gt_u                   Opcode = 0x004b
	OPCi32_le_s                   Opcode = 0x004c
	OPCi32_le_u                   Opcode = 0x004d
	OPCi32_ge_s                   Opcode = 0x004e
	OPCi32_ge_u                   Opcode = 0x004f
	OPCi64_eqz                    Opcode = 0x0050
	OPCi64_eq                     Opcode = 0x0051
	OPCi64_ne                     Opcode = 0x0052
	OPCi64_lt_s                   Opcode = 0x0053
	OPCi64_lt_u                   Opcode = 0x0054
	OPCi64_gt_s                   Opcode = 0x0055
	OPCi64_gt_u                   Opcode = 0x0056
	OPCi64_le_s                   Opcode = 0x0057
	OPCi64_le_u                   Opcode = 0x0058
	OPCi64_ge_s                   Opcode = 0x0059
	OPCi64_ge_u                   Opcode = 0x005a
	OPCf32_eq                     Opcode = 0x005b
	OPCf32_ne                     Opcode = 0x005c
	OPCf32_lt                     Opcode = 0x005d
	OPCf32_gt                     Opcode = 0x005e
	OPCf32_le                     Opcode = 0x005f
	OPCf32_ge                     Opcode = 0x0060
	OPCf64_eq                     Opcode = 0x0061
	OPCf64_ne                     Opcode = 0x0062
	OPCf64_lt                     Opcode = 0x0063
	OPCf64_gt                     Opcode = 0x0064
	OPCf64_le                     Opcode = 0x0065
	OPCf64_ge                     Opcode = 0x0066
	OPCi32_clz                    Opcode = 0x0067
	OPCi32_ctz                    Opcode = 0x0068
	OPCi32_popcnt                 Opcode = 0x0069
	OPCi32_add                    Opcode = 0x006a
	OPCi32_sub                    Opcode = 0x006b
	OPCi32_mul                    Opcode = 0x006c
	OPCi32_div_s                  Opcode = 0x006d
	OPCi32_div_u                  Opcode = 0x006e
	OPCi32_rem_s                  Opcode = 0x006f
	OPCi32_rem_u                  Opcode = 0x0070
	OPCi32_and_                   Opcode = 0x0071
	OPCi32_or_                    Opcode = 0x0072
	OPCi32_xor_                   Opcode = 0x0073
	OPCi32_shl                    Opcode = 0x0074
	OPCi32_shr_s                  Opcode = 0x0075
	OPCi32_shr_u                  Opcode = 0x0076
	OPCi32_rotl                   Opcode = 0x0077
	OPCi32_rotr                   Opcode = 0x0078
	OPCi64_clz                    Opcode = 0x0079
	OPCi64_ctz                    Opcode = 0x007a
	OPCi64_popcnt                 Opcode = 0x007b
	OPCi64_add                    Opcode = 0x007c
	OPCi64_sub                    Opcode = 0x007d
	OPCi64_mul                    Opcode = 0x007e
	OPCi64_div_s                  Opcode = 0x007f
	OPCi64_div_u                  Opcode = 0x0080
	OPCi64_rem_s                  Opcode = 0x0081
	OPCi64_rem_u                  Opcode = 0x0082
	OPCi64_and_                   Opcode = 0x0083
	OPCi64_or_                    Opcode = 0x0084
	OPCi64_xor_                   Opcode = 0x0085
	OPCi64_shl                    Opcode = 0x0086
	OPCi64_shr_s                  Opcode = 0x0087
	OPCi64_shr_u                  Opcode = 0x0088
	OPCi64_rotl                   Opcode = 0x0089
	OPCi64_rotr                   Opcode = 0x008a
	OPCf32_abs                    Opcode = 0x008b
	OPCf32_neg                    Opcode = 0x008c
	OPCf32_ceil                   Opcode = 0x008d
	OPCf32_floor                  Opcode = 0x008e
	OPCf32_trunc                  Opcode = 0x008f
	OPCf32_nearest                Opcode = 0x0090
	OPCf32_sqrt                   Opcode = 0x0091
	OPCf32_add                    Opcode = 0x0092
	OPCf32_sub                    Opcode = 0x0093
	OPCf32_mul                    Opcode = 0x0094
	OPCf32_div                    Opcode = 0x0095
	OPCf32_min                    Opcode = 0x0096
	OPCf32_max                    Opcode = 0x0097
	OPCf32_copysign               Opcode = 0x0098
	OPCf64_abs                    Opcode = 0x0099
	OPCf64_neg                    Opcode = 0x009a
	OPCf64_ceil                   Opcode = 0x009b
	OPCf64_floor                  Opcode = 0x009c
	OPCf64_trunc                  Opcode = 0x009d
	OPCf64_nearest                Opcode = 0x009e
	OPCf64_sqrt                   Opcode = 0x009f
	OPCf64_add                    Opcode = 0x00a0
	OPCf64_sub                    Opcode = 0x00a1
	OPCf64_mul                    Opcode = 0x00a2
	OPCf64_div                    Opcode = 0x00a3
	OPCf64_min                    Opcode = 0x00a4
	OPCf64_max                    Opcode = 0x00a5
	OPCf64_copysign               Opcode = 0x00a6
	OPCi32_wrap_i64               Opcode = 0x00a7
	OPCi32_trunc_s_f32            Opcode = 0x00a8
	OPCi32_trunc_u_f32            Opcode = 0x00a9
	OPCi32_trunc_s_f64            Opcode = 0x00aa
	OPCi32_trunc_u_f64            Opcode = 0x00ab
	OPCi64_extend_s_i32           Opcode = 0x00ac
	OPCi64_extend_u_i32           Opcode = 0x00ad
	OPCi64_trunc_s_f32            Opcode = 0x00ae
	OPCi64_trunc_u_f32            Opcode = 0x00af
	OPCi64_trunc_s_f64            Opcode = 0x00b0
	OPCi64_trunc_u_f64            Opcode = 0x00b1
	OPCf32_convert_s_i32          Opcode = 0x00b2
	OPCf32_convert_u_i32          Opcode = 0x00b3
	OPCf32_convert_s_i64          Opcode = 0x00b4
	OPCf32_convert_u_i64          Opcode = 0x00b5
	OPCf32_demote_f64             Opcode = 0x00b6
	OPCf64_convert_s_i32          Opcode = 0x00b7
	OPCf64_convert_u_i32          Opcode = 0x00b8
	OPCf64_convert_s_i64          Opcode = 0x00b9
	OPCf64_convert_u_i64          Opcode = 0x00ba
	OPCf64_promote_f32            Opcode = 0x00bb
	OPCi32_reinterpret_f32        Opcode = 0x00bc
	OPCi64_reinterpret_f64        Opcode = 0x00bd
	OPCf32_reinterpret_i32        Opcode = 0x00be
	OPCf64_reinterpret_i64        Opcode = 0x00bf
	OPCi32_extend8_s              Opcode = 0x00c0
	OPCi32_extend16_s             Opcode = 0x00c1
	OPCi64_extend8_s              Opcode = 0x00c2
	OPCi64_extend16_s             Opcode = 0x00c3
	OPCi64_extend32_s             Opcode = 0x00c4
	OPCref_null                   Opcode = 0x00d0
	OPCref_isnull                 Opcode = 0x00d1
	OPCref_func                   Opcode = 0x00d2
	OPCi32_trunc_s_sat_f32        Opcode = 0xfc00
	OPCi32_trunc_u_sat_f32        Opcode = 0xfc01
	OPCi32_trunc_s_sat_f64        Opcode = 0xfc02
	OPCi32_trunc_u_sat_f64        Opcode = 0xfc03
	OPCi64_trunc_s_sat_f32        Opcode = 0xfc04
	OPCi64_trunc_u_sat_f32        Opcode = 0xfc05
	OPCi64_trunc_s_sat_f64        Opcode = 0xfc06
	OPCi64_trunc_u_sat_f64        Opcode = 0xfc07
	OPCmemory_init                Opcode = 0xfc08
	OPCmemory_drop                Opcode = 0xfc09
	OPCmemory_copy                Opcode = 0xfc0a
	OPCmemory_fill                Opcode = 0xfc0b
	OPCtable_init                 Opcode = 0xfc0c
	OPCtable_drop                 Opcode = 0xfc0d
	OPCtable_copy                 Opcode = 0xfc0e
	OPCv128_const                 Opcode = 0xfd00
	OPCv128_load                  Opcode = 0xfd01
	OPCv128_store                 Opcode = 0xfd02
	OPCi8x16_splat                Opcode = 0xfd03
	OPCi16x8_splat                Opcode = 0xfd04
	OPCi32x4_splat                Opcode = 0xfd05
	OPCi64x2_splat                Opcode = 0xfd06
	OPCf32x4_splat                Opcode = 0xfd07
	OPCf64x2_splat                Opcode = 0xfd08
	OPCi8x16_extract_lane_s       Opcode = 0xfd09
	OPCi8x16_extract_lane_u       Opcode = 0xfd0a
	OPCi16x8_extract_lane_s       Opcode = 0xfd0b
	OPCi16x8_extract_lane_u       Opcode = 0xfd0c
	OPCi32x4_extract_lane         Opcode = 0xfd0d
	OPCi64x2_extract_lane         Opcode = 0xfd0e
	OPCf32x4_extract_lane         Opcode = 0xfd0f
	OPCf64x2_extract_lane         Opcode = 0xfd10
	OPCi8x16_replace_lane         Opcode = 0xfd11
	OPCi16x8_replace_lane         Opcode = 0xfd12
	OPCi32x4_replace_lane         Opcode = 0xfd13
	OPCi64x2_replace_lane         Opcode = 0xfd14
	OPCf32x4_replace_lane         Opcode = 0xfd15
	OPCf64x2_replace_lane         Opcode = 0xfd16
	OPCv8x16_shuffle              Opcode = 0xfd17
	OPCi8x16_add                  Opcode = 0xfd18
	OPCi16x8_add                  Opcode = 0xfd19
	OPCi32x4_add                  Opcode = 0xfd1a
	OPCi64x2_add                  Opcode = 0xfd1b
	OPCi8x16_sub                  Opcode = 0xfd1c
	OPCi16x8_sub                  Opcode = 0xfd1d
	OPCi32x4_sub                  Opcode = 0xfd1e
	OPCi64x2_sub                  Opcode = 0xfd1f
	OPCi8x16_mul                  Opcode = 0xfd20
	OPCi16x8_mul                  Opcode = 0xfd21
	OPCi32x4_mul                  Opcode = 0xfd22
	OPCi8x16_neg                  Opcode = 0xfd24
	OPCi16x8_neg                  Opcode = 0xfd25
	OPCi32x4_neg                  Opcode = 0xfd26
	OPCi64x2_neg                  Opcode = 0xfd27
	OPCi8x16_add_saturate_s       Opcode = 0xfd28
	OPCi8x16_add_saturate_u       Opcode = 0xfd29
	OPCi16x8_add_saturate_s       Opcode = 0xfd2a
	OPCi16x8_add_saturate_u       Opcode = 0xfd2b
	OPCi8x16_sub_saturate_s       Opcode = 0xfd2c
	OPCi8x16_sub_saturate_u       Opcode = 0xfd2d
	OPCi16x8_sub_saturate_s       Opcode = 0xfd2e
	OPCi16x8_sub_saturate_u       Opcode = 0xfd2f
	OPCi8x16_shl                  Opcode = 0xfd30
	OPCi16x8_shl                  Opcode = 0xfd31
	OPCi32x4_shl                  Opcode = 0xfd32
	OPCi64x2_shl                  Opcode = 0xfd33
	OPCi8x16_shr_s                Opcode = 0xfd34
	OPCi8x16_shr_u                Opcode = 0xfd35
	OPCi16x8_shr_s                Opcode = 0xfd36
	OPCi16x8_shr_u                Opcode = 0xfd37
	OPCi32x4_shr_s                Opcode = 0xfd38
	OPCi32x4_shr_u                Opcode = 0xfd39
	OPCi64x2_shr_s                Opcode = 0xfd3a
	OPCi64x2_shr_u                Opcode = 0xfd3b
	OPCv128_and                   Opcode = 0xfd3c
	OPCv128_or                    Opcode = 0xfd3d
	OPCv128_xor                   Opcode = 0xfd3e
	OPCv128_not                   Opcode = 0xfd3f
	OPCv128_bitselect             Opcode = 0xfd40
	OPCi8x16_any_true             Opcode = 0xfd41
	OPCi16x8_any_true             Opcode = 0xfd42
	OPCi32x4_any_true             Opcode = 0xfd43
	OPCi64x2_any_true             Opcode = 0xfd44
	OPCi8x16_all_true             Opcode = 0xfd45
	OPCi16x8_all_true             Opcode = 0xfd46
	OPCi32x4_all_true             Opcode = 0xfd47
	OPCi64x2_all_true             Opcode = 0xfd48
	OPCi8x16_eq                   Opcode = 0xfd49
	OPCi16x8_eq                   Opcode = 0xfd4a
	OPCi32x4_eq                   Opcode = 0xfd4b
	OPCf32x4_eq                   Opcode = 0xfd4d
	OPCf64x2_eq                   Opcode = 0xfd4e
	OPCi8x16_ne                   Opcode = 0xfd4f
	OPCi16x8_ne                   Opcode = 0xfd50
	OPCi32x4_ne                   Opcode = 0xfd51
	OPCf32x4_ne                   Opcode = 0xfd53
	OPCf64x2_ne                   Opcode = 0xfd54
	OPCi8x16_lt_s                 Opcode = 0xfd55
	OPCi8x16_lt_u                 Opcode = 0xfd56
	OPCi16x8_lt_s                 Opcode = 0xfd57
	OPCi16x8_lt_u                 Opcode = 0xfd58
	OPCi32x4_lt_s                 Opcode = 0xfd59
	OPCi32x4_lt_u                 Opcode = 0xfd5a
	OPCf32x4_lt                   Opcode = 0xfd5d
	OPCf64x2_lt                   Opcode = 0xfd5e
	OPCi8x16_le_s                 Opcode = 0xfd5f
	OPCi8x16_le_u                 Opcode = 0xfd60
	OPCi16x8_le_s                 Opcode = 0xfd61
	OPCi16x8_le_u                 Opcode = 0xfd62
	OPCi32x4_le_s                 Opcode = 0xfd63
	OPCi32x4_le_u                 Opcode = 0xfd64
	OPCf32x4_le                   Opcode = 0xfd67
	OPCf64x2_le                   Opcode = 0xfd68
	OPCi8x16_gt_s                 Opcode = 0xfd69
	OPCi8x16_gt_u                 Opcode = 0xfd6a
	OPCi16x8_gt_s                 Opcode = 0xfd6b
	OPCi16x8_gt_u                 Opcode = 0xfd6c
	OPCi32x4_gt_s                 Opcode = 0xfd6d
	OPCi32x4_gt_u                 Opcode = 0xfd6e
	OPCf32x4_gt                   Opcode = 0xfd71
	OPCf64x2_gt                   Opcode = 0xfd72
	OPCi8x16_ge_s                 Opcode = 0xfd73
	OPCi8x16_ge_u                 Opcode = 0xfd74
	OPCi16x8_ge_s                 Opcode = 0xfd75
	OPCi16x8_ge_u                 Opcode = 0xfd76
	OPCi32x4_ge_s                 Opcode = 0xfd77
	OPCi32x4_ge_u                 Opcode = 0xfd78
	OPCf32x4_ge                   Opcode = 0xfd7b
	OPCf64x2_ge                   Opcode = 0xfd7c
	OPCf32x4_neg                  Opcode = 0xfd7d
	OPCf64x2_neg                  Opcode = 0xfd7e
	OPCf32x4_abs                  Opcode = 0xfd7f
	OPCf64x2_abs                  Opcode = 0xfd80
	OPCf32x4_min                  Opcode = 0xfd81
	OPCf64x2_min                  Opcode = 0xfd82
	OPCf32x4_max                  Opcode = 0xfd83
	OPCf64x2_max                  Opcode = 0xfd84
	OPCf32x4_add                  Opcode = 0xfd85
	OPCf64x2_add                  Opcode = 0xfd86
	OPCf32x4_sub                  Opcode = 0xfd87
	OPCf64x2_sub                  Opcode = 0xfd88
	OPCf32x4_div                  Opcode = 0xfd89
	OPCf64x2_div                  Opcode = 0xfd8a
	OPCf32x4_mul                  Opcode = 0xfd8b
	OPCf64x2_mul                  Opcode = 0xfd8c
	OPCf32x4_sqrt                 Opcode = 0xfd8d
	OPCf64x2_sqrt                 Opcode = 0xfd8e
	OPCf32x4_convert_s_i32x4      Opcode = 0xfd8f
	OPCf32x4_convert_u_i32x4      Opcode = 0xfd90
	OPCf64x2_convert_s_i64x2      Opcode = 0xfd91
	OPCf64x2_convert_u_i64x2      Opcode = 0xfd92
	OPCi32x4_trunc_s_sat_f32x4    Opcode = 0xfd93
	OPCi32x4_trunc_u_sat_f32x4    Opcode = 0xfd94
	OPCi64x2_trunc_s_sat_f64x2    Opcode = 0xfd95
	OPCi64x2_trunc_u_sat_f64x2    Opcode = 0xfd96
	OPCatomic_wake                Opcode = 0xfe00
	OPCi32_atomic_wait            Opcode = 0xfe01
	OPCi64_atomic_wait            Opcode = 0xfe02
	OPCi32_atomic_load            Opcode = 0xfe10
	OPCi64_atomic_load            Opcode = 0xfe11
	OPCi32_atomic_load8_u         Opcode = 0xfe12
	OPCi32_atomic_load16_u        Opcode = 0xfe13
	OPCi64_atomic_load8_u         Opcode = 0xfe14
	OPCi64_atomic_load16_u        Opcode = 0xfe15
	OPCi64_atomic_load32_u        Opcode = 0xfe16
	OPCi32_atomic_store           Opcode = 0xfe17
	OPCi64_atomic_store           Opcode = 0xfe18
	OPCi32_atomic_store8          Opcode = 0xfe19
	OPCi32_atomic_store16         Opcode = 0xfe1a
	OPCi64_atomic_store8          Opcode = 0xfe1b
	OPCi64_atomic_store16         Opcode = 0xfe1c
	OPCi64_atomic_store32         Opcode = 0xfe1d
	OPCi32_atomic_rmw_add         Opcode = 0xfe1e
	OPCi64_atomic_rmw_add         Opcode = 0xfe1f
	OPCi32_atomic_rmw8_u_add      Opcode = 0xfe20
	OPCi32_atomic_rmw16_u_add     Opcode = 0xfe21
	OPCi64_atomic_rmw8_u_add      Opcode = 0xfe22
	OPCi64_atomic_rmw16_u_add     Opcode = 0xfe23
	OPCi64_atomic_rmw32_u_add     Opcode = 0xfe24
	OPCi32_atomic_rmw_sub         Opcode = 0xfe25
	OPCi64_atomic_rmw_sub         Opcode = 0xfe26
	OPCi32_atomic_rmw8_u_sub      Opcode = 0xfe27
	OPCi32_atomic_rmw16_u_sub     Opcode = 0xfe28
	OPCi64_atomic_rmw8_u_sub      Opcode = 0xfe29
	OPCi64_atomic_rmw16_u_sub     Opcode = 0xfe2a
	OPCi64_atomic_rmw32_u_sub     Opcode = 0xfe2b
	OPCi32_atomic_rmw_and         Opcode = 0xfe2c
	OPCi64_atomic_rmw_and         Opcode = 0xfe2d
	OPCi32_atomic_rmw8_u_and      Opcode = 0xfe2e
	OPCi32_atomic_rmw16_u_and     Opcode = 0xfe2f
	OPCi64_atomic_rmw8_u_and      Opcode = 0xfe30
	OPCi64_atomic_rmw16_u_and     Opcode = 0xfe31
	OPCi64_atomic_rmw32_u_and     Opcode = 0xfe32
	OPCi32_atomic_rmw_or          Opcode = 0xfe33
	OPCi64_atomic_rmw_or          Opcode = 0xfe34
	OPCi32_atomic_rmw8_u_or       Opcode = 0xfe35
	OPCi32_atomic_rmw16_u_or      Opcode = 0xfe36
	OPCi64_atomic_rmw8_u_or       Opcode = 0xfe37
	OPCi64_atomic_rmw16_u_or      Opcode = 0xfe38
	OPCi64_atomic_rmw32_u_or      Opcode = 0xfe39
	OPCi32_atomic_rmw_xor         Opcode = 0xfe3a
	OPCi64_atomic_rmw_xor         Opcode = 0xfe3b
	OPCi32_atomic_rmw8_u_xor      Opcode = 0xfe3c
	OPCi32_atomic_rmw16_u_xor     Opcode = 0xfe3d
	OPCi64_atomic_rmw8_u_xor      Opcode = 0xfe3e
	OPCi64_atomic_rmw16_u_xor     Opcode = 0xfe3f
	OPCi64_atomic_rmw32_u_xor     Opcode = 0xfe40
	OPCi32_atomic_rmw_xchg        Opcode = 0xfe41
	OPCi64_atomic_rmw_xchg        Opcode = 0xfe42
	OPCi32_atomic_rmw8_u_xchg     Opcode = 0xfe43
	OPCi32_atomic_rmw16_u_xchg    Opcode = 0xfe44
	OPCi64_atomic_rmw8_u_xchg     Opcode = 0xfe45
	OPCi64_atomic_rmw16_u_xchg    Opcode = 0xfe46
	OPCi64_atomic_rmw32_u_xchg    Opcode = 0xfe47
	OPCi32_atomic_rmw_cmpxchg     Opcode = 0xfe48
	OPCi64_atomic_rmw_cmpxchg     Opcode = 0xfe49
	OPCi32_atomic_rmw8_u_cmpxchg  Opcode = 0xfe4a
	OPCi32_atomic_rmw16_u_cmpxchg Opcode = 0xfe4b
	OPCi64_atomic_rmw8_u_cmpxchg  Opcode = 0xfe4c
	OPCi64_atomic_rmw16_u_cmpxchg Opcode = 0xfe4d
	OPCi64_atomic_rmw32_u_cmpxchg Opcode = 0xfe4e
	OPCblock                      Opcode = 0x0002
	OPCloop                       Opcode = 0x0003
	OPCif_                        Opcode = 0x0004
	OPCelse_                      Opcode = 0x0005
	OPCend                        Opcode = 0x000b
	OPCtry_                       Opcode = 0xfb02
	OPCcatch_                     Opcode = 0xfb03
	OPCcatch_all                  Opcode = 0xfb04

	OPCMaxSingleByteOpcode Opcode = 0xdf
)
