using BigDawgs.EconomyService.DTOs;
using BigDawgs.EconomyService.Models;

namespace BigDawgs.EconomyService.Services;

public class MarketService
{
    private static readonly PriceBalancingSettings Balancing = new();
    private static readonly Dictionary<string, List<decimal>> PriceHistory = new(StringComparer.OrdinalIgnoreCase);

    private static readonly Dictionary<string, decimal> BasePrices = new(StringComparer.OrdinalIgnoreCase)
    {
        ["BonesOfMeat"] = 10m,
        ["DogCoins"] = 5m
    };

    public EconomyCalculationResponseDto GetPrices()
    {
        var prices = new List<MarketPrice>
        {
            CreateDefaultPrice("BonesOfMeat"),
            CreateDefaultPrice("DogCoins")
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
                Demand = resource.CurrentDemand,
                PriceHistory = UpdatePriceHistory(resource.ResourceType, currentPrice)
            });
        }

        return MapToResponse(prices);
    }

    private static MarketPrice CreateDefaultPrice(string resourceType)
    {
        var basePrice = BasePrices[resourceType];

        return new MarketPrice
        {
            ResourceType = resourceType,
            BasePrice = basePrice,
            CurrentPrice = basePrice,
            Supply = 100,
            Demand = 100,
            PriceHistory = UpdatePriceHistory(resourceType, basePrice)
        };
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
                Demand = p.Demand,
                PriceHistory = p.PriceHistory
            }).ToList()
        };
    }

    private static decimal CalculatePrice(decimal basePrice, decimal previousPrice, int supply, int demand)
    {
        supply = Math.Max(0, supply);
        demand = Math.Max(0, demand);

        var safeSupply = Math.Max(1, supply);
        var demandSupplyRatio = (decimal)demand / safeSupply;

        var marketFactor = 1m + ((demandSupplyRatio - 1m) * Balancing.Sensitivity);
        marketFactor = Math.Clamp(
            marketFactor,
            Balancing.MinMarketFactor,
            Balancing.MaxMarketFactor
        );

        var rawPrice = previousPrice * marketFactor;

        var smoothedPrice =
            (previousPrice * Balancing.PreviousPriceWeight) +
            (rawPrice * Balancing.NewPriceWeight);

        var minPrice = basePrice * Balancing.MinPriceMultiplier;
        var maxPrice = basePrice * Balancing.MaxPriceMultiplier;

        return Math.Round(Math.Clamp(smoothedPrice, minPrice, maxPrice), 2);
    }

    private static List<decimal> UpdatePriceHistory(string resourceType, decimal currentPrice)
    {
        if (!PriceHistory.ContainsKey(resourceType))
        {
            PriceHistory[resourceType] = new List<decimal>();
        }

        PriceHistory[resourceType].Add(currentPrice);

        if (PriceHistory[resourceType].Count > Balancing.MaxHistoryEntries)
        {
            PriceHistory[resourceType].RemoveAt(0);
        }

        return new List<decimal>(PriceHistory[resourceType]);
    }

    public class PriceBalancingSettings
    {
        public decimal Sensitivity { get; init; } = 0.25m;
        public decimal MinMarketFactor { get; init; } = 0.50m;
        public decimal MaxMarketFactor { get; init; } = 2.00m;
        public decimal MinPriceMultiplier { get; init; } = 0.50m;
        public decimal MaxPriceMultiplier { get; init; } = 2.00m;

        public decimal PreviousPriceWeight { get; init; } = 0.70m;
        public decimal NewPriceWeight { get; init; } = 0.30m;
        public int MaxHistoryEntries { get; init; } = 10;
    }
}