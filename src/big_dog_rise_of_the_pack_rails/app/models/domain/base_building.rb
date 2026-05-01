module Domain
  class BaseBuilding
    attr_reader :level, :upgrade_cost, :image_path, :name, :production, :description, :is_constructing, :construction_time_left, :bones_cost, :dog_coins_cost, :dogs_cost

    def initialize(building_data, description, image_path = nil)
      @level = building_data.fetch("Level", 0)
      @image_path = image_path
      @bones_cost = building_data.fetch("UpgradeCostDogBones", 0)
      @dog_coins_cost = building_data.fetch("UpgradeCostDogCoins", 0)
      @dogs_cost = building_data.fetch("UpgradeCostDogs", 0)
      @production = building_data.fetch("production", nil)
      @name = building_data.fetch("Key", "Unknown Building").gsub("_", " ").capitalize
      @description = description
      @is_constructing = building_data.fetch("IsConstructing", false)
      if @is_constructing
        @construction_time_left = Time.parse(building_data.fetch("CompletesAt")) - Time.now
      else
        @construction_time_left = nil
      end
    end

    def building_data_for_construcktion
      {
        "building_name" => name,
        "UpgradeCostDogBones" => bones_cost,
        "UpgradeCostDogCoins" => dog_coins_cost,
        "UpgradeCostDogs" => dogs_cost,
      }
    end
    

    def cost
      "Bones: #{bones_cost} \n Coins: #{dog_coins_cost} \n Dogs: #{dogs_cost}"
    end
    
  end
end