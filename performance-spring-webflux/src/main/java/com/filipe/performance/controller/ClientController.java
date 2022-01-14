package com.filipe.performance.controller;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.reactive.function.client.WebClient;

import reactor.core.publisher.Flux;
import com.filipe.performance.dto.Product;

@RestController
public class ClientController {
	
	private final WebClient wc = WebClient.builder()
											.baseUrl("http://localhost:8082")
											.build();

    @GetMapping(value = "/performance-webflux")
    public Flux<Product> getUserUsingWebfluxWebclient(@RequestParam long delay) {
        return wc.get()
        		.uri("/product/?delay={delay}", delay)
        		.retrieve()
        		.bodyToFlux(Product.class);
    }

}

