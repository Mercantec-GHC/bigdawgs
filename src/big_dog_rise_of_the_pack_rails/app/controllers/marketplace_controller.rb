class MarketplaceController < ApplicationController
  def show
    @market_place =  JSON.parse(File.read(Rails.root.join("buildings.json"))).find { |hash| hash["Key"] == "market"}
    puts "Market place data: #{@market_place.fetch("Level")}"
  end

  def create
  end
end
