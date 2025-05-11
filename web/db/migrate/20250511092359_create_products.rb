class CreateProducts < ActiveRecord::Migration[8.0]
  def change
    create_table :products do |t|
      t.string :name, null: false
      t.string :repository, index: { unique: true }, null: false

      t.timestamps
    end

    create_join_table :products, :users
  end
end
