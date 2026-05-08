class Settlement < ApplicationRecord
  belongs_to :user
  validates :x, presence: true
  validates :y, presence: true

  def location_array
    [self.x.to_i, self.y.to_i]
  end
  
end
