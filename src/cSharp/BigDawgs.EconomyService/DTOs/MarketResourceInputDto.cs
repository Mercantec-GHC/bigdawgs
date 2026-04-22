namespace BigDawgs.EconomyService.DTOs;

public class MarketResourceInputDto
{
    public string ResourceType { get; set; } = string.Empty;
    public int CurrentSupply { get; set; }
    public int CurrentDemand { get; set; }
    public decimal PreviousPrice { get; set; }
}