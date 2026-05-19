class MapsController < ApplicationController
  skip_before_action :require_login, only: [:new, :tile_popup, :create]
  before_action :all_players_locations, only: [:show, :new]

  def show
  end

  def new
  end

  def create
    new_Settlement = current_user.settlements.new(x: params[:x], y: params[:y])
    if new_Settlement.valid?
      new_Settlement.save
      flash[:success] = "Settlement started welcome pack"
    end
    redirect_to root_path
  end

  def tile_popup
    @x = params[:x]
    @y = params[:y]
    settlement = Settlement.find_by(x: @x, y: @y)
    if settlement
      @free = false
      player = settlement.user
      @dog_house_level = FaradayService.fetch_data("/buildings/the_doghouse/#{player.id}", token: current_user.encoded_json_web_token)&.fetch("building")&.fetch("level", 1)
      player.id == current_user.id ? @base_name = "Your" : @base_name = "#{player.name}s"   
    else
      @free = true
    end

  end

  def all_players_locations
    @all_players_locations ||= Settlement.all&.map do |settlement|
      [settlement.x.to_i, settlement.y.to_i]
    end
  end
end
