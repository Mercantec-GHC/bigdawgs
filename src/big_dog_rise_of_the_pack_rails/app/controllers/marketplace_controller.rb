class MarketplaceController < ApplicationController
  def show
    @market_place =  JSON.parse(File.read(Rails.root.join("buildings.json"))).find { |hash| hash["Key"] == "market"}
    @price = FaradayService.fetch_data("/market/prices", connection_type: "MARKETPLACE_BASE_URL")
    puts "______________________________"
    puts "MARKETPLACE PRICE: #{@price}"
    puts "______________________________"

  end

  def create
  end
end
