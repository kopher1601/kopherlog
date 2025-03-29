package jp.co.kopher.kopherlog.exception

class PostNotFound : KopherlogException("Post not found") {
    override val code: Int
        get() = 404
}