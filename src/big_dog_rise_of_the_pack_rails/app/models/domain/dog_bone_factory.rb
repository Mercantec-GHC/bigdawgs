module Domain
  class DogBoneFactory < BaseBuilding
    def initialize(building_data)
      super(building_data,
      <<~DESC,
“Where bones become power.”
Your pack can’t survive on loyalty alone. The Dog Bone Factory produces a steady supply of Bones of Meat, fueling your dogs and keeping them strong, focused, and ready to grow.
👉 Produces: Bones of Meat
👉 Higher level → higher production
        DESC
        "dog_bone_factory.png"
        )
    end
  end
end