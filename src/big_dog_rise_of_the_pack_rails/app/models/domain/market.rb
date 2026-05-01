module Domain
  class Market < BaseBuilding
    def initialize(building_data)
      super(building_data, 
      <<~DESC,
“Trade, thrive, dominate.”
The Market connects your pack to the outside world. Trade resources, optimize your economy, and turn surplus into profit.
👉 Enables resource trading
        DESC
        "market.png"
        )
    end
  end
end