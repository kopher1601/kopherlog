package jp.co.kopher.kopherlog.repository

import com.querydsl.jpa.impl.JPAQueryFactory
import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.domain.QPost

class PostRepositoryImpl(
    private val jpaQueryFactory: JPAQueryFactory,
) : PostRepositoryCustom {
    override fun getList(page: Int): List<Post> {
        return jpaQueryFactory.selectFrom(QPost.post)
            .limit(10)
            .offset(((page - 1) * 10).toLong())
            .orderBy(QPost.post._id.desc())
            .fetch()
    }

}