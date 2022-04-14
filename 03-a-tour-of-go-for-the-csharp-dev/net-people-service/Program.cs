var builder = WebApplication.CreateBuilder(args);

builder.WebHost.ConfigureKestrel(options => options.ListenLocalhost(9874));

// Add services to the container.
builder.Services.AddSingleton<IPeopleProvider, HardCodedPeopleProvider>();

// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.MapGet("/people", (IPeopleProvider provider) => provider.GetPeople())
    .WithName("GetPeople");

app.MapGet("/people/{id}",
    async (int id, IPeopleProvider provider) =>
    {
        await Task.Delay(1000);
        var result = provider.GetPerson(id);
        return result switch
        {
            null => Results.NoContent(),
            _ => Results.Json(result)
        };
    })
    .WithName("GetPerson");

app.MapGet("/people/ids", (IPeopleProvider provider) => provider.GetPeople().Select(p => p.Id).ToList())
    .WithName("GetIds");

app.Run();
