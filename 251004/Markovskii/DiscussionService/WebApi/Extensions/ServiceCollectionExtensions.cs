﻿using Infrastructure.Kafka;
using WebApi.Kafka;
using WebApi.Settings;

namespace WebApi.Extensions;

public static class ServiceCollectionExtensions
{
    public static IServiceCollection AddKafkaConsumers(
        this IServiceCollection services,
        IConfigurationRoot config
    )
    {
        services.Configure<KafkaSettings>(config.GetSection("Kafka"));

        services.AddSingleton<KafkaResponseDispatcher<PostResponse>>();
        services.AddSingleton<PostHandler>();
        services.AddHostedService<KafkaResponseService>();
        return services;
    }
}