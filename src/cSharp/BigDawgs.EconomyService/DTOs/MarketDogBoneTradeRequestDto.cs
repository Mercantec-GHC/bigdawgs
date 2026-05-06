namespace BigDawgs.EconomyService.DTOs;

public class MarketDogBoneTradeRequestDto
{
    public MarketDogBoneTradeDto Resources { get; set; } = new();
}

public class MarketDogBoneTradeDto
{
    public string Type { get; set; } = string.Empty; 
    public int Amount { get; set; }
}