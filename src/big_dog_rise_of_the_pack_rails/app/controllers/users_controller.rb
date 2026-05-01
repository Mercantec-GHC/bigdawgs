class UsersController < ApplicationController
  skip_before_action :require_login

  def new
    @user = User.new
  end

  def create
    @user = User.new(user_params)
    User.transaction do
      @user.save!
        create_account_response = FaradayService.post("/buildings/create", token: @user.encoded_json_web_token)
        create_resource_bag_response = FaradayService.post("/resources/create", token: @user.encoded_json_web_token)
        unless create_account_response.success?
          raise "Failed to create account in engine: #{create_account_response.status} - #{create_account_response.body}"
        end
        unless create_resource_bag_response.success?
          raise "Failed to create resource bag in engine: #{create_resource_bag_response.status} - #{create_resource_bag_response.body}"
        end
    end
    redirect_to session_path, notice: "User created successfully."
  rescue => e
    Rails.logger.error("User creation failed: #{e.message}")
    flash.now[:alert] = "Could not create account."
    render :new, status: :unprocessable_entity
  end

  private

  def user_params
    params.require(:user).permit(:email, :password, :password_confirmation, :name)
  end
end
