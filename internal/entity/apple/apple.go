package apple

import (
	"time"
)

// Apple Object Model
type Apple struct {
	AlamatCabAMS     string    `firestore:"AlamatCabAMS" json:"alamat_cab_ams"`
	ApoOutApoteker   string    `firestore:"ApoOut_Apoteker" json:"apo_out_apoteker"`
	ApoOutSIA        string    `firestore:"ApoOut_SIA" json:"apo_out_sia"`
	DPP              int       `firestore:"DPP" json:"dpp"` // CIMB
	Kepada           string    `firestore:"Kepada" json:"kepada"`
	KodeOutAMS       string    `firestore:"KodeOutAMS" json:"kode_out_ams"`
	NOPL             string    `firestore:"NOPL" json:"no_pl"`
	NamaCabsAMS      string    `firestore:"NamaCabsAMS" json:"nama_cabs_ams"`
	NamaGudang       string    `firestore:"NamaGudang" json:"nama_gudang"`
	NamaOut          string    `firestore:"NamaOut" json:"nama_out"`
	OrdLclNoPO       string    `firestore:"OrdLcl_NoPO" json:"ord_lcl_no_po"`
	OrdLclTglPO      string    `firestore:"OrdLcl_TglPO" json:"ord_lcl_tgl_po"`
	OutAddress       string    `firestore:"OutAddress" json:"out_address"`
	OutCode          string    `firestore:"Out_Code" json:"out_code"`
	OutName          string    `firestore:"Out_Name" json:"out_name"`
	THPDistName      string    `firestore:"THP_DistName" json:"thp_dist_name"`
	THPNoPL          string    `firestore:"THP_NoPL" json:"thp_no_pl"`
	THPNoPOD         string    `firestore:"THP_NoPOD" json:"thp_no_pod"`
	THPTglPL         time.Time `firestore:"THP_TglPL" json:"thp_tgl_pl"`
	AlamatPajak      string    `firestore:"alamatPajak" json:"alamat_pajak"` // CIMB
	Apj              string    `firestore:"apj" json:"apj"`
	ApjTujuan        string    `firestore:"apjTujuan" json:"apj_tujuan"`
	Consup           string    `firestore:"consup" json:"consup"`
	ExtraDiskon      string    `firestore:"extraDiskon" json:"extra_diskon"` // CIMB
	Finsuptop        int       `firestore:"finsup_top" json:"finsup_top"`
	IjinDari         string    `firestore:"ijinDari" json:"ijin_dari"`
	IjinTujuan       string    `firestore:"ijinTujuan" json:"ijin_tujuan"`
	NamaOutlet       string    `firestore:"namaOutlet" json:"nama_outlet"`
	NamaPajak        string    `firestore:"namaPajak" json:"nama_pajak"` // CIMB
	NpwpDari         string    `firestore:"npwpDari" json:"npwp_dari"`
	NpwpPajak        string    `firestore:"npwpPajak" json:"npwp_pajak"` // CIMB
	NpwpTujuan       string    `firestore:"npwpTujuan" json:"npwp_tujuan"`
	OutaddressDari   string    `firestore:"outaddressDari" json:"out_address_dari"`
	OutaddressTujuan string    `firestore:"outaddressTujuan" json:"out_address_tujuan"`
	OutnameDari      string    `firestore:"outnameDari" json:"out_name_dari"`
	OutnameTujuan    string    `firestore:"outnameTujuan" json:"out_name_tujuan"`
	PaymentMethod    string    `firestore:"paymentMethod" json:"payment_method"` // flag
	Pembuat          string    `firestore:"pembuat" json:"pembuat"`
	PpnPajak         int       `firestore:"ppnPajak" json:"ppn_pajak"` // CIMB
	PrintCount       int       `firestore:"printCount" json:"print_count"`
	Printed          string    `firestore:"printed" json:"printed"`
	Sika             string    `firestore:"sika" json:"sika"`
	SikaTujuan       string    `firestore:"sikaTujuan" json:"sika_tujuan"`
	Supaddress       string    `firestore:"sup_address" json:"sup_address"`
	TelpDari         string    `firestore:"telpDari" json:"telp_dari"`
	TelpTujuan       string    `firestore:"telpTujuan" json:"telp_tujuan"`
	TglTransf        string    `firestore:"tglTransf" json:"tgl_transf"`
	TotalBerat       float64   `firestore:"totalBerat" json:"total_berat"`
	TotalDiskon      float64   `firestore:"totalDiskon" json:"total_diskon"` // CIMB
	TotalHarga       int       `firestore:"totalHarga" json:"total_harga"`   // CIMB
	TotalPPN         int       `firestore:"totalPPN" json:"total_ppn"`       // CIMB
	TotalQTY         int       `firestore:"totalQTY" json:"total_qty"`
	TransFD          []transFD `json:"trans_fd"`
	TransFH          string    `firestore:"transFH" json:"trans_fh"` // flag API
}

// transFD Object Model
type transFD struct {
	ProMedUnit             int     `firestore:"Pro_MedUnit" json:"pro_med_unit"`
	ProName                string  `firestore:"Pro_Name" json:"pro_name"`
	TransFDCategoryProduct int     `firestore:"TransFD_CategoryProduct" json:"trans_fd_category_product"`
	TransfDBatchNumber     string  `firestore:"TransfD_BatchNumber" json:"trans_fd_batch_number"`
	TransfDED              string  `firestore:"TransfD_ED" json:"trans_fd_ed"`
	TransfDNoOrder         string  `firestore:"TransfD_NoOrder" json:"trans_fd_no_order"`
	TransfDNoSP            string  `firestore:"TransfD_NoSP" json:"trans_fd_no_sp"`
	TransfDNoTransf        string  `firestore:"TransfD_NoTransf" json:"trans_fd_no_transf"`
	TransfDOutCodeOrder    string  `firestore:"TransfD_OutCodeOrder" json:"trans_fd_out_code_order"`
	TransfDOutCodeSP       string  `firestore:"TransfD_OutCodeSP" json:"trans_fd_out_code_sp"`
	TransfDOutCodeTransf   string  `firestore:"TransfD_OutCodeTransf" json:"trans_fd_out_code_transf"`
	TransfDProCod          string  `firestore:"TransfD_ProCod" json:"trans_fd_pro_cod"`
	TransfDQty             int     `firestore:"TransfD_Qty" json:"trans_fd_qty"`
	TransfDQtyStk          int     `firestore:"TransfD_QtyStk" json:"trans_fd_qty_stk"`
	TransfDQtyScan         int     `firestore:"TransfD_Qty_Scan" json:"trans_fd_qty_scan"`
	TransfDUserID          string  `firestore:"TransfD_UserID" json:"trans_fd_user_id"`
	WeiHValAvg             float64 `firestore:"WeiH_ValAvg" json:"weih_val_avg"`
	WeiHValMax             float64 `firestore:"WeiH_ValMax" json:"weih_val_max"`
	WeiHValMin             float64 `firestore:"WeiH_ValMin" json:"weih_val_min"`
	Packname               string  `firestore:"pack_name" json:"pack_name"`
}
