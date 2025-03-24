package jp.co.kopher.kopherlog.response

data class PostResponse(
    val id: Long,
    var title: String,
    val content: String,
) {

    init {
        title = title.substring(0, title.length.coerceAtMost(10))
    }
}