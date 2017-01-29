class CreateIssues < ActiveRecord::Migration[5.0]
  def change
    create_table :issues do |t|
      t.string :title
      t.text :description
      t.integer :duration_hours
      t.references :sprint, foreign_key: true
      t.references :board, foreign_key: true
      t.references :assigned_to, foreign_key: true
      t.datetime :completed_at
      t.timestamps
    end
  end
end
