class Board < ApplicationRecord
  belongs_to :sprint
  belongs_to :next_board
  has_many :issues
end
