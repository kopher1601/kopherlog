package jp.co.kopher.kopherlog.service

import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.repository.PostRepository
import jp.co.kopher.kopherlog.request.PostCreate
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Service

@Service
class PostService(
    private val postRepository: PostRepository,
) {

    val log = LoggerFactory.getLogger(PostService::class.java)

    fun write(postCreate: PostCreate) {

        val post = Post(postCreate.title, postCreate.title)
        postRepository.save(post)
    }

}