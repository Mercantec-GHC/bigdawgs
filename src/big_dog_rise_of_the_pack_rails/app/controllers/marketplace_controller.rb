class MarketplaceController < ApplicationController
  def show
    @market_place = FaradayService.fetch_data("/buildings", token: current_user.encoded_json_web_token).find{|building| building["key"] == "market" }
    @price = FaradayService.fetch_data("/market/prices",token: current_user.encoded_json_web_token, connection_type: "MARKETPLACE_BASE_URL").fetch("Resources").fetch("current_dog_coins_price")
    @dog_bones = current_user.cashed_resources&.fetch("dog_bones", 0).to_i
    @dog_coins = current_user.cashed_resources&.fetch("dog_coins", 0).to_i
  end

  def create
    if can_afford?
      process_transaction
      ResourcesFetchingService.save_current_resources_to_cashe(current_user)
    else
      flash[:error] = "You don't have enough resources to complete this transaction."
    end
    redirect_to marketplace_show_path
  end

  def market_history
    @market_history = FaradayService.fetch_data("/market/history", token: current_user.encoded_json_web_token, connection_type: "MARKETPLACE_BASE_URL").fetch("resources")
    @users = User.where(id: @market_history.map { |t| t["userId"] })
  end
  
  private

  def process_transaction
    transaction_response = FaradayService.post("/transaction/trade", token: current_user.encoded_json_web_token, body: transaction_hash(buying_type, pay_amount, selling_type, recived_amount))
    if transaction_response.success?
      response = FaradayService.post("/market/trade", token: current_user.encoded_json_web_token, connection_type: "MARKETPLACE_BASE_URL", body: { "resources": transaction_value_hash })
      if response.success?
        flash[:success] = "Transaction successful!" 
      else
        flash[:error] = "Transaction failed: #{response.body}"
      end
    else
      flash[:error] = "Transaction failed: #{transaction_response.body}"
    end
  end

  def recived_amount
    buying_type == "dogcoin" ?  pay_amount / marketplace_price_params[:price].to_f : pay_amount * marketplace_price_params[:price].to_f
  end

  def pay_amount
    marketplace_params.values.first.to_i
  end

  def transaction_value_hash
    transaction_value_hash ||= marketplace_params[:dog_bones_amount] != nil ? {type: "buy", amount: marketplace_params[:dog_bones_amount].to_i} : {type: "sell", amount: marketplace_params[:dog_coins_amount].to_i}
  end

  def transaction_hash(buy_type, buy_amount, recive_type, recive_amount)
    transaction_hash ||= {"spent"=> buy_type, "spent_amount"=> buy_amount, "receive"=> recive_type, "receive_amount"=> recive_amount.to_i}
  end

  def can_afford?
    transaction_value_hash[:type] == "buy" ? current_user.cashed_resources&.fetch("dog_bones", 0).to_i >= transaction_value_hash[:amount] : current_user.cashed_resources&.fetch("dog_coins", 0).to_i >= transaction_value_hash[:amount]
  end

  def buying_type
    marketplace_params[:dog_coins_amount] == nil ?   "dogbones" : "dogcoin"
  end

  def selling_type
    marketplace_params[:dog_coins_amount] == nil ?  "dogcoin" : "dogbones" 
  end
  
  def marketplace_params
    params.require(:marketplace).permit(:dog_coins_amount, :dog_bones_amount)
  end

  def marketplace_price_params
    params.require(:marketplace).permit(:price)
  end

end
