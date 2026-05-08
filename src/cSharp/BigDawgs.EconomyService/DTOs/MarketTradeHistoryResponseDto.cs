namespace BigDawgs.EconomyService.DTOs;

public class MarketTradeHistoryResponseDto
{
    public List<MarketTradeHistoryDto> Resources { get; set; } = new();
}

public class MarketTradeHistoryDto
{
    public string UserId { get; set; } = string.Empty;
    public string Type { get; set; } = string.Empty;
    public int Amount { get; set; }
    public decimal PriceAtTrade { get; set; }
    public decimal TradeValue { get; set; }
    public int SupplyAfterTrade { get; set; }
    public int DemandAfterTrade { get; set; }
    public DateTime CreatedAt { get; set; }
}