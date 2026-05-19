class ResourceBarController < ApplicationController
  def show
    @resources_fetching_service = ResourcesFetchingService.new(current_user)
    @resources_fetching_service.save_to_cashe
  end
end
