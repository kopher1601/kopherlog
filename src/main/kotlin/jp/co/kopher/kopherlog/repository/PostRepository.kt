package jp.co.kopher.kopherlog.repository

import jp.co.kopher.kopherlog.domain.Post
import org.springframework.data.jpa.repository.JpaRepository

interface PostRepository : JpaRepository<Post, Long>, PostRepositoryCustom