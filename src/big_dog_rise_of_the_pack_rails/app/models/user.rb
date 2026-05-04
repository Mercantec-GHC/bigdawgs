class User < ApplicationRecord
  has_secure_password

  validates :email, presence: true, uniqueness: true
  validates :name, presence: true
  validates :password, length: { minimum: 6 }, if: -> { password.present? }
  validates :password_confirmation, presence: true, if: -> { password.present? }
  validates :name, length: { maximum: 50 }, allow_blank: false, uniqueness: true


  def encoded_json_web_token
    token = JsonWebTokenService.encode({ user_id: id, email: email, iss: "big_dog_rails", exp: 5.minutes.from_now.to_i })
  end
  
  def cashed_resources
    Rails.cache.read([:resource_data, id])
  end
  
end
