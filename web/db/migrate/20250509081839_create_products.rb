class CreateProducts < ActiveRecord::Migration[8.0]
  def change
    create_table :products do |t|
      t.string :name, null: false
      t.string :slug, index: { unique: true }, null: false

      t.timestamps
    end
  end
end
