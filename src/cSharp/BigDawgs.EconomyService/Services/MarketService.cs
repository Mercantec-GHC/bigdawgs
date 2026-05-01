using BigDawgs.EconomyService.DTOs;

namespace BigDawgs.EconomyService.Services;

public class MarketService
{
    private static readonly PriceBalancingSettings Balancing = new();

    private const string DogCoins = "DogCoins";

    private static readonly Dictionary<string, List<decimal>> PriceHistory = new(StringComparer.OrdinalIgnoreCase);

    private static readonly Dictionary<string, decimal> BasePrices = new(StringComparer.OrdinalIgnoreCase)
    {
        [DogCoins] = 5m
    };

    private int _dogBoneSupply = 100;
    private int _dogBoneDemand = 100;
    private decimal _currentDogCoinsPrice = 5m;

    private readonly List<MarketTradeHistory> _tradeHistory = new();

    public MarketDogBonePriceResponseDto GetPrices()
    {
        return new MarketDogBonePriceResponseDto
        {
            Resources = new MarketDogBonePriceDto
            {
                CurrentDogCoinsPrice = _currentDogCoinsPrice
            }
        };
    }

    public MarketDogBonePriceResponseDto HandleTrade(MarketDogBoneTradeRequestDto request)
    {
        var trade = request.Resources;
        var type = trade.Type.Trim().ToLower();

        if (type != "buy" && type != "sell")
        {
            throw new ArgumentException("Type must be either 'buy' or 'sell'.");
        }

        if (trade.Amount <= 0)
        {
            throw new ArgumentException("Amount must be higher than 0.");
        }

        if (type == "buy")
        {
            _dogBoneDemand += trade.Amount;
            _dogBoneSupply = Math.Max(0, _dogBoneSupply - trade.Amount);
        }
        else
        {
            _dogBoneSupply += trade.Amount;
            _dogBoneDemand = Math.Max(0, _dogBoneDemand - trade.Amount);
        }

        _currentDogCoinsPrice = CalculatePrice(
            BasePrices[DogCoins],
            _currentDogCoinsPrice,
            _dogBoneSupply,
            _dogBoneDemand
        );

        _tradeHistory.Add(new MarketTradeHistory
        {
            Type = type,
            Amount = trade.Amount,
            PriceAtTrade = _currentDogCoinsPrice,
            SupplyAfterTrade = _dogBoneSupply,
            DemandAfterTrade = _dogBoneDemand,
            CreatedAt = DateTime.UtcNow
        });

        UpdatePriceHistory(DogCoins, _currentDogCoinsPrice);

        return GetPrices();
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
            previousPrice * Balancing.PreviousPriceWeight +
            rawPrice * Balancing.NewPriceWeight;

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

    private class MarketTradeHistory
    {
        public string Type { get; set; } = string.Empty;
        public int Amount { get; set; }
        public decimal PriceAtTrade { get; set; }
        public int SupplyAfterTrade { get; set; }
        public int DemandAfterTrade { get; set; }
        public DateTime CreatedAt { get; set; }
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

    public MarketDogBonePriceResponseDto RunBotTrade()
    {
        var botTrade = DecideBotTrade();

        if (botTrade is null)
        {
            return GetPrices();
        }

        return HandleTrade(new MarketDogBoneTradeRequestDto
        {
            Resources = botTrade
        });
    }

    private MarketDogBoneTradeDto? DecideBotTrade()
    {
        if (_currentDogCoinsPrice <= 4m)
        {
            return new MarketDogBoneTradeDto
            {
                Type = "buy",
                Amount = 10
            };
        }

        if (_currentDogCoinsPrice >= 8m)
        {
            return new MarketDogBoneTradeDto
            {
                Type = "sell",
                Amount = 10
            };
        }

        return null;
    }

    public decimal GetCurrentDogCoinsPrice()
    {
        return _currentDogCoinsPrice;
    }
}