module Domain
  class DogBoneFactory < BaseBuilding
    def initialize(building_data)
      change_building_data(building_data)
      super(building_data,
      <<~DESC,
“Where bones become power.”
Your pack can’t survive on loyalty alone. The Dog Bone Factory produces a steady supply of Dog bones, fueling your dogs and keeping them strong, focused, and ready to grow.
👉 Produces: Dog bones
👉 Higher level → higher production
        DESC
        "dog_bone_factory.png"
        )
    end

    def change_building_data(building_data)
      building_data.merge!("Key" => "dog_bone_factory", "UpgradeCostDogBones" => 100000, "UpgradeCostDogCoins" => 500000, "UpgradeCostDogs" => 10000)
    end 
  end
end