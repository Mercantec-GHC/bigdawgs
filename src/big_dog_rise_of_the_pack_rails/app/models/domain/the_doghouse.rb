module Domain
  class TheDoghouse < BaseBuilding
    def initialize(building_data)
      super(building_data, 
      <<~DESC,
“The heart of the pack.”
The Doghouse is your main base and the center of your pack’s power. From here, you expand your territory, unlock new buildings, and control your progression.
👉 Sets max level for other buildings
👉 Increases overall capacity"
      DESC
      "the_doghouse.png")
    end
  end
end
# 👉 Unlocks new buildings
