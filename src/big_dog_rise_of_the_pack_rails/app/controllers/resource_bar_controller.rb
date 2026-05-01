class ResourceBarController < ApplicationController
  def show
    resource_data = FaradayService.fetch_data("/resources/getBag", token: current_user.encoded_json_web_token)
    resource_bag = resource_data.fetch("resourcesBag", {})
    @dog_bones = resource_bag.fetch("dogbones")&.fetch("Amount", 0) || 0
    @dog_coins = resource_bag.fetch("dogcoin")&.fetch("Amount", 0) || 0
    @dogs = resource_bag.fetch("dogs")&.fetch("Amount", 0) || 0

    Rails.cache.write([:resource_data, current_user.id], { dog_bones: @dog_bones, dog_coins: @dog_coins, dogs: @dogs }, expires_in: 15.seconds)
  end
end
