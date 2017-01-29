class CreateSprints < ActiveRecord::Migration[5.0]
  def change
    create_table :sprints do |t|
      t.string :title
      t.text :description
      t.datetime :due_to
      t.datetime :completed_at

      t.timestamps
    end
  end
end
