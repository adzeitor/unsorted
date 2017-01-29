class CreateBoards < ActiveRecord::Migration[5.0]
  def change
    create_table :boards do |t|
      t.string :title
      t.text :description
      t.references :sprint, foreign_key: true
      t.references :next_board, foreign_key: true

      t.timestamps
    end
  end
end
