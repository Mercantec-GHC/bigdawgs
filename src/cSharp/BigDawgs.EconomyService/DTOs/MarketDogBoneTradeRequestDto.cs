namespace BigDawgs.EconomyService.DTOs;

public class MarketDogBoneTradeRequestDto
{
    public MarketDogBoneTradeDto Resources { get; set; } = new();
}

public class MarketDogBoneTradeDto
{
    public string Type { get; set; } = string.Empty; // buy or sell
    public int Amount { get; set; }
}