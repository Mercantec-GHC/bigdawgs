# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[8.0].define(version: 2026_04_22_082852) do
  # These are extensions that must be enabled in order to support this database
  enable_extension "pg_catalog.plpgsql"

  create_table "buildings", force: :cascade do |t|
    t.timestamptz "created_at"
    t.timestamptz "updated_at"
    t.timestamptz "deleted_at"
    t.bigint "user_id", null: false
    t.text "key", null: false
    t.bigint "level", default: 1, null: false
    t.bigint "count", default: 1, null: false
    t.boolean "is_constructing", default: false, null: false
    t.timestamptz "started_at"
    t.timestamptz "completes_at"
    t.index ["deleted_at"], name: "idx_buildings_deleted_at"
    t.index ["user_id", "key"], name: "idx_user_building_key", unique: true
  end

  create_table "resource_balances", force: :cascade do |t|
    t.timestamptz "created_at"
    t.timestamptz "updated_at"
    t.timestamptz "deleted_at"
    t.bigint "user_id", null: false
    t.text "resource_key", null: false
    t.bigint "amount", default: 0, null: false
    t.bigint "capacity", default: 0, null: false
    t.index ["deleted_at"], name: "idx_resource_balances_deleted_at"
    t.index ["resource_key"], name: "idx_user_resource_key", unique: true
    t.index ["user_id"], name: "idx_resource_balances_user_id"
  end

  create_table "users", force: :cascade do |t|
    t.string "name"
    t.string "email"
    t.string "password_digest"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end
end
