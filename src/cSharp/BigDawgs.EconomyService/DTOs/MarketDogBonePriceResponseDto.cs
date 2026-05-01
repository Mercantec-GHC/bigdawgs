using System.Text.Json.Serialization;

namespace BigDawgs.EconomyService.DTOs;

public class MarketDogBonePriceResponseDto
{
    [JsonPropertyName("Resources")]
    public MarketDogBonePriceDto Resources { get; set; } = new();
}

public class MarketDogBonePriceDto
{
    [JsonPropertyName("current_dog_coins_price")]
    public decimal CurrentDogCoinsPrice { get; set; }
}