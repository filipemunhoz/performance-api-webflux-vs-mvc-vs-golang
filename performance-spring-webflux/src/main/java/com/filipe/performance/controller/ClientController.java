package com.filipe.performance.controller;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.reactive.function.client.WebClient;

import reactor.core.publisher.Mono;

@RestController
public class ClientController {
	

    @GetMapping(value = "/performance-webflux")
    public Mono<List> getUserUsingWebfluxWebclient(@RequestParam long delay) {
        return WebClient.builder()
        		.baseUrl("http://localhost:8082")
        		.build()
        		.get()
        		.uri("/product/?delay={delay}", delay)
        		.retrieve()
        		.bodyToMono(List.class);
    }

}

