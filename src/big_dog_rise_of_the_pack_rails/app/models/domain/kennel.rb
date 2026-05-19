module Domain
  class Kennel < BaseBuilding
    attr_reader :building_data  
    def initialize(building_data)
      @building_data = building_data
      super(building_data, 
      <<~DESC,
“Grow your pack. Expand your power.”
The Kennel is where your dogs are trained, housed, and prepared for expansion. A bigger pack means more strength, more production, and greater dominance.
👉 Increases number of dogs
        DESC
        "kennel.png"
        )
    end
    def production
      "Current production: #{building_data.fetch("production_per_tick").fetch("dogs")} dogs per tick."
    end
  end
end