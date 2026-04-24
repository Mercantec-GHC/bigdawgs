class UsersController < ApplicationController
  skip_before_action :require_login

  def new
    @user = User.new
  end

  def create
    @user = User.new(user_params)
    User.transaction do
      @user.save!
      raise StandardError unless FaradayService.post_with_success("https://bd-engine.mags.dk/users/#{@user.id}/buildings/create", token: @user.encoded_json_web_token)
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
