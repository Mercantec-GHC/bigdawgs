class BuildingController < ApplicationController
  def show
  end

  def update

  end

  def index
    building_test = FaradayService.fetch_data("/buildings", token: current_user.encoded_json_web_token)
    buildings = File.read(Rails.root.join('buildings.json'))
    @buildings_view_model = ViewModels::Buildings.new(JSON.parse(buildings))
  end

  def destroy

  end

  def delete

  end

  def create
    puts "Creating building with params: #{params[:building]}"
  end
  
end
