using BigDawgs.EconomyService.DTOs;
using BigDawgs.EconomyService.Models;

namespace BigDawgs.EconomyService.Services;

public class MarketService
{
    private static readonly Dictionary<string, decimal> BasePrices = new(StringComparer.OrdinalIgnoreCase)
    {
        ["BonesOfMeat"] = 10m,
        ["DogCoins"] = 5m
    };

    public EconomyCalculationResponseDto GetPrices()
    {
        var prices = new List<MarketPrice>
        {
            new MarketPrice
            {
                ResourceType = "BonesOfMeat",
                BasePrice = 10m,
                CurrentPrice = 10m,
                Supply = 100,
                Demand = 100
            },
            new MarketPrice
            {
                ResourceType = "DogCoins",
                BasePrice = 5m,
                CurrentPrice = 5m,
                Supply = 100,
                Demand = 100
            }
        };

        return MapToResponse(prices);
    }

    public EconomyCalculationResponseDto CalculatePrices(EconomyCalculationRequestDto request)
    {
        var prices = new List<MarketPrice>();

        foreach (var resource in request.Resources)
        {
            var basePrice = BasePrices.TryGetValue(resource.ResourceType, out var configuredBasePrice)
                ? configuredBasePrice
                : resource.PreviousPrice;

            var currentPrice = CalculatePrice(
                basePrice,
                resource.PreviousPrice,
                resource.CurrentSupply,
                resource.CurrentDemand
            );

            prices.Add(new MarketPrice
            {
                ResourceType = resource.ResourceType,
                BasePrice = basePrice,
                CurrentPrice = currentPrice,
                Supply = resource.CurrentSupply,
                Demand = resource.CurrentDemand
            });
        }

        return MapToResponse(prices);
    }

    private static EconomyCalculationResponseDto MapToResponse(List<MarketPrice> prices)
    {
        return new EconomyCalculationResponseDto
        {
            Resources = prices.Select(p => new MarketPriceDto
            {
                ResourceType = p.ResourceType,
                BasePrice = p.BasePrice,
                CurrentPrice = p.CurrentPrice,
                Supply = p.Supply,
                Demand = p.Demand
            }).ToList()
        };
    }

    private static decimal CalculatePrice(decimal basePrice, decimal previousPrice, int supply, int demand)
    {
        if (supply < 0) supply = 0;
        if (demand < 0) demand = 0;

        var safeSupply = Math.Max(1, supply);
        var ratio = (decimal)demand / safeSupply;

        var marketFactor = 1m + ((ratio - 1m) * 0.25m);
        marketFactor = Math.Clamp(marketFactor, 0.50m, 2.00m);

        var rawPrice = previousPrice * marketFactor;

        var minPrice = basePrice * 0.50m;
        var maxPrice = basePrice * 2.00m;

        return Math.Round(Math.Clamp(rawPrice, minPrice, maxPrice), 2);
    }
}