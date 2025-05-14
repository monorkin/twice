class CreateProducts < ActiveRecord::Migration[8.0]
  def change
    create_table :products do |t|
      t.string :name, null: false
      t.string :repository, index: { unique: true }, null: false

      t.timestamps
    end

    create_table :licenses do |t|
      t.string :key, index: { unique: true }, null: false
      t.belongs_to :product, null: false
      t.belongs_to :owner, null: false, foreign_key: { to_table: :users }

      t.timestamps
    end
  end
end
