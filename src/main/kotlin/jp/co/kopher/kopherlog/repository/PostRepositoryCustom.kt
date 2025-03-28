package jp.co.kopher.kopherlog.repository

import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.request.PostSearch

interface PostRepositoryCustom {

    fun getList(search: PostSearch): List<Post>

}