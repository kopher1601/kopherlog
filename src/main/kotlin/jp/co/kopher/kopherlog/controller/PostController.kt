package jp.co.kopher.kopherlog.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class PostController {

    @GetMapping("/posts")
    fun get(): String {
        return "Hello World"
    }

}