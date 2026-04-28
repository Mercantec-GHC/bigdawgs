class ApplicationController < ActionController::Base
  allow_browser versions: :modern
  helper_method :current_user
  before_action :require_login

  def current_user
    @current_user ||= User.find_by(id: session[:user_id])
  end

  def require_login
    redirect_to session_path unless current_user
  end
end
