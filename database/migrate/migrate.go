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
		target VARCHAR(50), -- 'product' or 'transaction'

		-- Audit Fields
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		updated_at TIMESTAMPTZ,
		updated_by UUID,
		deleted_at TIMESTAMPTZ,
		deleted_by UUID
		);
	`
}

func CreateDiscountTargetProductTableMigration() string{
	return `
		CREATE TABLE IF NOT EXISTS ndi_discount_product_target (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		discount_id UUID NOT NULL,
		target_type VARCHAR(100), -- e.g., 'category', 'brand', or 'sku'
		target_id UUID NOT NULL,   -- ID of the related product/category
		max_total_quota INTEGER,

		-- Audit Fields
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		updated_at TIMESTAMPTZ,
		updated_by UUID
		);
	`
}

func CreateDiscountTargetTransactionTableMigration() string{
	return `
		CREATE TABLE IF NOT EXISTS discount_transaction_target (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		discount_id UUID NOT NULL,
		max_total_quota INTEGER,

		-- Audit Fields
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		updated_at TIMESTAMPTZ,
		updated_by UUID
		);
	`
}

func CreateAppliedDiscountTableMigration() string{
	return `
		CREATE TABLE IF NOT EXISTS discount_transaction_target (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		discount_id UUID NOT NULL,
		max_total_quota INTEGER,

		-- Audit Fields
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		created_by UUID,
		updated_at TIMESTAMPTZ,
		updated_by UUID
		);
	`
}