package jp.co.kopher.kopherlog.controller

import jakarta.validation.Valid
import jp.co.kopher.kopherlog.request.PostCreate
import jp.co.kopher.kopherlog.service.PostService
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RestController

@RestController
class PostController(
    private val postService: PostService,
) {

    @PostMapping("/posts")
    fun post(@RequestBody @Valid request: PostCreate) {
        postService.write(request)
    }

}