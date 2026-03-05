package migrate

func CreatePgcryptoExtensionMigration() string {
	return `
	    CREATE EXTENSION IF NOT EXISTS pgcrypto;
	`
}

func CreateTableMigration() string {
	return `
	CREATE TABLE IF NOT EXISTS migrations (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL,
		migrated_at TIMESTAMPTZ NOT NULL DEFAULT now()
	);
	`
}

func CreateDiscountTableMigration() string {
	return `
		CREATE TABLE IF NOT EXISTS ndi_discount (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		type VARCHAR(50) NOT NULL, -- 'percentage' or 'nominal'
		value DECIMAL(15, 2) NOT NULL, -- percentage 10 for 10% or nominal 5000 for $5,000
		start_date TIMESTAMPTZ NOT NULL,
		end_date TIMESTAMPTZ NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		updated_at TIMESTAMPTZ,
		updated_by UUID,
		deleted_at TIMESTAMPTZ,
		deleted_by UUID
		);
	`
}

func CreateDiscountTargetProductTableMigration() string {
	return `
		CREATE TABLE IF NOT EXISTS ndi_discount_product_target (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		discount_id UUID NOT NULL,
		target_id UUID NOT NULL, -- Merujuk pada ID Produk
		product_name VARCHAR(255),
		max_total_quota INTEGER,
		price_before_discount DECIMAL(15, 2),
		total_discount DECIMAL(15, 2),
		price_after_discount DECIMAL(15, 2),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		updated_at TIMESTAMPTZ,
		updated_by UUID,
		deleted_at TIMESTAMPTZ,
		deleted_by UUID
	);
	`
}

func CreateDiscountTargetTransactionTableMigration() string {
	return `
		CREATE TABLE IF NOT EXISTS ndi_discount_transaction_target (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		discount_id UUID NOT NULL,
		max_total_quota INTEGER,
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		updated_at TIMESTAMPTZ,
		updated_by UUID,
		deleted_at TIMESTAMPTZ,
		deleted_by UUID
	);
	`
}

func CreateProductDiscountAppliedTableMigration() string {
	return `
		CREATE TABLE IF NOT EXISTS ndi_discount_product_applied (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		discount_product_target_id UUID,
		customer_id UUID,
		customer_name VARCHAR(255),
		transaction_date TIMESTAMPTZ,
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		deleted_at TIMESTAMPTZ,
		deleted_by UUID
	);
	`
}

func CreateTransactionDiscountAppliedTableMigration() string {
	return `
		CREATE TABLE IF NOT EXISTS ndi_discount_transaction_applied (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		discount_transaction_target_id UUID NOT NULL,
		target_id UUID, -- Biasanya merujuk pada Transaction ID
		price_before_discount DECIMAL(15, 2),
		total_discount DECIMAL(15, 2),
		price_after_discount DECIMAL(15, 2),
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		deleted_at TIMESTAMPTZ,
		deleted_by UUID
	);
	`
}
