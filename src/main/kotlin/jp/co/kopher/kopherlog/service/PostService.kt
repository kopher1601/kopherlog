package jp.co.kopher.kopherlog.service

import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.repository.PostRepository
import jp.co.kopher.kopherlog.request.PostCreate
import jp.co.kopher.kopherlog.response.PostResponse
import org.slf4j.LoggerFactory
import org.springframework.data.domain.Pageable
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

    fun getList(pageable: Pageable): List<PostResponse> {
        return postRepository.getList(1).map { PostResponse.from(it) }.toList()
    }

}