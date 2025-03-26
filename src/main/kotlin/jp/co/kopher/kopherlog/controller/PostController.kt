package jp.co.kopher.kopherlog.controller

import jakarta.validation.Valid
import jp.co.kopher.kopherlog.request.PostCreate
import jp.co.kopher.kopherlog.response.PostResponse
import jp.co.kopher.kopherlog.service.PostService
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*

@RestController
class PostController(
    private val postService: PostService,
) {

    @ResponseStatus(HttpStatus.CREATED)
    @PostMapping("/posts")
    fun post(@RequestBody @Valid request: PostCreate) {
        postService.write(request)
    }

    @GetMapping("/posts/{postId}")
    fun get(@PathVariable postId: Long): PostResponse {
        return postService.get(postId)
    }

    @GetMapping("/posts")
    fun getList(): List<PostResponse> {
        return postService.getList()
    }

}