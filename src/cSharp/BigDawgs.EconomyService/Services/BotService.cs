using BigDawgs.EconomyService.DTOs;

namespace BigDawgs.EconomyService.Services;

public class BotService
{
    private readonly MarketService _marketService;

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
        });
    }

    private static MarketDogBoneTradeDto DecideBotTrade(decimal currentPrice)
    {
        if (currentPrice <= 4m)
        {
            return new MarketDogBoneTradeDto
            {
                Type = "buy",
                Amount = 10
            };
        }

        return new MarketDogBoneTradeDto
        {
            Type = "sell",
            Amount = 10
        };
    }
}