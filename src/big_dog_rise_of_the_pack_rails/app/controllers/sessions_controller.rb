class SessionsController < ApplicationController
  skip_before_action :require_login
  
  def show
  end

  def create
    puts "Attempting to log in with email: #{session_params[:email]}"
    user = User.find_by(email: session_params[:email])
    if user&.authenticate(session_params[:password])
      session[:user_id] = user.id
      redirect_to root_path, notice: "Logged in successfully."
    else
      flash.now[:alert] = "Invalid email or password."
      render :show, status: :unprocessable_entity
    end
  end

  private

  def session_params
    params.require(:session).permit(:email, :password)
  end
  
end
