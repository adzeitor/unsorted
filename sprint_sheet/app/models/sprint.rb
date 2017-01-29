class Sprint < ApplicationRecord
  has_many :boards
  has_many :issues
end
