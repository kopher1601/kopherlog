package jp.co.kopher.kopherlog.repository

import jp.co.kopher.kopherlog.domain.Post

interface PostRepositoryCustom {

    fun getList(page: Int): List<Post>

}