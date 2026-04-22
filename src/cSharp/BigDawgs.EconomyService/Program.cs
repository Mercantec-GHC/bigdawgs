using BigDawgs.EconomyService.Services;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddSingleton<MarketService>();
builder.Services.AddControllers();

var app = builder.Build();

app.MapControllers();

app.Run();