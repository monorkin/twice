class CreateCustomers < ActiveRecord::Migration[8.0]
  def change
    create_table :customers do |t|
      t.string :email, index: { unique: true }, null: false
      t.string :license_key, index: { unique: true }, null: false

      t.timestamps
    end
  end
end
