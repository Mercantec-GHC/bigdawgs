using BigDawgs.EconomyService.DTOs;
using BigDawgs.EconomyService.Services;

namespace BigDawgs.EconomyService.Tests;

public class MarketServiceTests
{
    [Fact]
    public void GetPrices_Returns_Default_Resources()
    {
        var service = new MarketService();

        var result = service.GetPrices();

        Assert.NotNull(result);
        Assert.NotNull(result.Resources);
        Assert.True(result.Resources.Count >= 2);

        Assert.Contains(result.Resources, r => r.ResourceType == "BonesOfMeat");
        Assert.Contains(result.Resources, r => r.ResourceType == "DogCoins");
    }

    [Fact]
    public void CalculatePrices_Returns_Calculated_Prices_For_All_Resources()
    {
        var service = new MarketService();

        var request = new EconomyCalculationRequestDto
        {
            Resources = new List<MarketResourceInputDto>
            {
                new MarketResourceInputDto
                {
                    ResourceType = "BonesOfMeat",
                    CurrentSupply = 80,
                    CurrentDemand = 120,
                    PreviousPrice = 10m
                },
                new MarketResourceInputDto
                {
                    ResourceType = "DogCoins",
                    CurrentSupply = 150,
                    CurrentDemand = 90,
                    PreviousPrice = 5m
                }
            }
        };

        var result = service.CalculatePrices(request);

        Assert.NotNull(result);
        Assert.Equal(2, result.Resources.Count);
        Assert.Contains(result.Resources, r => r.ResourceType == "BonesOfMeat");
        Assert.Contains(result.Resources, r => r.ResourceType == "DogCoins");
    }

    [Fact]
    public void CalculatePrices_Increases_Price_When_Demand_Is_Higher_Than_Supply()
    {
        var service = new MarketService();

        var request = new EconomyCalculationRequestDto
        {
            Resources = new List<MarketResourceInputDto>
            {
                new MarketResourceInputDto
                {
                    ResourceType = "BonesOfMeat",
                    CurrentSupply = 50,
                    CurrentDemand = 150,
                    PreviousPrice = 10m
                }
            }
        };

        var result = service.CalculatePrices(request);

        var resource = result.Resources.First();

        Assert.True(resource.CurrentPrice > 10m);
    }

    [Fact]
    public void CalculatePrices_Decreases_Price_When_Supply_Is_Higher_Than_Demand()
    {
        var service = new MarketService();

        var request = new EconomyCalculationRequestDto
        {
            Resources = new List<MarketResourceInputDto>
            {
                new MarketResourceInputDto
                {
                    ResourceType = "DogCoins",
                    CurrentSupply = 200,
                    CurrentDemand = 50,
                    PreviousPrice = 5m
                }
            }
        };

        var result = service.CalculatePrices(request);

        var resource = result.Resources.First();

        Assert.True(resource.CurrentPrice < 5m);
    }

    [Fact]
    public void CalculatePrices_Never_Goes_Below_Minimum_Clamp()
    {
        var service = new MarketService();

        var request = new EconomyCalculationRequestDto
        {
            Resources = new List<MarketResourceInputDto>
            {
                new MarketResourceInputDto
                {
                    ResourceType = "BonesOfMeat",
                    CurrentSupply = 10000,
                    CurrentDemand = 1,
                    PreviousPrice = 10m
                }
            }
        };

        var result = service.CalculatePrices(request);

        var resource = result.Resources.First();

        Assert.True(resource.CurrentPrice >= 5m);
    }

    [Fact]
    public void CalculatePrices_Never_Goes_Above_Maximum_Clamp()
    {
        var service = new MarketService();

        var request = new EconomyCalculationRequestDto
        {
            Resources = new List<MarketResourceInputDto>
            {
                new MarketResourceInputDto
                {
                    ResourceType = "DogCoins",
                    CurrentSupply = 1,
                    CurrentDemand = 10000,
                    PreviousPrice = 5m
                }
            }
        };

        var result = service.CalculatePrices(request);

        var resource = result.Resources.First();

        Assert.True(resource.CurrentPrice <= 10m);
    }
}