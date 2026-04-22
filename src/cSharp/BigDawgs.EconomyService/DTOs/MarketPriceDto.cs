namespace BigDawgs.EconomyService.DTOs;

public class MarketPriceDto
{
    public string ResourceType { get; set; } = string.Empty;
    public decimal BasePrice { get; set; }
    public decimal CurrentPrice { get; set; }
    public int Supply { get; set; }
    public int Demand { get; set; }
}