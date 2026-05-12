class MarketplaceController < ApplicationController
  def show
    @market_place = buildings = FaradayService.fetch_data("/buildings", token: current_user.encoded_json_web_token).find{|building| building["Key"] == "market" }
    @price = FaradayService.fetch_data("/market/prices",token: current_user.encoded_json_web_token, connection_type: "MARKETPLACE_BASE_URL").fetch("Resources").fetch("current_dog_coins_price")
    @dog_bones = current_user.cashed_resources&.fetch("dog_bones", 0).to_i
    @dog_coins = current_user.cashed_resources&.fetch("dog_coins", 0).to_i
  end

  def create
    if can_afford?
      process_transaction
    else
      flash[:error] = "You don't have enough resources to complete this transaction."
    end
    redirect_to marketplace_show_path
  end

  def market_history
    
  end
  

  private

  def process_transaction
    response = FaradayService.post("/market/trade", token: current_user.encoded_json_web_token, connection_type: "MARKETPLACE_BASE_URL", body: { "resources": transaction_hash })
    if response.success?
      flash[:success] = "Transaction successful!" 
    else
      flash[:error] = "Transaction failed: #{response.body}"
    end
  end

  def transaction_hash
    transaction_hash ||= marketplace_params[:dog_bones_amount] != nil ? {type: "buy", amount: marketplace_params[:dog_bones_amount].to_i} : {type: "sell", amount: marketplace_params[:dog_coins_amount].to_i}
  end

  def can_afford?
    transaction_hash[:type] == "buy" ? current_user.cashed_resources&.fetch("dog_bones", 0).to_i >= transaction_hash[:amount] : current_user.cashed_resources&.fetch("dog_coins", 0).to_i >= transaction_hash[:amount]
  end

  def marketplace_params
    params.require(:marketplace).permit(:dog_coins_amount, :dog_bones_amount)
  end
end
