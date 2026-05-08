class BuildingController < ApplicationController
  def show
  end

  def update

  end

  def index
    buildings = FaradayService.fetch_data("/buildings", token: current_user.encoded_json_web_token)
    @buildings_view_model = ViewModels::Buildings.new(buildings)
  end

  def destroy

  end

  def delete

  end

  def create
    if can_build?
      building_name = create_params[:building_name] 
      building_name = "meat_factory" if building_name == "Dog bone factory"
      response = FaradayService.post("/buildings/#{building_name.downcase.gsub(' ', '_')}/upgrade", token: current_user.encoded_json_web_token, connection_type: "ENGINE_BASE_URL")
      if response.success?
        flash[:success] = "Building upgrade started successfully!"
      else
        flash[:error] = "Failed to upgrade building. Please try again."
      end
      redirect_to building_index_path
    else
      flash[:error] = "You don't have enough resources to upgrade this building."
      redirect_to building_index_path
    end
  end
  

  private

  def can_build?
    create_params[:UpgradeCostDogBones].to_i <= current_user.cashed_resources&.fetch("dog_bones", 0).to_i &&
    create_params[:UpgradeCostDogCoins].to_i <= current_user.cashed_resources&.fetch("dog_coins", 0).to_i &&
    create_params[:UpgradeCostDogs].to_i <= current_user.cashed_resources&.fetch("dogs", 0).to_i 
  end

  def create_params
    params.require(:building).permit(:UpgradeCostDogBones, :UpgradeCostDogCoins, :UpgradeCostDogs, :building_name)
  end
  
end
