package jp.co.kopher.kopherlog.request

import kotlin.math.max
import kotlin.math.min

data class PostSearch(
    val page: Int = 1,
    val size: Int = 10,
) {

    fun offset(): Long {
        val MAX_SIZE = 2000
        return ((max(page, 1) - 1) * min(size, MAX_SIZE)).toLong()
    }
}