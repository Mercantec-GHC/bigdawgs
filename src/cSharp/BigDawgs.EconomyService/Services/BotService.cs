using BigDawgs.EconomyService.DTOs;

namespace BigDawgs.EconomyService.Services;

public class BotService
{
    private readonly MarketService _marketService;
    private static readonly Random Random = new();

    public BotService(MarketService marketService)
    {
        _marketService = marketService;
    }

    public MarketDogBonePriceResponseDto SimulateBotTrade()
    {
        var currentPrice = _marketService.GetCurrentDogCoinsPrice();

        var trade = DecideBotTrade(currentPrice);

        return _marketService.HandleTrade(new MarketDogBoneTradeRequestDto
        {
            Resources = trade
        }, "bot");
    }

    private static MarketDogBoneTradeDto DecideBotTrade(decimal currentPrice)
    {
        string type;

        if (currentPrice <= 4m)
        {
            type = Random.Next(100) < 75 ? "buy" : "sell";
        }
        else if (currentPrice >= 8m)
        {
            type = Random.Next(100) < 75 ? "sell" : "buy";
        }
        else
        {
            type = Random.Next(2) == 0 ? "buy" : "sell";
        }

        var amount = Random.Next(1, 21); 

        return new MarketDogBoneTradeDto
        {
            Type = type,
            Amount = amount
        };
    }
}