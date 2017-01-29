class Issue < ApplicationRecord
  belongs_to :sprint
  belongs_to :board
  belongs_to :assigned_to, class_name: "User"

  scope :backlog, -> { where('board_id IS NULL AND completed_at IS NULL') }
  scope :completed, -> { where('completed_at IS NOT NULL') }
  scope :active, -> { where('completed_at IS NULL') }
end
