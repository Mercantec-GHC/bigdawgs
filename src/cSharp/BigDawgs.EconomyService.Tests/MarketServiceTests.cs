using BigDawgs.EconomyService.DTOs;
using BigDawgs.EconomyService.Services;

namespace BigDawgs.EconomyService.Tests;

public class MarketServiceTests
{
    [Fact]
    public void GetPrices_Returns_Current_DogCoins_Price()
    {
        var service = new MarketService();

        var result = service.GetPrices();

        Assert.NotNull(result);
        Assert.NotNull(result.Resources);
        Assert.Equal(5m, result.Resources.CurrentDogCoinsPrice);
    }

    [Fact]
    public void HandleTrade_Increases_Price_When_Bot_Or_Player_Buys()
    {
        var service = new MarketService();

        var before = service.GetPrices().Resources.CurrentDogCoinsPrice;

        var request = new MarketDogBoneTradeRequestDto
        {
            Resources = new MarketDogBoneTradeDto
            {
                Type = "buy",
                Amount = 100
            }
        };

        var result = service.HandleTrade(request);

        Assert.True(result.Resources.CurrentDogCoinsPrice > before);
    }

    [Fact]
    public void HandleTrade_Decreases_Price_When_Bot_Or_Player_Sells()
    {
        var service = new MarketService();

        var before = service.GetPrices().Resources.CurrentDogCoinsPrice;

        var result = service.HandleTrade(new MarketDogBoneTradeRequestDto
        {
            Resources = new MarketDogBoneTradeDto
            {
                Type = "sell",
                Amount = 100
            }
        });

        Assert.True(result.Resources.CurrentDogCoinsPrice < before);
    }

    [Fact]
    public void HandleTrade_Does_Not_Go_Below_Minimum_Price()
    {
        var service = new MarketService();

        for (var i = 0; i < 20; i++)
        {
            service.HandleTrade(new MarketDogBoneTradeRequestDto
            {
                Resources = new MarketDogBoneTradeDto
                {
                    Type = "sell",
                    Amount = 10000
                }
            });
        }

        var result = service.GetPrices();

        Assert.True(result.Resources.CurrentDogCoinsPrice >= 2.5m);
    }

    [Fact]
    public void HandleTrade_Does_Not_Go_Above_Maximum_Price()
    {
        var service = new MarketService();

        for (var i = 0; i < 20; i++)
        {
            service.HandleTrade(new MarketDogBoneTradeRequestDto
            {
                Resources = new MarketDogBoneTradeDto
                {
                    Type = "buy",
                    Amount = 10000
                }
            });
        }

        var result = service.GetPrices();

        Assert.True(result.Resources.CurrentDogCoinsPrice <= 10m);
    }

    [Fact]
    public void HandleTrade_Is_Deterministic_For_Same_Input()
    {
        var firstService = new MarketService();
        var secondService = new MarketService();

        var request = new MarketDogBoneTradeRequestDto
        {
            Resources = new MarketDogBoneTradeDto
            {
                Type = "buy",
                Amount = 100
            }
        };

        var first = firstService.HandleTrade(request);
        var second = secondService.HandleTrade(request);

        Assert.Equal(
            first.Resources.CurrentDogCoinsPrice,
            second.Resources.CurrentDogCoinsPrice
        );
    }

    [Fact]
    public void HandleTrade_Rejects_Invalid_Trade_Type()
    {
        var service = new MarketService();

        var request = new MarketDogBoneTradeRequestDto
        {
            Resources = new MarketDogBoneTradeDto
            {
                Type = "trade",
                Amount = 100
            }
        };

        Assert.Throws<ArgumentException>(() => service.HandleTrade(request));
    }

    [Fact]
    public void HandleTrade_Rejects_Zero_Amount()
    {
        var service = new MarketService();

        var request = new MarketDogBoneTradeRequestDto
        {
            Resources = new MarketDogBoneTradeDto
            {
                Type = "buy",
                Amount = 0
            }
        };

        Assert.Throws<ArgumentException>(() => service.HandleTrade(request));
    }

    [Fact]
    public void HandleTrade_Trims_And_Lowercases_Type()
    {
        var service = new MarketService();

        var request = new MarketDogBoneTradeRequestDto
        {
            Resources = new MarketDogBoneTradeDto
            {
                Type = " BUY ",
                Amount = 100
            }
        };

        var result = service.HandleTrade(request);

        Assert.True(result.Resources.CurrentDogCoinsPrice > 5m);
    }
}