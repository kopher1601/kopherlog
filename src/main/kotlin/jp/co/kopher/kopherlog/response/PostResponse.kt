package jp.co.kopher.kopherlog.response

import jp.co.kopher.kopherlog.domain.Post

data class PostResponse(
    val id: Long,
    var title: String,
    val content: String,
) {

    init {
        title = title.substring(0, title.length.coerceAtMost(10))
    }

    companion object {
        fun from(post: Post): PostResponse {
            return PostResponse(
                id = post.id!!,
                title = post.title,
                content = post.content
            )
        }
    }
}