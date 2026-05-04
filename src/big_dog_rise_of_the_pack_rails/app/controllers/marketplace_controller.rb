class MarketplaceController < ApplicationController
  def show
    @market_place =  JSON.parse(File.read(Rails.root.join("buildings.json"))).find { |hash| hash["Key"] == "market"}
    @price = FaradayService.fetch_data("/market/prices", connection_type: "MARKETPLACE_BASE_URL").fetch("Resources").fetch("current_dog_coins_price")
  end

  def create
  end
end
