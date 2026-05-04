module ViewModels
  class Buildings
    attr_reader :buildings, :building_data
    attr_accessor :kennel, :the_doghouse, :dog_bone_factory, :dog_coin_den, :market

    def initialize(buildings_data)
      @buildings = []
      @building_data = buildings_data
      create_buildings
    end

    def create_buildings
      building_data.each do |building_data|
        case building_data.fetch('Key')
        when 'kennel'
          @kennel = Domain::Kennel.new(building_data)
          buildings << @kennel
        when 'the_doghouse'
          @the_doghouse = Domain::TheDoghouse.new(building_data)
          buildings << @the_doghouse
        when 'dog_bone_factory'
          @dog_bone_factory = Domain::DogBoneFactory.new(building_data)
          buildings << @dog_bone_factory
        when 'dog_coin_den'
          @dog_coin_den = Domain::DogCoinDen.new(building_data)
          buildings << @dog_coin_den
        when 'market'
          @market = Domain::Market.new(building_data)
          buildings << @market
        end
      end
    end
    def to_json
      {
        kennel: @kennel,
        the_doghouse: @the_doghouse,
        dog_bone_factory: @dog_bone_factory,
        dog_coin_den: @dog_coin_den,
        market: @market
      }.to_json
    end
  end
end