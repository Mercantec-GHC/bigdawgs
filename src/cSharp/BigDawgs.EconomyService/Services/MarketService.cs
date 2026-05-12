using BigDawgs.EconomyService.DTOs;

namespace BigDawgs.EconomyService.Services;

public class MarketService
{
    private static readonly PriceBalancingSettings Balancing = new();

    private const string DogCoins = "DogCoins";

    private readonly object _lock = new();

    private readonly Dictionary<string, List<decimal>> PriceHistory =
        new(StringComparer.OrdinalIgnoreCase);

    private static readonly Dictionary<string, decimal> BasePrices =
        new(StringComparer.OrdinalIgnoreCase)
        {
            [DogCoins] = 5m
        };

    private int _dogBoneSupply = 100;
    private int _dogBoneDemand = 100;
    private decimal _currentDogCoinsPrice = 5m;

    private readonly List<MarketTradeHistory> _tradeHistory = new();

    private static readonly Random Random = new();

    public MarketDogBonePriceResponseDto GetPrices(
        decimal? priceAtTrade = null,
        decimal? tradeValue = null)
    {
        lock (_lock)
        {
            return new MarketDogBonePriceResponseDto
            {
                Resources = new MarketDogBonePriceDto
                {
                    CurrentDogCoinsPrice = _currentDogCoinsPrice,
                    PriceAtTrade = priceAtTrade,
                    TradeValue = tradeValue
                }
            };
        }
    }

    public MarketDogBonePriceResponseDto HandleTrade(
        MarketDogBoneTradeRequestDto request,
        string userId = "bot")
    {
        lock (_lock)
        {
            var trade = request.Resources;
            var type = trade.Type.Trim().ToLower();

            if (type != "buy" && type != "sell")
                throw new ArgumentException(
                    "Type must be either 'buy' or 'sell'.");

            if (trade.Amount <= 0)
                throw new ArgumentException(
                    "Amount must be higher than 0.");

            var priceBeforeTrade = _currentDogCoinsPrice;
            var tradeValue = priceBeforeTrade * trade.Amount;

            if (type == "buy")
            {
                _dogBoneDemand += trade.Amount;
                _dogBoneSupply = Math.Max(
                    0,
                    _dogBoneSupply - trade.Amount);
            }
            else
            {
                _dogBoneSupply += trade.Amount;
                _dogBoneDemand = Math.Max(
                    0,
                    _dogBoneDemand - trade.Amount);
            }

            _currentDogCoinsPrice = CalculatePrice(
                BasePrices[DogCoins],
                _currentDogCoinsPrice,
                _dogBoneSupply,
                _dogBoneDemand,
                type,
                trade.Amount
            );

            _tradeHistory.Add(new MarketTradeHistory
            {
                UserId = userId,
                Type = type,
                Amount = trade.Amount,
                PriceAtTrade = priceBeforeTrade,
                TradeValue = tradeValue,
                SupplyAfterTrade = _dogBoneSupply,
                DemandAfterTrade = _dogBoneDemand,
                CreatedAt = DateTime.UtcNow
            });

            if (_tradeHistory.Count > 20)
            {
                _tradeHistory.RemoveAt(0);
            }

            UpdatePriceHistory(
                DogCoins,
                _currentDogCoinsPrice);

            return new MarketDogBonePriceResponseDto
            {
                Resources = new MarketDogBonePriceDto
                {
                    CurrentDogCoinsPrice = _currentDogCoinsPrice,
                    PriceAtTrade = priceBeforeTrade,
                    TradeValue = tradeValue
                }
            };
        }
    }

    private static decimal CalculatePrice(
        decimal basePrice,
        decimal previousPrice,
        int supply,
        int demand,
        string tradeType,
        int tradeAmount)
    {
        supply = Math.Max(0, supply);
        demand = Math.Max(0, demand);

        var minPrice = basePrice * Balancing.MinPriceMultiplier;
        var maxPrice = basePrice * Balancing.MaxPriceMultiplier;

        var impactPerUnit = 0.01m;

        var marketPressure =
            supply > 0
                ? (decimal)demand / supply
                : 2m;

        var pressureModifier =
            Math.Clamp(
                marketPressure,
                0.5m,
                2m);

        var impact =
            1m + (tradeAmount * impactPerUnit * pressureModifier);

        var newPrice =
            tradeType == "buy"
                ? previousPrice * impact
                : previousPrice / impact;

        return Math.Round(
            Math.Clamp(newPrice, minPrice, maxPrice),
            2);
    }

    private List<decimal> UpdatePriceHistory(
        string resourceType,
        decimal currentPrice)
    {
        if (!PriceHistory.ContainsKey(resourceType))
        {
            PriceHistory[resourceType] =
                new List<decimal>();
        }

        PriceHistory[resourceType]
            .Add(currentPrice);

        if (PriceHistory[resourceType].Count >
            Balancing.MaxHistoryEntries)
        {
            PriceHistory[resourceType]
                .RemoveAt(0);
        }

        return new List<decimal>(
            PriceHistory[resourceType]);
    }

    private class MarketTradeHistory
    {
        public string UserId { get; set; } = string.Empty;
        public string Type { get; set; } = string.Empty;
        public int Amount { get; set; }
        public decimal PriceAtTrade { get; set; }
        public decimal TradeValue { get; set; }
        public int SupplyAfterTrade { get; set; }
        public int DemandAfterTrade { get; set; }
        public DateTime CreatedAt { get; set; }
    }

    public MarketTradeHistoryResponseDto GetTradeHistory(
        int limit = 20)
    {
        lock (_lock)
        {
            limit = Math.Clamp(limit, 1, 20);

            return new MarketTradeHistoryResponseDto
            {
                Resources = _tradeHistory
                    .OrderByDescending(x => x.CreatedAt)
                    .Take(limit)
                    .Select(x => new MarketTradeHistoryDto
                    {
                        UserId = x.UserId,
                        Type = x.Type,
                        Amount = x.Amount,
                        PriceAtTrade = x.PriceAtTrade,
                        TradeValue = x.TradeValue,
                        SupplyAfterTrade = x.SupplyAfterTrade,
                        DemandAfterTrade = x.DemandAfterTrade,
                        CreatedAt = x.CreatedAt
                    })
                    .ToList()
            };
        }
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

        return HandleTrade(
            new MarketDogBoneTradeRequestDto
            {
                Resources = botTrade
            },
            "bot");
    }

    private MarketDogBoneTradeDto? DecideBotTrade()
    {
        lock (_lock)
        {
            string type;

            if (_currentDogCoinsPrice <= 4m)
            {
                type =
                    Random.Next(100) < 75
                        ? "buy"
                        : "sell";
            }
            else if (_currentDogCoinsPrice >= 8m)
            {
                type =
                    Random.Next(100) < 75
                        ? "sell"
                        : "buy";
            }
            else
            {
                type =
                    Random.Next(2) == 0
                        ? "buy"
                        : "sell";
            }

            var amount = Random.Next(1, 21);

            return new MarketDogBoneTradeDto
            {
                Type = type,
                Amount = amount
            };
        }
    }

    public decimal GetCurrentDogCoinsPrice()
    {
        lock (_lock)
        {
            return _currentDogCoinsPrice;
        }
    }
}