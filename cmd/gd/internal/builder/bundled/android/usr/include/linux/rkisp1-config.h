/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_RKISP1_CONFIG_H
#define _UAPI_RKISP1_CONFIG_H
#include <linux/types.h>
#define RKISP1_CIF_ISP_MODULE_DPCC (1U << 0)
#define RKISP1_CIF_ISP_MODULE_BLS (1U << 1)
#define RKISP1_CIF_ISP_MODULE_SDG (1U << 2)
#define RKISP1_CIF_ISP_MODULE_HST (1U << 3)
#define RKISP1_CIF_ISP_MODULE_LSC (1U << 4)
#define RKISP1_CIF_ISP_MODULE_AWB_GAIN (1U << 5)
#define RKISP1_CIF_ISP_MODULE_FLT (1U << 6)
#define RKISP1_CIF_ISP_MODULE_BDM (1U << 7)
#define RKISP1_CIF_ISP_MODULE_CTK (1U << 8)
#define RKISP1_CIF_ISP_MODULE_GOC (1U << 9)
#define RKISP1_CIF_ISP_MODULE_CPROC (1U << 10)
#define RKISP1_CIF_ISP_MODULE_AFC (1U << 11)
#define RKISP1_CIF_ISP_MODULE_AWB (1U << 12)
#define RKISP1_CIF_ISP_MODULE_IE (1U << 13)
#define RKISP1_CIF_ISP_MODULE_AEC (1U << 14)
#define RKISP1_CIF_ISP_MODULE_WDR (1U << 15)
#define RKISP1_CIF_ISP_MODULE_DPF (1U << 16)
#define RKISP1_CIF_ISP_MODULE_DPF_STRENGTH (1U << 17)
#define RKISP1_CIF_ISP_CTK_COEFF_MAX 0x100
#define RKISP1_CIF_ISP_CTK_OFFSET_MAX 0x800
#define RKISP1_CIF_ISP_AE_MEAN_MAX_V10 25
#define RKISP1_CIF_ISP_AE_MEAN_MAX_V12 81
#define RKISP1_CIF_ISP_AE_MEAN_MAX RKISP1_CIF_ISP_AE_MEAN_MAX_V12
#define RKISP1_CIF_ISP_HIST_BIN_N_MAX_V10 16
#define RKISP1_CIF_ISP_HIST_BIN_N_MAX_V12 32
#define RKISP1_CIF_ISP_HIST_BIN_N_MAX RKISP1_CIF_ISP_HIST_BIN_N_MAX_V12
#define RKISP1_CIF_ISP_AFM_MAX_WINDOWS 3
#define RKISP1_CIF_ISP_DEGAMMA_CURVE_SIZE 17
#define RKISP1_CIF_ISP_BDM_MAX_TH 0xff
#define RKISP1_CIF_ISP_BLS_START_H_MAX 0x00000fff
#define RKISP1_CIF_ISP_BLS_STOP_H_MAX 0x00000fff
#define RKISP1_CIF_ISP_BLS_START_V_MAX 0x00000fff
#define RKISP1_CIF_ISP_BLS_STOP_V_MAX 0x00000fff
#define RKISP1_CIF_ISP_BLS_SAMPLES_MAX 0x00000012
#define RKISP1_CIF_ISP_BLS_FIX_SUB_MAX 0x00000fff
#define RKISP1_CIF_ISP_BLS_FIX_SUB_MIN 0xfffff000
#define RKISP1_CIF_ISP_BLS_FIX_MASK 0x00001fff
#define RKISP1_CIF_ISP_AWB_MAX_GRID 1
#define RKISP1_CIF_ISP_AWB_MAX_FRAMES 7
#define RKISP1_CIF_ISP_GAMMA_OUT_MAX_SAMPLES_V10 17
#define RKISP1_CIF_ISP_GAMMA_OUT_MAX_SAMPLES_V12 34
#define RKISP1_CIF_ISP_GAMMA_OUT_MAX_SAMPLES RKISP1_CIF_ISP_GAMMA_OUT_MAX_SAMPLES_V12
#define RKISP1_CIF_ISP_LSC_SECTORS_TBL_SIZE 8
#define RKISP1_CIF_ISP_LSC_SAMPLES_MAX 17
#define RKISP1_CIF_ISP_HISTOGRAM_WEIGHT_GRIDS_SIZE_V10 25
#define RKISP1_CIF_ISP_HISTOGRAM_WEIGHT_GRIDS_SIZE_V12 81
#define RKISP1_CIF_ISP_HISTOGRAM_WEIGHT_GRIDS_SIZE RKISP1_CIF_ISP_HISTOGRAM_WEIGHT_GRIDS_SIZE_V12
#define RKISP1_CIF_ISP_DPCC_METHODS_MAX 3
#define RKISP1_CIF_ISP_DPCC_MODE_STAGE1_ENABLE (1U << 2)
#define RKISP1_CIF_ISP_DPCC_OUTPUT_MODE_STAGE1_INCL_G_CENTER (1U << 0)
#define RKISP1_CIF_ISP_DPCC_OUTPUT_MODE_STAGE1_INCL_RB_CENTER (1U << 1)
#define RKISP1_CIF_ISP_DPCC_OUTPUT_MODE_STAGE1_G_3X3 (1U << 2)
#define RKISP1_CIF_ISP_DPCC_OUTPUT_MODE_STAGE1_RB_3X3 (1U << 3)
#define RKISP1_CIF_ISP_DPCC_SET_USE_STAGE1_USE_SET(n) ((n) << 0)
#define RKISP1_CIF_ISP_DPCC_SET_USE_STAGE1_USE_FIX_SET (1U << 3)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_PG_GREEN_ENABLE (1U << 0)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_LC_GREEN_ENABLE (1U << 1)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_RO_GREEN_ENABLE (1U << 2)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_RND_GREEN_ENABLE (1U << 3)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_RG_GREEN_ENABLE (1U << 4)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_PG_RED_BLUE_ENABLE (1U << 8)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_LC_RED_BLUE_ENABLE (1U << 9)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_RO_RED_BLUE_ENABLE (1U << 10)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_RND_RED_BLUE_ENABLE (1U << 11)
#define RKISP1_CIF_ISP_DPCC_METHODS_SET_RG_RED_BLUE_ENABLE (1U << 12)
#define RKISP1_CIF_ISP_DPCC_LINE_THRESH_G(v) ((v) << 0)
#define RKISP1_CIF_ISP_DPCC_LINE_THRESH_RB(v) ((v) << 8)
#define RKISP1_CIF_ISP_DPCC_LINE_MAD_FAC_G(v) ((v) << 0)
#define RKISP1_CIF_ISP_DPCC_LINE_MAD_FAC_RB(v) ((v) << 8)
#define RKISP1_CIF_ISP_DPCC_PG_FAC_G(v) ((v) << 0)
#define RKISP1_CIF_ISP_DPCC_PG_FAC_RB(v) ((v) << 8)
#define RKISP1_CIF_ISP_DPCC_RND_THRESH_G(v) ((v) << 0)
#define RKISP1_CIF_ISP_DPCC_RND_THRESH_RB(v) ((v) << 8)
#define RKISP1_CIF_ISP_DPCC_RG_FAC_G(v) ((v) << 0)
#define RKISP1_CIF_ISP_DPCC_RG_FAC_RB(v) ((v) << 8)
#define RKISP1_CIF_ISP_DPCC_RO_LIMITS_n_G(n,v) ((v) << ((n) * 4))
#define RKISP1_CIF_ISP_DPCC_RO_LIMITS_n_RB(n,v) ((v) << ((n) * 4 + 2))
#define RKISP1_CIF_ISP_DPCC_RND_OFFS_n_G(n,v) ((v) << ((n) * 4))
#define RKISP1_CIF_ISP_DPCC_RND_OFFS_n_RB(n,v) ((v) << ((n) * 4 + 2))
#define RKISP1_CIF_ISP_DPF_MAX_NLF_COEFFS 17
#define RKISP1_CIF_ISP_DPF_MAX_SPATIAL_COEFFS 6
#define RKISP1_CIF_ISP_COMPAND_NUM_POINTS 64
#define RKISP1_CIF_ISP_STAT_AWB (1U << 0)
#define RKISP1_CIF_ISP_STAT_AUTOEXP (1U << 1)
#define RKISP1_CIF_ISP_STAT_AFM (1U << 2)
#define RKISP1_CIF_ISP_STAT_HIST (1U << 3)
enum rkisp1_cif_isp_version {
  RKISP1_V10 = 10,
  RKISP1_V11,
  RKISP1_V12,
  RKISP1_V13,
  RKISP1_V_IMX8MP,
};
enum rkisp1_cif_isp_histogram_mode {
  RKISP1_CIF_ISP_HISTOGRAM_MODE_DISABLE,
  RKISP1_CIF_ISP_HISTOGRAM_MODE_RGB_COMBINED,
  RKISP1_CIF_ISP_HISTOGRAM_MODE_R_HISTOGRAM,
  RKISP1_CIF_ISP_HISTOGRAM_MODE_G_HISTOGRAM,
  RKISP1_CIF_ISP_HISTOGRAM_MODE_B_HISTOGRAM,
  RKISP1_CIF_ISP_HISTOGRAM_MODE_Y_HISTOGRAM
};
enum rkisp1_cif_isp_awb_mode_type {
  RKISP1_CIF_ISP_AWB_MODE_MANUAL,
  RKISP1_CIF_ISP_AWB_MODE_RGB,
  RKISP1_CIF_ISP_AWB_MODE_YCBCR
};
enum rkisp1_cif_isp_flt_mode {
  RKISP1_CIF_ISP_FLT_STATIC_MODE,
  RKISP1_CIF_ISP_FLT_DYNAMIC_MODE
};
enum rkisp1_cif_isp_exp_ctrl_autostop {
  RKISP1_CIF_ISP_EXP_CTRL_AUTOSTOP_0 = 0,
  RKISP1_CIF_ISP_EXP_CTRL_AUTOSTOP_1 = 1,
};
enum rkisp1_cif_isp_exp_meas_mode {
  RKISP1_CIF_ISP_EXP_MEASURING_MODE_0,
  RKISP1_CIF_ISP_EXP_MEASURING_MODE_1,
};
struct rkisp1_cif_isp_window {
  __u16 h_offs;
  __u16 v_offs;
  __u16 h_size;
  __u16 v_size;
};
struct rkisp1_cif_isp_bls_fixed_val {
  __s16 r;
  __s16 gr;
  __s16 gb;
  __s16 b;
};
struct rkisp1_cif_isp_bls_config {
  __u8 enable_auto;
  __u8 en_windows;
  struct rkisp1_cif_isp_window bls_window1;
  struct rkisp1_cif_isp_window bls_window2;
  __u8 bls_samples;
  struct rkisp1_cif_isp_bls_fixed_val fixed_val;
};
struct rkisp1_cif_isp_dpcc_methods_config {
  __u32 method;
  __u32 line_thresh;
  __u32 line_mad_fac;
  __u32 pg_fac;
  __u32 rnd_thresh;
  __u32 rg_fac;
};
struct rkisp1_cif_isp_dpcc_config {
  __u32 mode;
  __u32 output_mode;
  __u32 set_use;
  struct rkisp1_cif_isp_dpcc_methods_config methods[RKISP1_CIF_ISP_DPCC_METHODS_MAX];
  __u32 ro_limits;
  __u32 rnd_offs;
};
struct rkisp1_cif_isp_gamma_corr_curve {
  __u16 gamma_y[RKISP1_CIF_ISP_DEGAMMA_CURVE_SIZE];
};
struct rkisp1_cif_isp_gamma_curve_x_axis_pnts {
  __u32 gamma_dx0;
  __u32 gamma_dx1;
};
struct rkisp1_cif_isp_sdg_config {
  struct rkisp1_cif_isp_gamma_corr_curve curve_r;
  struct rkisp1_cif_isp_gamma_corr_curve curve_g;
  struct rkisp1_cif_isp_gamma_corr_curve curve_b;
  struct rkisp1_cif_isp_gamma_curve_x_axis_pnts xa_pnts;
};
struct rkisp1_cif_isp_lsc_config {
  __u16 r_data_tbl[RKISP1_CIF_ISP_LSC_SAMPLES_MAX][RKISP1_CIF_ISP_LSC_SAMPLES_MAX];
  __u16 gr_data_tbl[RKISP1_CIF_ISP_LSC_SAMPLES_MAX][RKISP1_CIF_ISP_LSC_SAMPLES_MAX];
  __u16 gb_data_tbl[RKISP1_CIF_ISP_LSC_SAMPLES_MAX][RKISP1_CIF_ISP_LSC_SAMPLES_MAX];
  __u16 b_data_tbl[RKISP1_CIF_ISP_LSC_SAMPLES_MAX][RKISP1_CIF_ISP_LSC_SAMPLES_MAX];
  __u16 x_grad_tbl[RKISP1_CIF_ISP_LSC_SECTORS_TBL_SIZE];
  __u16 y_grad_tbl[RKISP1_CIF_ISP_LSC_SECTORS_TBL_SIZE];
  __u16 x_size_tbl[RKISP1_CIF_ISP_LSC_SECTORS_TBL_SIZE];
  __u16 y_size_tbl[RKISP1_CIF_ISP_LSC_SECTORS_TBL_SIZE];
  __u16 config_width;
  __u16 config_height;
};
struct rkisp1_cif_isp_ie_config {
  __u16 effect;
  __u16 color_sel;
  __u16 eff_mat_1;
  __u16 eff_mat_2;
  __u16 eff_mat_3;
  __u16 eff_mat_4;
  __u16 eff_mat_5;
  __u16 eff_tint;
};
struct rkisp1_cif_isp_cproc_config {
  __u8 c_out_range;
  __u8 y_in_range;
  __u8 y_out_range;
  __u8 contrast;
  __u8 brightness;
  __u8 sat;
  __u8 hue;
};
struct rkisp1_cif_isp_awb_meas_config {
  struct rkisp1_cif_isp_window awb_wnd;
  __u32 awb_mode;
  __u8 max_y;
  __u8 min_y;
  __u8 max_csum;
  __u8 min_c;
  __u8 frames;
  __u8 awb_ref_cr;
  __u8 awb_ref_cb;
  __u8 enable_ymax_cmp;
};
struct rkisp1_cif_isp_awb_gain_config {
  __u16 gain_red;
  __u16 gain_green_r;
  __u16 gain_blue;
  __u16 gain_green_b;
};
struct rkisp1_cif_isp_flt_config {
  __u32 mode;
  __u8 grn_stage1;
  __u8 chr_h_mode;
  __u8 chr_v_mode;
  __u32 thresh_bl0;
  __u32 thresh_bl1;
  __u32 thresh_sh0;
  __u32 thresh_sh1;
  __u32 lum_weight;
  __u32 fac_sh1;
  __u32 fac_sh0;
  __u32 fac_mid;
  __u32 fac_bl0;
  __u32 fac_bl1;
};
struct rkisp1_cif_isp_bdm_config {
  __u8 demosaic_th;
};
struct rkisp1_cif_isp_ctk_config {
  __u16 coeff[3][3];
  __u16 ct_offset[3];
};
enum rkisp1_cif_isp_goc_mode {
  RKISP1_CIF_ISP_GOC_MODE_LOGARITHMIC,
  RKISP1_CIF_ISP_GOC_MODE_EQUIDISTANT
};
struct rkisp1_cif_isp_goc_config {
  __u32 mode;
  __u16 gamma_y[RKISP1_CIF_ISP_GAMMA_OUT_MAX_SAMPLES];
};
struct rkisp1_cif_isp_hst_config {
  __u32 mode;
  __u8 histogram_predivider;
  struct rkisp1_cif_isp_window meas_window;
  __u8 hist_weight[RKISP1_CIF_ISP_HISTOGRAM_WEIGHT_GRIDS_SIZE];
};
struct rkisp1_cif_isp_aec_config {
  __u32 mode;
  __u32 autostop;
  struct rkisp1_cif_isp_window meas_window;
};
struct rkisp1_cif_isp_afc_config {
  __u8 num_afm_win;
  struct rkisp1_cif_isp_window afm_win[RKISP1_CIF_ISP_AFM_MAX_WINDOWS];
  __u32 thres;
  __u32 var_shift;
};
enum rkisp1_cif_isp_dpf_gain_usage {
  RKISP1_CIF_ISP_DPF_GAIN_USAGE_DISABLED,
  RKISP1_CIF_ISP_DPF_GAIN_USAGE_NF_GAINS,
  RKISP1_CIF_ISP_DPF_GAIN_USAGE_LSC_GAINS,
  RKISP1_CIF_ISP_DPF_GAIN_USAGE_NF_LSC_GAINS,
  RKISP1_CIF_ISP_DPF_GAIN_USAGE_AWB_GAINS,
  RKISP1_CIF_ISP_DPF_GAIN_USAGE_AWB_LSC_GAINS,
  RKISP1_CIF_ISP_DPF_GAIN_USAGE_MAX
};
enum rkisp1_cif_isp_dpf_rb_filtersize {
  RKISP1_CIF_ISP_DPF_RB_FILTERSIZE_13x9,
  RKISP1_CIF_ISP_DPF_RB_FILTERSIZE_9x9,
};
enum rkisp1_cif_isp_dpf_nll_scale_mode {
  RKISP1_CIF_ISP_NLL_SCALE_LINEAR,
  RKISP1_CIF_ISP_NLL_SCALE_LOGARITHMIC,
};
struct rkisp1_cif_isp_dpf_nll {
  __u16 coeff[RKISP1_CIF_ISP_DPF_MAX_NLF_COEFFS];
  __u32 scale_mode;
};
struct rkisp1_cif_isp_dpf_rb_flt {
  __u32 fltsize;
  __u8 spatial_coeff[RKISP1_CIF_ISP_DPF_MAX_SPATIAL_COEFFS];
  __u8 r_enable;
  __u8 b_enable;
};
struct rkisp1_cif_isp_dpf_g_flt {
  __u8 spatial_coeff[RKISP1_CIF_ISP_DPF_MAX_SPATIAL_COEFFS];
  __u8 gr_enable;
  __u8 gb_enable;
};
struct rkisp1_cif_isp_dpf_gain {
  __u32 mode;
  __u16 nf_r_gain;
  __u16 nf_b_gain;
  __u16 nf_gr_gain;
  __u16 nf_gb_gain;
};
struct rkisp1_cif_isp_dpf_config {
  struct rkisp1_cif_isp_dpf_gain gain;
  struct rkisp1_cif_isp_dpf_g_flt g_flt;
  struct rkisp1_cif_isp_dpf_rb_flt rb_flt;
  struct rkisp1_cif_isp_dpf_nll nll;
};
struct rkisp1_cif_isp_dpf_strength_config {
  __u8 r;
  __u8 g;
  __u8 b;
};
struct rkisp1_cif_isp_isp_other_cfg {
  struct rkisp1_cif_isp_dpcc_config dpcc_config;
  struct rkisp1_cif_isp_bls_config bls_config;
  struct rkisp1_cif_isp_sdg_config sdg_config;
  struct rkisp1_cif_isp_lsc_config lsc_config;
  struct rkisp1_cif_isp_awb_gain_config awb_gain_config;
  struct rkisp1_cif_isp_flt_config flt_config;
  struct rkisp1_cif_isp_bdm_config bdm_config;
  struct rkisp1_cif_isp_ctk_config ctk_config;
  struct rkisp1_cif_isp_goc_config goc_config;
  struct rkisp1_cif_isp_dpf_config dpf_config;
  struct rkisp1_cif_isp_dpf_strength_config dpf_strength_config;
  struct rkisp1_cif_isp_cproc_config cproc_config;
  struct rkisp1_cif_isp_ie_config ie_config;
};
struct rkisp1_cif_isp_isp_meas_cfg {
  struct rkisp1_cif_isp_awb_meas_config awb_meas_config;
  struct rkisp1_cif_isp_hst_config hst_config;
  struct rkisp1_cif_isp_aec_config aec_config;
  struct rkisp1_cif_isp_afc_config afc_config;
};
struct rkisp1_params_cfg {
  __u32 module_en_update;
  __u32 module_ens;
  __u32 module_cfg_update;
  struct rkisp1_cif_isp_isp_meas_cfg meas;
  struct rkisp1_cif_isp_isp_other_cfg others;
};
struct rkisp1_cif_isp_compand_bls_config {
  __u32 r;
  __u32 gr;
  __u32 gb;
  __u32 b;
};
struct rkisp1_cif_isp_compand_curve_config {
  __u8 px[RKISP1_CIF_ISP_COMPAND_NUM_POINTS];
  __u32 x[RKISP1_CIF_ISP_COMPAND_NUM_POINTS];
  __u32 y[RKISP1_CIF_ISP_COMPAND_NUM_POINTS];
};
struct rkisp1_cif_isp_awb_meas {
  __u32 cnt;
  __u8 mean_y_or_g;
  __u8 mean_cb_or_b;
  __u8 mean_cr_or_r;
};
struct rkisp1_cif_isp_awb_stat {
  struct rkisp1_cif_isp_awb_meas awb_mean[RKISP1_CIF_ISP_AWB_MAX_GRID];
};
struct rkisp1_cif_isp_bls_meas_val {
  __u16 meas_r;
  __u16 meas_gr;
  __u16 meas_gb;
  __u16 meas_b;
};
struct rkisp1_cif_isp_ae_stat {
  __u8 exp_mean[RKISP1_CIF_ISP_AE_MEAN_MAX];
  struct rkisp1_cif_isp_bls_meas_val bls_val;
};
struct rkisp1_cif_isp_af_meas_val {
  __u32 sum;
  __u32 lum;
};
struct rkisp1_cif_isp_af_stat {
  struct rkisp1_cif_isp_af_meas_val window[RKISP1_CIF_ISP_AFM_MAX_WINDOWS];
};
struct rkisp1_cif_isp_hist_stat {
  __u32 hist_bins[RKISP1_CIF_ISP_HIST_BIN_N_MAX];
};
struct rkisp1_cif_isp_stat {
  struct rkisp1_cif_isp_awb_stat awb;
  struct rkisp1_cif_isp_ae_stat ae;
  struct rkisp1_cif_isp_af_stat af;
  struct rkisp1_cif_isp_hist_stat hist;
};
struct rkisp1_stat_buffer {
  __u32 meas_type;
  __u32 frame_id;
  struct rkisp1_cif_isp_stat params;
};
enum rkisp1_ext_params_block_type {
  RKISP1_EXT_PARAMS_BLOCK_TYPE_BLS,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_DPCC,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_SDG,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_AWB_GAIN,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_FLT,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_BDM,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_CTK,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_GOC,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_DPF,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_DPF_STRENGTH,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_CPROC,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_IE,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_LSC,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_AWB_MEAS,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_HST_MEAS,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_AEC_MEAS,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_AFC_MEAS,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_COMPAND_BLS,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_COMPAND_EXPAND,
  RKISP1_EXT_PARAMS_BLOCK_TYPE_COMPAND_COMPRESS,
};
#define RKISP1_EXT_PARAMS_FL_BLOCK_DISABLE (1U << 0)
#define RKISP1_EXT_PARAMS_FL_BLOCK_ENABLE (1U << 1)
struct rkisp1_ext_params_block_header {
  __u16 type;
  __u16 flags;
  __u32 size;
};
struct rkisp1_ext_params_bls_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_bls_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_dpcc_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_dpcc_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_sdg_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_sdg_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_lsc_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_lsc_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_awb_gain_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_awb_gain_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_flt_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_flt_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_bdm_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_bdm_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_ctk_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_ctk_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_goc_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_goc_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_dpf_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_dpf_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_dpf_strength_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_dpf_strength_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_cproc_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_cproc_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_ie_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_ie_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_awb_meas_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_awb_meas_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_hst_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_hst_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_aec_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_aec_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_afc_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_afc_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_compand_bls_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_compand_bls_config config;
} __attribute__((aligned(8)));
struct rkisp1_ext_params_compand_curve_config {
  struct rkisp1_ext_params_block_header header;
  struct rkisp1_cif_isp_compand_curve_config config;
} __attribute__((aligned(8)));
#define RKISP1_EXT_PARAMS_MAX_SIZE (sizeof(struct rkisp1_ext_params_bls_config) + sizeof(struct rkisp1_ext_params_dpcc_config) + sizeof(struct rkisp1_ext_params_sdg_config) + sizeof(struct rkisp1_ext_params_lsc_config) + sizeof(struct rkisp1_ext_params_awb_gain_config) + sizeof(struct rkisp1_ext_params_flt_config) + sizeof(struct rkisp1_ext_params_bdm_config) + sizeof(struct rkisp1_ext_params_ctk_config) + sizeof(struct rkisp1_ext_params_goc_config) + sizeof(struct rkisp1_ext_params_dpf_config) + sizeof(struct rkisp1_ext_params_dpf_strength_config) + sizeof(struct rkisp1_ext_params_cproc_config) + sizeof(struct rkisp1_ext_params_ie_config) + sizeof(struct rkisp1_ext_params_awb_meas_config) + sizeof(struct rkisp1_ext_params_hst_config) + sizeof(struct rkisp1_ext_params_aec_config) + sizeof(struct rkisp1_ext_params_afc_config) + sizeof(struct rkisp1_ext_params_compand_bls_config) + sizeof(struct rkisp1_ext_params_compand_curve_config) + sizeof(struct rkisp1_ext_params_compand_curve_config))
enum rksip1_ext_param_buffer_version {
  RKISP1_EXT_PARAM_BUFFER_V1 = 1,
};
struct rkisp1_ext_params_cfg {
  __u32 version;
  __u32 data_size;
  __u8 data[RKISP1_EXT_PARAMS_MAX_SIZE];
};
#endif
