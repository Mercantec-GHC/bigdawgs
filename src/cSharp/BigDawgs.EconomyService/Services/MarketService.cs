namespace BigDawgs.EconomyService.Services
{
    public class MarketService
    {
        public object GetPrices()
        {
            return new
            {
                bonesOfMeat = 10,
                dogCoins = 5
            };
        }
    }
}
