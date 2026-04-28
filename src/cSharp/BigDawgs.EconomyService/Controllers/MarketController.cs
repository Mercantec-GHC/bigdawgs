using BigDawgs.EconomyService.DTOs;
using BigDawgs.EconomyService.Services;
using Microsoft.AspNetCore.Mvc;

namespace BigDawgs.EconomyService.Controllers;

[ApiController]
[Route("market")]
public class MarketController : ControllerBase
{
    private readonly MarketService _marketService;

    public MarketController(MarketService marketService)
    {
        _marketService = marketService;
    }

    [HttpGet("prices")]
    public ActionResult<EconomyCalculationResponseDto> GetPrices()
    {
        return Ok(_marketService.GetPrices());
    }

    [HttpPost("prices/calculate")]
    public ActionResult<EconomyCalculationResponseDto> CalculatePrices([FromBody] EconomyCalculationRequestDto request)
    {
        return Ok(_marketService.CalculatePrices(request));
    }
}