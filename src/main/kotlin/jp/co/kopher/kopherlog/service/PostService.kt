package jp.co.kopher.kopherlog.service

import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.repository.PostRepository
import jp.co.kopher.kopherlog.request.PostCreate
import jp.co.kopher.kopherlog.request.PostEdit
import jp.co.kopher.kopherlog.request.PostSearch
import jp.co.kopher.kopherlog.response.PostResponse
import org.slf4j.LoggerFactory
import org.springframework.data.repository.findByIdOrNull
import org.springframework.stereotype.Service

@Service
class PostService(
    private val postRepository: PostRepository,
) {

    val log = LoggerFactory.getLogger(PostService::class.java)

    fun write(postCreate: PostCreate) {
        val post = Post(
            _title = postCreate.title,
            _content = postCreate.content
        )
        postRepository.save(post)
    }

    fun get(id: Long): PostResponse {
        val post = postRepository.findByIdOrNull(id)
            ?: throw IllegalArgumentException("Post not found")

        return PostResponse(
            id = post.id!!,
            title = post.title,
            content = post.content
        )
    }

    fun getList(search: PostSearch): List<PostResponse> {
        return postRepository.getList(search).map { PostResponse.from(it) }.toList()
    }

    fun edit(id: Long, postEdit: PostEdit) {
        val post = postRepository.findByIdOrNull(id)
            ?: throw IllegalArgumentException("Post not found")

        post.edit(postEdit)
    }

}