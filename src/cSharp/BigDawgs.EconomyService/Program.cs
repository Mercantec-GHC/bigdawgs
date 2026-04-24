using BigDawgs.EconomyService.Services;

var builder = WebApplication.CreateBuilder(args);

// Add services
builder.Services.AddSingleton<MarketService>();

var app = builder.Build();

app.MapGet("/", () => "Economy Service is running 🐶");

// Realistic endpoint
app.MapGet("/market/prices", (MarketService marketService) =>
{
    return marketService.GetPrices();
});

app.Run();