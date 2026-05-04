module Domain
  class Kennel < BaseBuilding  
    def initialize(building_data)
      super(building_data, 
      <<~DESC,
“Grow your pack. Expand your power.”
The Kennel is where your dogs are trained, housed, and prepared for expansion. A bigger pack means more strength, more production, and greater dominance.
👉 Increases number of dogs
        DESC
        "kennel.png"
        )
    end
  end
end