package shared

type DocType string

const (
	DocTypeUNKNOWN         DocType = "UNKNOWN"
	DocTypeNEEDSPREDICTION DocType = "NEEDS_PREDICTION"
	DocTypeFILLER          DocType = "FILLER"
	DocTypeFORMW2          DocType = "FORM_W2"
	DocTypeFORM1099INT     DocType = "FORM_1099_INT"
	DocTypeFORM1099DIV     DocType = "FORM_1099_DIV"
	DocTypeFORM1099MISC    DocType = "FORM_1099_MISC"
	DocTypeFORM1099B       DocType = "FORM_1099_B"
	DocTypeFORM1099        DocType = "FORM_1099"
	DocTypeFORM1099TXNS    DocType = "FORM_1099_TXNS"
	DocTypeFORM1098        DocType = "FORM_1098"
	DocTypeFORM1099R       DocType = "FORM_1099_R"
	DocTypeFORM1099NEC     DocType = "FORM_1099_NEC"
)
