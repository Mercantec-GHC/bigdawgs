using BigDawgs.EconomyService.DTOs;
using BigDawgs.EconomyService.Services;
using Microsoft.AspNetCore.Mvc;


namespace BigDawgs.EconomyService.Controllers;

[ApiController]
[Route("market")]
public class MarketController : ControllerBase
{
    private readonly MarketService _marketService;
    private readonly BotService _botService;

    public MarketController(MarketService marketService, BotService botService)
    {
        _marketService = marketService;
        _botService = botService;
    }

    [HttpGet("prices")]
    public ActionResult<MarketDogBonePriceResponseDto> GetPrices()
    {
        return Ok(_marketService.GetPrices());
    }

    [HttpPost("trade")]
    public ActionResult<MarketDogBonePriceResponseDto> CalculatePrices([FromBody] MarketDogBoneTradeRequestDto request)
    {
        return Ok(_marketService.HandleTrade(request));
    }

    [HttpPost("bot/trade")]
    public ActionResult<MarketDogBonePriceResponseDto> RunBotTrade()
    {
        return Ok(_marketService.RunBotTrade());
    }

    [HttpPost("bot/simulate")]
    public ActionResult<MarketDogBonePriceResponseDto> SimulateBotTrade()
    {
        return Ok(_botService.SimulateBotTrade());
    }
}