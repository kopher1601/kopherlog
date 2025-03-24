package jp.co.kopher.kopherlog.controller

import jakarta.validation.Valid
import jp.co.kopher.kopherlog.request.PostCreate
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RestController

@RestController
class PostController {

    @PostMapping("/posts")
    fun post(@RequestBody @Valid request: PostCreate): String {
        return "Hello World"
    }

}