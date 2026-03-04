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
		CreateProductDiscountAppliedTable(),
		CreateTransactionDiscountAppliedTable(),
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

func CreateProductDiscountAppliedTable() Migration {
	return Migration{
		Name: "create_product_discount_applied_table",
		SQL:  CreateProductDiscountAppliedTableMigration(),
	}
}

func CreateTransactionDiscountAppliedTable() Migration {
	return Migration{
		Name: "create_transaction_discount_applied_table",
		SQL:  CreateTransactionDiscountAppliedTableMigration(),
	}
}
