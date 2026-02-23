package migrate

type Migration struct {
	Name string
	SQL  string
}

// Daftar semua migrasi
func GetMigrations() []Migration {
	return []Migration{
		CreatePgcryptoExtension(),
		CreateDiscountTable(),
		CreateDiscountTargetProductTable(),
		CreateDiscountTargetTransactionTable(),
		CreateAppliedDiscountTable(),
	}
}

func CreatePgcryptoExtension() Migration {
	return Migration{
		Name: "create_pgcrypto_extension",
		SQL:  CreatePgcryptoExtensionMigration(),
	}
}


func CreateDiscountTable() Migration {
	return Migration{
		Name: "create_discount_table",
		SQL:  CreateDiscountTableMigration(),
	}
}

func CreateDiscountTargetProductTable() Migration {
	return Migration{
		Name: "create_discount_product_target_table",
		SQL:  CreateDiscountTargetProductTableMigration(),
	}
}

func CreateDiscountTargetTransactionTable() Migration {
	return Migration{
		Name: "create_discount_transaction_target_table",
		SQL:  CreateDiscountTargetTransactionTableMigration(),
	}
}

func CreateAppliedDiscountTable() Migration {
	return Migration{
		Name: "create_applied_discount_table",
		SQL:  CreateAppliedDiscountTableMigration(),
	}
}
