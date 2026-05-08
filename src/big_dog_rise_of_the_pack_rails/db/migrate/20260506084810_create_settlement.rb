class CreateSettlement < ActiveRecord::Migration[8.0]
  def change
    create_table :settlements do |t|
      t.string :x
      t.string :y

      t.timestamps
    end
  end
end
