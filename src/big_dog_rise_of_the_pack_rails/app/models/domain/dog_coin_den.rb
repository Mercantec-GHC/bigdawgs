module Domain
  class DogCoinDen < BaseBuilding
    attr_reader :building_data
    def initialize(building_data)
      @building_data = building_data
      super(building_data, 
      <<~DESC,
“Where the real deals are buried.”
Not everything here is strictly legal—but it works. The Dog Coin Den generates coins through trade, deals, and maybe a few shady operations behind the scenes.
👉 Produces: Coins
        DESC
        "dog_coin_den.png"
        )
    end

    def production
      "Current production: #{building_data.fetch("production_per_tick").fetch("dog_coins")} dog coins per tick"
    end
  end
end