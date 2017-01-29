class Board < ApplicationRecord
  belongs_to :sprint
  belongs_to :next_board
  has_many :issues

  def issues_duration
    self.issues.map(&:duration_hours).reduce(&:+) || 0
  end
end
