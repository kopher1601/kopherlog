package jp.co.kopher.kopherlog.repository

import com.querydsl.jpa.impl.JPAQueryFactory
import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.domain.QPost
import jp.co.kopher.kopherlog.request.PostSearch

class PostRepositoryImpl(
    private val jpaQueryFactory: JPAQueryFactory,
) : PostRepositoryCustom {
    override fun getList(search: PostSearch): List<Post> {
        return jpaQueryFactory.selectFrom(QPost.post)
            .limit(search.size.toLong())
            .offset(search.offset())
            .orderBy(QPost.post._id.desc())
            .fetch()
    }

}