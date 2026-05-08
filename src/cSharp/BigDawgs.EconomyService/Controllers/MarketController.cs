using System.Security.Claims;
using BigDawgs.EconomyService.DTOs;
using BigDawgs.EconomyService.Services;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace BigDawgs.EconomyService.Controllers;

[ApiController]
[Route("market")]
[Authorize]
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
        var userId = GetUserId();

        if (string.IsNullOrWhiteSpace(userId))
            return Unauthorized("User id was not found in JWT.");

        return Ok(_marketService.GetPrices());
    }

    [HttpPost("trade")]
    public ActionResult<MarketDogBonePriceResponseDto> CalculatePrices(
        [FromBody] MarketDogBoneTradeRequestDto request)
    {
        var userId = GetUserId();

        if (string.IsNullOrWhiteSpace(userId))
            return Unauthorized("User id was not found in JWT.");

        return Ok(_marketService.HandleTrade(request, userId));
    }

    [HttpPost("bot/trade")]
    public ActionResult<MarketDogBonePriceResponseDto> RunBotTrade()
    {
        var userId = GetUserId();

        if (string.IsNullOrWhiteSpace(userId))
            return Unauthorized("User id was not found in JWT.");

        return Ok(_marketService.RunBotTrade());
    }

    [HttpPost("bot/simulate")]
    public ActionResult<MarketDogBonePriceResponseDto> SimulateBotTrade()
    {
        var userId = GetUserId();

        if (string.IsNullOrWhiteSpace(userId))
            return Unauthorized("User id was not found in JWT.");

        return Ok(_botService.SimulateBotTrade());
    }

    [HttpGet("history")]
    public ActionResult<MarketTradeHistoryResponseDto> GetTradeHistory([FromQuery] int limit = 20)
    {
        var userId = GetUserId();

        if (string.IsNullOrWhiteSpace(userId))
            return Unauthorized("User id was not found in JWT.");

        return Ok(_marketService.GetTradeHistory(limit));
    }

    private string? GetUserId()
    {
        return User.FindFirstValue(ClaimTypes.NameIdentifier)
            ?? User.FindFirstValue("sub")
            ?? User.FindFirstValue("id")
            ?? User.FindFirstValue("user_id");
    }
}