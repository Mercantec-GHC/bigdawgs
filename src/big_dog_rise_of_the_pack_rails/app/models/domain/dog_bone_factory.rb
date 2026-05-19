module Domain
  class DogBoneFactory < BaseBuilding
    attr_reader :building_data
    def initialize(building_data)
      @building_data = building_data
      change_building_data
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

    def change_building_data
      building_data.merge!("key" => "dog_bone_factory")
    end 
    
    def production
      "Current production: #{building_data.fetch("production_per_tick").fetch("dog_bones")} dog bones per tick"
    end
  end
end