class BuildingController < ApplicationController
  def show
  end

  def update

  end

  def index
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
