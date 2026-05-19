class ResourcesFetchingService
  attr_reader :user

  def self.save_current_resources_to_cashe(user)
    new(user).save_to_cashe
  end 
  def initialize(user)
    @user = user
  end

  def resource_data
    @resource_data ||= FaradayService.fetch_data("/resources/getBag", token: user.encoded_json_web_token)
  end

  def resource_bag
    @resource_bag ||= resource_data.fetch("resourcesBag", {})
  end

  def dog_bones
   @dog_bones = resource_bag.fetch("dogbones")&.fetch("Amount", 0) || 0
  end

  def dog_coins
    @dog_coins ||= resource_bag.fetch("dogcoin")&.fetch("Amount", 0) || 0
  end

  def dogs
    @dogs ||= resource_bag.fetch("dogs")&.fetch("Amount", 0) || 0
  end

  def save_to_cashe
    Rails.cache.write([:resource_data, user.id], { "dog_bones" => dog_bones, "dog_coins" => dog_coins, "dogs" => dogs }, expires_in: 15.seconds)
  end

  def resource_cap
    @resource_cap ||= FaradayService.fetch_data("/resources/getCap", token: user.encoded_json_web_token)
  end
  
end