using BigDawgs.EconomyService.Services;

namespace BigDawgs.EconomyService.Tests;

public class BotServiceTests
{
    [Fact]
    public void SimulateBotTrade_Can_Influence_Market_Price()
    {
        var marketService = new MarketService();
        var botService = new BotService(marketService);

        var before = marketService.GetPrices().Resources.CurrentDogCoinsPrice;

        var result = botService.SimulateBotTrade();

        var after = result.Resources.CurrentDogCoinsPrice;

        Assert.NotEqual(before, after);
    }
}