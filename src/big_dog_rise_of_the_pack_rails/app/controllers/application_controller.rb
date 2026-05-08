class ApplicationController < ActionController::Base
  allow_browser versions: :modern
  helper_method :current_user
  before_action :require_login

  def current_user
    @current_user ||= User.find_by(id: session[:user_id])
  end

  def require_login
    if current_user
      if current_user.settlements.empty?
        redirect_to new_map_path
      end
    else
      redirect_to session_path 
    end
  end
end
